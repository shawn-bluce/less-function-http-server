#!/usr/bin/env python3


import json
import requests


config = json.load(open('config.json'))

host = config.get('host')
port = config.get('port')
protocol = config.get('protocol')

url = '{}://{}:{}/asdf'.format(protocol, host, port)
data = {
    'key_1': 'value_1',
    'key_2': 'value_2'
}
headers = {
    'Content-Type': 'application/json'
}
resp = requests.post(url, json=data, headers=headers)
print(resp)
