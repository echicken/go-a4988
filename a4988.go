package a4988

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

// Driver contains pin mappings
type Driver struct {
	pinDir    rpio.Pin
	pinStep   rpio.Pin
	pinMS1    rpio.Pin
	pinMS2    rpio.Pin
	pinMS3    rpio.Pin
	pinEnable rpio.Pin
}

// Init returns a new stepper driver
func Init(dir uint8, step uint8, ms1 uint8, ms2 uint8, ms3 uint8, enable uint8) (driver Driver, err error) {

	err = rpio.Open()
	if err != nil {
		return driver, err
	}

	driver.pinDir = rpio.Pin(dir)
	driver.pinStep = rpio.Pin(step)
	driver.pinMS1 = rpio.Pin(ms1)
	driver.pinMS2 = rpio.Pin(ms2)
	driver.pinMS3 = rpio.Pin(ms3)
	driver.pinEnable = rpio.Pin(enable)

	driver.pinDir.Output()
	driver.pinStep.Output()
	driver.pinMS1.Output()
	driver.pinMS2.Output()
	driver.pinMS3.Output()
	driver.pinEnable.Output()

	driver.pinDir.Low()
	driver.pinStep.Low()
	driver.pinMS1.Low()
	driver.pinMS2.Low()
	driver.pinMS3.Low()
	driver.pinEnable.Low()

	return driver, nil

}

// Enable the stepper driver
func (driver *Driver) Enable() {
	driver.pinEnable.Low()
}

// Disable the stepper driver
func (driver *Driver) Disable() {
	driver.pinEnable.High()
}

// Direction true/false alters the stepper's rotational direction
func (driver *Driver) Direction(dir bool) {
	if dir {
		driver.pinDir.High()
	} else {
		driver.pinDir.Low()
	}
}

// StepSize sets the stepper's microstep increment (0 = Full, 1 = Half, 2 = Quarter, 3 = Eighth, 4 = Sixteenth)
func (driver *Driver) StepSize(ss int) {
	switch ss {
	case 0: // Full
		driver.pinMS1.Low()
		driver.pinMS2.Low()
		driver.pinMS3.Low()
		break
	case 1: // Half
		driver.pinMS1.High()
		driver.pinMS2.Low()
		driver.pinMS3.Low()
		break
	case 2: // Quarter
		driver.pinMS1.Low()
		driver.pinMS2.High()
		driver.pinMS3.Low()
		break
	case 3: // Eighth
		driver.pinMS1.High()
		driver.pinMS2.High()
		driver.pinMS3.Low()
		break
	case 4: // Sixteenth
		driver.pinMS1.High()
		driver.pinMS2.High()
		driver.pinMS3.High()
		break
	default:
		break
	}
}

func (driver *Driver) step() {
	driver.pinStep.High()
	time.Sleep(time.Millisecond)
	driver.pinStep.Low()
	time.Sleep(time.Millisecond)
}

// Turn the stepper n steps
func (driver *Driver) Turn(steps int) {
	for i := 0; i < steps; i++ {
		driver.step()
	}
}

// Close rpio
func (driver *Driver) Close() {
	rpio.Close()
}
