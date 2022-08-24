from functools import wraps
from flask import request, jsonify

import jwt
import os

secret_key = os.getenv('JWT_SECRET')

def token_required(f):
    @wraps(f)
    def decorated(*args, **kwargs):
        token = None

        if 'Authorization' in request.headers:
            auth_header = request.headers['Authorization'].split()
            if len(auth_header) >= 2:
                token = auth_header[1]  

        if token is None:
            return jsonify({
                'msg' : 'invalid header'
            }), 400

        try:
            data = jwt.decode(token, secret_key, algorithms='HS256')
        except Exception as e:
            print(f'Error: {e}')
            return jsonify({
                'msg' : 'invalid token'
            }), 401

        return f(data, *args, **kwargs)
  
    return decorated