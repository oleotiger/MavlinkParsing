package main

import (
	"flag"
	"time"
)

func main() {
	rxaddr := flag.String("addr", ":14550", "address to listen on for mavlink")
	flag.Parse()
	//rxaddr1 := flag.String("addr", ":14551", "address to listen on for bluetooth data")
	//flag.Parse()	

	vehicle := new(Vehicle)
	zipDataA := NewZipDataA(vehicle)
	zipDataB := NewZipDataB(vehicle)
	frameChannel := make(chan Frame, 2)
	PX4Port := NewSerialPort("/dev/ttyAMA0", 4800)

		
	go channel2Port(frameChannel, PX4Port)
	go frame2Channel(frameChannel, 35714*time.Microsecond, zipDataA) //28Hz
	go frame2Channel(frameChannel, 200*time.Millisecond, zipDataB)   //5Hz
	
	go listernTCP(":14551", vehicle)
	//go wristbandRecv(vehicle)
	mavlinkRecv(*rxaddr, vehicle)
}
