import helpers.caching as cached
import dateutil.parser as date_parser

import statistics
import datetime
import requests

def _get_currency_rates(base, target):
    '''
    _get_currency_rates get currency rate from base to target
    '''
    api_resource_endpoint = f'https://v6.exchangerate-api.com/v6/62014c6b2760b7b441a3dab8/latest/{base.upper()}'

    def get_api_resource():
        return requests.get(api_resource_endpoint)

    req = cached.get_cache(api_resource_endpoint, 120, get_api_resource)

    resp = req.json()

    if resp['result'] != 'success':
        raise Exception('failed to get currency rates')

    if target.upper() not in resp['conversion_rates']:
        raise Exception('target currency not found')

    return resp['conversion_rates'][target.upper()]

def _get_summary_provinces(data):
    '''
    _get_summary_provinces get summary aggregate of provinces data from the resource api
    '''
    provinces = dict()

    for item in data:
        if 'area_provinsi' in item and item['area_provinsi'] is not None:
            province = item['area_provinsi']
            if province not in provinces:
                provinces[province] = {
                    'sizes': [],
                    'prices': []
                }
                
            if 'price' in item and item['price'] is not None:
                provinces[province]['prices'].append(int(item['price']))

            if 'size' in item and item['size'] is not None:
                provinces[province]['sizes'].append(int(item['size']))
    
    res = dict()

    for province, values in provinces.items():
        res[province] = {
            'size': {
                'min': min(values['sizes']),
                'max': max(values['sizes']),
                'median': statistics.median(values['sizes']),
                'avg': statistics.mean(values['sizes']),
            },
            'price': {
                'min': min(values['prices']),
                'max': max(values['prices']),
                'median': statistics.median(values['prices']),
                'avg': statistics.mean(values['prices']),
            },
        }

    return res

def fetch():
    '''
    fetch get data from the resource api and add price_usd field
    '''
    api_resource_endpoint = 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list'

    def get_api_resource():
        return requests.get(api_resource_endpoint)

    req = cached.get_cache(api_resource_endpoint, 60, get_api_resource)

    resp = req.json()
    resp = sorted(resp, key=lambda x: (x['tgl_parsed'] is None, x['tgl_parsed']))

    to_usd_rates = _get_currency_rates('IDR', 'USD')

    for item in resp:
        price_usd = None

        if 'price' in item and item['price'] is not None:
            price_usd = str(round(int(item['price']) * to_usd_rates, 2))

        item.update({ 
            "price_usd": price_usd
        })

    res = resp

    return res

def aggregate():
    '''
    aggregate get data from the resource api then create summary of provinces by weekly
    '''
    api_resource_endpoint = 'https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list'

    def get_api_resource():
        return requests.get(api_resource_endpoint)

    req = cached.get_cache(api_resource_endpoint, 60, get_api_resource)

    resp = req.json()
    resp = sorted(resp, key=lambda x: (x['tgl_parsed'] is None, x['tgl_parsed']))

    res = list()

    # by default, set start_week to current week's monday
    start_week = datetime.datetime.today()
    start_week -=  datetime.timedelta(days=start_week.weekday() % 7)

    # if there is record(s), set start_week to earliest record week's monday
    if len(resp) > 0 and resp[0]['tgl_parsed'] is not None:
        start_week = date_parser.parse(resp[0]['tgl_parsed'])
        start_week -= datetime.timedelta(days=start_week.weekday() % 7)

    end_week = start_week + datetime.timedelta(days=6)

    records_week = []
    for item in resp:
        if 'tgl_parsed' not in item or item['tgl_parsed'] is None:
            continue

        date_parsed = date_parser.parse(item['tgl_parsed'])

        if date_parsed >= end_week:
            # add to list
            summary_provinces = _get_summary_provinces(records_week)
            res.append({
                'start_week': start_week,
                'end_week': end_week,
                'summary_provinces': summary_provinces
            })

            # clear the list for recording next week
            records_week.clear()

            # set new start_week and end_week
            start_week = end_week + datetime.timedelta(days=1)
            end_week = start_week + datetime.timedelta(days=6)

        records_week.append(item)

    # add to list
    summary_provinces = _get_summary_provinces(records_week)
    res.append({
        'start_week': start_week,
        'end_week': end_week,
        'summary_provinces': summary_provinces
    })

    # clear the list
    records_week.clear()

    return res