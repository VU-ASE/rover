#!/bin/bash

curl -u debix:debix -X GET "http://192.168.0.112/logs/vu-ase/controller/1.4.0?lines=3" -H "Accept: application/json" | python3 -c "import json, sys; [print(__import__('codecs').decode(__import__('codecs').encode(line, 'latin1', 'backslashreplace'), 'unicode_escape')) for line in json.load(sys.stdin)]"
