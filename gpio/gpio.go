package gpio

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -lgpio
#include "gpio.h"
*/
import "C"
import "errors"

const (
	SETUP_OK           = int(C.SETUP_OK)
	SETUP_DEVMEM_FAIL  = int(C.SETUP_DEVMEM_FAIL)
	SETUP_MALLOC_FAIL  = int(C.SETUP_MALLOC_FAIL)
	SETUP_MMAP_FAIL    = int(C.SETUP_MMAP_FAIL)
	SETUP_CPUINFO_FAIL = int(C.SETUP_CPUINFO_FAIL)
	SETUP_NOT_RPI_FAIL = int(C.SETUP_NOT_RPI_FAIL)

	INPUT  = int(C.INPUT)
	OUTPUT = int(C.OUTPUT)
	ALT0   = int(C.ALT0)

	HIGH = int(C.HIGH)
	LOW  = int(C.LOW)

	PUD_OFF  = int(C.PUD_OFF)
	PUD_DOWN = int(C.PUD_DOWN)
	PUD_UP   = int(C.PUD_UP)
)

func SetupError(status int) error {
	switch status {
	case SETUP_DEVMEM_FAIL:
		return errors.New("SETUP_DEVMEM_FAIL")
	case SETUP_MALLOC_FAIL:
		return errors.New("SETUP_MALLOC_FAIL")
	case SETUP_MMAP_FAIL:
		return errors.New("SETUP_MMAP_FAIL")
	case SETUP_CPUINFO_FAIL:
		return errors.New("SETUP_CPUINFO_FAIL")
	case SETUP_NOT_RPI_FAIL:
		return errors.New("SETUP_NOT_RPI_FAIL")
	}

	return nil
}

func Setup() int {
	return int(C.setup())
}

func SetupGPIO(gpio, direction, pud int) {
	C.setup_gpio(C.int(gpio), C.int(direction), C.int(pud))
}

func FunctionGPIO(gpio int) int {
	return int(C.gpio_function(C.int(gpio)))
}

func OutputGPIO(gpio, value int) {
	C.output_gpio(C.int(gpio), C.int(value))
}

func InputGPIO(gpio int) int {
	return int(C.input_gpio(C.int(gpio)))
}

func SetRisingEvent(gpio, enabled int) {
	C.set_rising_event(C.int(gpio), C.int(enabled))
}

func SetFallingEvent(gpio, enabled int) {
	C.set_falling_event(C.int(gpio), C.int(enabled))
}

func SetHighEvent(gpio, enable int) {
	C.set_high_event(C.int(gpio), C.int(enable))
}

func SetLowEvent(gpio, enable int) {
	C.set_low_event(C.int(gpio), C.int(enable))
}

func EventDetected(gpio int) int {
	return int(C.eventdetected(C.int(gpio)))
}

func Cleanup() {
	C.cleanup()
}
