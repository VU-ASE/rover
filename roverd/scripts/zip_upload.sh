#!/bin/bash

# How to zip
# cd imaging
# zip -r ../imaging.zip bin/imaging service.yaml

DIR=valid_alias
ADDRESS=192.168.0.112

sudo rm -rf /home/debix/.rover/vu-ase


curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspace/rover/roverd/roverd/example-pipelines/$DIR/actuator.zip" \
  http://$ADDRESS/upload

echo ""

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspace/rover/roverd/roverd/example-pipelines/$DIR/funny.zip" \
  http://$ADDRESS/upload

echo ""

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspace/rover/roverd/roverd/example-pipelines/$DIR/controller.zip" \
  http://$ADDRESS/upload

echo ""

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspace/rover/roverd/roverd/example-pipelines/$DIR/imaging.zip" \
  http://$ADDRESS/upload


echo ""

echo Done



# curl http://192.168.0.112/pipeline \
#   -u debix:debix \
#   -X POST \
#   -H "Content-Type: application/json" \
#   -d '
#   [
#     {
#       "fq": {
#         "author": "vu-ase",
#         "name": "imaging",
#         "version": "1.0.0"
#       }
#     },
#     {
#       "fq": {
#         "author": "vu-ase",
#         "name": "actuator",
#         "version": "1.0.0"
#       }
#     },
#     {
#       "fq": {
#         "author": "vu-ase",
#         "name": "funny",
#         "version": "1.0.0"
#       }
#     }
#   ]'


