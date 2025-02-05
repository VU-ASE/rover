#!/bin/bash

# How to zip
# cd imaging
# zip -r ../imaging.zip bin/imaging service.yaml

DIR=normal

sudo rm -rf /home/debix/.rover/vu-ase


curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspaces/roverd/roverd/example-pipelines/$DIR/actuator.zip" \
  http://localhost/upload

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspaces/roverd/roverd/example-pipelines/$DIR/controller.zip" \
  http://localhost/upload

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspaces/roverd/roverd/example-pipelines/$DIR/imaging.zip" \
  http://localhost/upload

curl -u debix:debix \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "content=@/workspaces/roverd/roverd/example-pipelines/$DIR/transceiver.zip" \
  http://localhost/upload



echo
