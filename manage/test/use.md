# Test script

1. PLug in the Rover
2. Update to latest roverctl `roverctl update`

## Test the basic pipeline

```bash
cd ./test
./all.sh <ROVER_NUMBER> normal install
# Put the Rover on teh ground
./all.sh <ROVER_NUMBER> normal run 
```

## Test the safe pipeline 

```bash
cd ./test
./all.sh <ROVER_NUMBER> safe install
./all.sh <ROVER_NUMBER> safe build
# Put the Rover on teh ground
./all.sh <ROVER_NUMBER> safe run 
```