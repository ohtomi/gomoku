#!/usr/bin/env python3

import os
import sys
import json

msg = {
    'greet': 'hello, {}'.format(os.environ['GOMOKU']),
    'method': os.environ['METHOD'],
    'url': sys.argv[1]
}
print(json.dumps(msg))
