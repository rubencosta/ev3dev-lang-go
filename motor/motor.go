// Package motor provides API for interacting with EV3's motors.
package motor

import (
	"errors"
	"fmt"
)

type port struct {
	port string
}

type command struct {
	id        string
	supported bool
}

type commands []command

// Motor type
type Motor struct {
	command
	commands
	countPerRoot           int
	driverName             string
	dutyCycle              int
	dutyCycleSp            int
	encoderPolarity        string
	polarity               string
	portName               string
	position               int
	positionP              int
	positionI              int
	positionD              int
	positionSp             int
	speed                  int
	speedSp                int
	rampUpSp               int
	rampDownSp             int
	speedRegulationEnabled int
	speedRegulationP       int
	speedRegulationI       int
	speedRegulationD       int
	state                  []string
	stopCommand            string
	stopCommands           []string
	timeSp                 int
	deviceIndex            int
	connected              bool
}

type portName string
type driverName string

const (
	outPortA portName = "outA"
	outPortB portName = "outB"
	outPortC portName = "outC"
	outPortD portName = "outD"
)

// Connect to motor
func Connect(p portName, d driverName) (Motor, error) {
	if p == "" {
		// TODO: If the specified port is blank or unspecified/undefined/null, the
		// available devices should be enumerated until a suitable device is found.
		// Any device is suitable when it's type is known to be compatible with the
		// controlling class, and it meets any other requirements specified by the caller.
	} else {
		// TODO: Look for a motor in the `p` portName
	}

	if d == "" {
		// TODO: The motor driver that should be driving the target motor (generally
		// specifies the type of motor). Can be left empty or undefined (in the
		// languages that support it) to specify a wildcard.
	}
	return Motor{}, errors.New("XXX")
}

// Reset all of the motor's parameters and state
func (m *Motor) Reset() {
	//TODO: Sets the command motor property to reset, which causes the motor driver
	// to reset all of the motor's parameters and state.
	fmt.Printf("asda")
}
