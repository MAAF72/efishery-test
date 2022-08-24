from time import time 

# cache dict singleton 
cache = dict()

def get_cache(key, duration, updater):
    '''
    get_cache get cache from cache dict singleton, create if not exist and update if expired
    '''
    if key not in cache:
        cache[key] = {
            'result': updater(),
            'last_request': int(time()),
            'duration': duration,
            'updater': updater,
        }
    else:
        if cache[key]['last_request'] + cache[key]['duration'] < int(time()):
            cache[key]['result'] = cache[key]['updater']()
            cache[key]['last_request'] = int(time())

    return cache[key]['result']