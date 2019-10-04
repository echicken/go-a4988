package a4988

import (
	"time"
	"github.com/stianeikeland/go-rpio"
)

type Driver struct {
	pin_dir rpio.Pin
	pin_step rpio.Pin
	pin_ms1 rpio.Pin
	pin_ms2 rpio.Pin
	pin_ms3 rpio.Pin
	pin_enable rpio.Pin
}

func Init(dir uint8, step uint8, ms1 uint8, ms2 uint8, ms3 uint8, enable uint8) (err error, driver Driver) {

	err = rpio.Open();
	if err != nil {
		return err, driver
	}

	driver.pin_dir = rpio.Pin(dir)
	driver.pin_step = rpio.Pin(step)
	driver.pin_ms1 = rpio.Pin(ms1)
	driver.pin_ms2 = rpio.Pin(ms2)
	driver.pin_ms3 = rpio.Pin(ms3)
	driver.pin_enable = rpio.Pin(enable)

	driver.pin_dir.Output()
	driver.pin_step.Output()
	driver.pin_ms1.Output()
	driver.pin_ms2.Output()
	driver.pin_ms3.Output()
	driver.pin_enable.Output();

	driver.pin_dir.Low()
	driver.pin_step.Low()
	driver.pin_ms1.Low()
	driver.pin_ms2.Low()
	driver.pin_ms3.Low()
	driver.pin_enable.Low()

	return nil, driver

}

func (driver *Driver) Enable() {
	driver.pin_enable.Low()
}

func (driver *Driver) Disable() {
	driver.pin_enable.High()
}

func (driver *Driver) Direction(dir bool) {
	if dir {
		driver.pin_dir.High()
	} else {
		driver.pin_dir.Low()
	}
}

func (driver *Driver) StepSize(ss int) {
	switch ss {
		case 0: // Full
			driver.pin_ms1.Low()
			driver.pin_ms2.Low()
			driver.pin_ms3.Low()
			break
		case 1: // Half
			driver.pin_ms1.High()
			driver.pin_ms2.Low()
			driver.pin_ms3.Low()
			break
		case 2: // Quarter
			driver.pin_ms1.Low()
			driver.pin_ms2.High()
			driver.pin_ms3.Low()
			break
		case 3: // Eighth
			driver.pin_ms1.High()
			driver.pin_ms2.High()
			driver.pin_ms3.Low()
			break
		case 4: // Sixteenth
			driver.pin_ms1.High()
			driver.pin_ms2.High()
			driver.pin_ms3.High()
			break
		default:
			break
	}
}

func (driver *Driver) step() {
	driver.pin_step.High()
	time.Sleep(time.Millisecond)
	driver.pin_step.Low()
	time.Sleep(time.Millisecond)
}

func (driver *Driver) Turn(steps int) {
	for i := 0; i < steps; i++ {
		driver.step()
	}
}

func (driver *Driver) Close() {
	rpio.Close();
}