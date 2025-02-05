#!/bin/bash

cd roverd/example-pipelines

for dir in *; do
    echo $dir

    cd $dir
    pwd

    cd actuator
    zip -r ../actuator.zip bin/actuator service.yaml
    cd ..

    cd imaging
    zip -r ../imaging.zip bin/imaging service.yaml Makefile
    cd ..

    cd controller
    zip -r ../controller.zip bin/controller service.yaml
    cd ..
    
    cd transceiver
    zip -r ../transceiver.zip bin/transceiver service.yaml
    cd ..

    cd ..
done

