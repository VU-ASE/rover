#!/bin/bash


PROJECT_ROOT=/workspaces/roverd
TEST_FILES=$PROJECT_ROOT/rovervalidate/src/testfiles


# --- /etc/roverd/rover.yaml ---
mkdir -p /etc/roverd

cp $TEST_FILES/roverd-yaml/valid/common.yaml /etc/roverd/rover.yaml
# cp $TEST_FILES/roverd-yaml/invalid/invalid-name.yaml /etc/roverd/rover.yaml

# --- /home/debix/rover/ ---
mkdir -p /home/debix/.rover
chown debix:debix /home/debix/.rover

# --- /etc/rover ---
echo "14" > /etc/rover
echo "bunny" >> /etc/rover
echo -n "debix" | sha256sum | cut -d ' ' -f 1 >> /etc/rover


# cp roverd/examples/imaging-service.yaml /home/debix/.rover/vu-ase/imaging/1.0.0/service.yaml
# cp roverd/examples/controller-service.yaml /home/debix/.rover/vu-ase/controller/1.0.1/service.yaml
# cp roverd/examples/actuator-service.yaml /home/debix/.rover/vu-ase/actuator/1.0.5/service.yaml
