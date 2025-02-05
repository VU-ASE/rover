package main

import (
	"fmt"
	"os"
	"time"

	pb_outputs "github.com/VU-ASE/rovercom/packages/go/outputs"
	roverlib "github.com/VU-ASE/roverlib-go/src"

	"github.com/rs/zerolog/log"
)

// The main user space program
// this program has all you need from roverlib: service identity, reading, writing and configuration
func run(service roverlib.Service, configuration *roverlib.ServiceConfiguration) error {
	if configuration == nil {
		return fmt.Errorf("configuration cannot be accessed")
	}

	//
	// Access the service identity, who am I?
	//
	log.Info().Msgf("Hello world, a new service '%s' was born at version %s", *service.Name, *service.Version)

	//
	// Access the service configuration, to use runtime parameters
	//
	exampleNum, err := configuration.GetFloatSafe("number-example")
	if err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}
	log.Info().Msgf("Fetched runtime configuration example number: %f", exampleNum)

	exampleString, err := configuration.GetStringSafe("string-example")
	if err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}
	log.Info().Msgf("Fetched runtime configuration example string: %s", exampleString)

	exampleStringTunable, err := configuration.GetStringSafe("tunable-string-example")
	if err != nil {
		return fmt.Errorf("failed to get configuration: %v", err)
	}
	log.Info().Msgf("Fetched runtime configuration example tunable string: %s", exampleStringTunable)

	//
	// Writing to an output that other services can read (see service.yaml to understand the output name)
	//
	writeStream := service.GetWriteStream("example-output")
	if writeStream == nil {
		return fmt.Errorf("failed to get write stream")
	}

	// Try to write a simple rovercom message, as if we are sending RPM data
	err = writeStream.Write(&pb_outputs.SensorOutput{
		SensorId:  404,                            // let the receiver know who we are! (if you have multiple sensors in one service)
		Timestamp: uint64(time.Now().UnixMilli()), // current time in milliseconds (useful for debugging)
		Status:    0,                              // we are chilling
		SensorOutput: &pb_outputs.SensorOutput_RpmOuput{
			RpmOuput: &pb_outputs.RpmSensorOutput{
				LeftRpm:  1000,
				RightRpm: 1200,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to write to stream: %v", err)
	}

	// You don't like using protobuf messages? No problem, you can write raw bytes too
	err = writeStream.WriteBytes([]byte("Hello world!"))
	if err != nil {
		return fmt.Errorf("failed to write to stream: %v", err)
	}

	//
	// Reading from an input, to get data from other services (see service.yaml to understand the input name)
	//
	readStream := service.GetReadStream("example-input", "rpm-data")
	if readStream == nil {
		return fmt.Errorf("failed to get read stream")
	}

	// Try to read a simple rovercom message, as if we are receiving RPM data
	pbData, err := readStream.Read()
	if err != nil {
		log.Error().Msgf("failed to read from stream: %v", err)
	} else {
		// Find out if we actually have rpm data
		if pbData.GetRpmOuput() == nil {
			log.Error().Msgf("expected RPM data, but got something else")
		} else {
			log.Info().Msgf("Received RPM data: %f, %f", pbData.GetRpmOuput().GetLeftRpm(), pbData.GetRpmOuput().GetRightRpm())
		}
	}

	// You don't like using protobuf messages? No problem, you can read raw bytes too
	rawData, err := readStream.ReadBytes()
	if err != nil {
		log.Error().Msgf("failed to read from stream: %v", err)
	} else {
		log.Info().Msgf("Received raw data: %s", string(rawData))
	}

	//
	// Now do something else fun, see if our "example-string-tunable" is updated
	//
	curr := exampleStringTunable
	for {
		log.Info().Msg("Checking for tunable string update")

		// We are not using the safe version here, because using locks is boring
		// (this is perfectly fine if you are constantly polling the value)
		// nb: this is not a blocking call, it will return the last known value
		newVal, err := configuration.GetString("tunable-string-example")
		if err != nil {
			return fmt.Errorf("failed to get configuration: %v", err)
		}

		if curr != newVal {
			log.Info().Msgf("Tunable string updated: %s -> %s", curr, newVal)
			curr = newVal
		}

		// Don't waste CPU cycles
		time.Sleep(1 * time.Second)
	}
}

// This function gets called when roverd wants to terminate the service
func onTerminate(sig os.Signal) error {
	log.Info().Str("signal", sig.String()).Msg("Terminating service")

	//
	// ...
	// Any clean up logic here
	// ...
	//

	return nil
}

// This is just a wrapper to run the user program
// it is not recommended to put any other logic here
func main() {
	roverlib.Run(run, onTerminate)
}
