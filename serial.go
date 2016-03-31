package main

import (
	"encoding/binary"
	"github.com/tarm/serial"
	"log"
)

func NewSerialPort(dev string, baud int) *serial.Port {
	serialConfig := &serial.Config{Name: dev, Baud: baud}
	port, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func channel2Port(c chan Frame, port *serial.Port) {
	for {
		frame := <-c
		binary.Write(port, binary.BigEndian, frame.frameHeader)
		binary.Write(port, binary.LittleEndian, frame.data)
	}

}

func readByte(port *serial.Port) byte {
	buf := make([]byte, 1)
	_, err := port.Read(buf)
	if err != nil {
		log.Panic(err)
	}
	return buf[0]
}
