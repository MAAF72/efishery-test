from flask import Blueprint, jsonify
from middlewares.jwt import token_required

import controllers.fish as ctrl_fish

fish = Blueprint('fish', __name__)

@token_required
def _fetch(ctx):
    '''
    _fetch fetch endpoint view
    '''
    return jsonify({
        'msg': ctrl_fish.fetch(),
    }), 200

@token_required
def _aggregate(ctx):
    '''
    _aggregate aggregate endpoint view
    '''
    if ctx['role'] != 'admin':
        return 'access denied', 403

    return jsonify({
        'msg': ctrl_fish.aggregate(),
    }), 200

# register the endpoints
fish.add_url_rule('/fetch', methods=['GET'], view_func=_fetch)
fish.add_url_rule('/aggregate', methods=['GET'], view_func=_aggregate)


