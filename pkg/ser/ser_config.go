package ser

import (
	"log"

	"github.com/leehayford/go_serial/pkg/utl"
	"go.bug.st/serial.v1"
)

func RegisterRoutes() {
	RegisterPlummetronRoutes()
}

func GetSerialPorts() {
	ports, serr := serial.GetPortsList()
	if serr != nil {
		utl.Log(serr)
	}
	if len(ports) == 0 {
		log.Println("No serial ports found!")
	}
	for _, port := range ports {
		log.Println("Found port: \t", port)
	}
}

func OpenSerialPort(name string) serial.Port {
	log.Println("Opening serial port...")
	mode := &serial.Mode{
		BaudRate: 230400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, serr := serial.Open(name, mode)
	if serr != nil {
		log.Fatal("Error opening serial port\n", serr)
	}
	log.Println("Serial port open...")
	return port
}
