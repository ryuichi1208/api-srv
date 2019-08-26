#!/usr/bin/env python
# -*- coding: utf-8 -*-

from flask import Flask
from flask import abort, jsonify, request
import json

app = Flask(__name__)

@app.route('/')
def index():
    return 'Hello world'

@app.route('/api/flask/hello')
def api():
    return jsonify({'ip': request.environ.get('HTTP_X_REAL_IP', request.remote_addr), 'NEW': 'Server'}), 200

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=80)
