#!/usr/bin/env python3

import os
import sys

print('hello, {}. method is {}, url is {}'.format(os.environ['GOMOKU'], os.environ['METHOD'], sys.argv[1]))
