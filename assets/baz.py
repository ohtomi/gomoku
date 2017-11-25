#!/usr/bin/env python3

import datetime

f = open('baz.txt', 'w')
f.write('HTTP request received at {}'.format(datetime.datetime.now()))
f.close()
