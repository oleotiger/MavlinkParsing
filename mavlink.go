package main

import (
	"github.com/liamstask/go-mavlink/mavlink"
	"log"
	"math"
	"net"
	"time"
)

func mavlinkRecv(addr string, vehicle *Vehicle) {

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	conn, listenerr := net.ListenUDP("udp", udpAddr)
	if listenerr != nil {
		log.Fatal(listenerr)
	}

	log.Println("listening on", udpAddr)

	dec := mavlink.NewDecoder(conn)
	dec.Dialects.Add(mavlink.DialectPixhawk)

	for {
		pkt, err := dec.Decode()
		if err != nil {
			//log.Println("Decode fail:", err)
			continue
		}
		//log.Println("Decode success")
		switch pkt.MsgID {
		case mavlink.MSG_ID_ATTITUDE:
			var attitude mavlink.Attitude
			if err := attitude.Unpack(pkt); err == nil {
				vehicle.pitch = attitude.Pitch / math.Pi * 180
				vehicle.yaw = attitude.Yaw / math.Pi * 180
				vehicle.roll = attitude.Roll / math.Pi * 180
				vehicle.ahrsTimeStamp = time.Now()
			}
		case mavlink.MSG_ID_GLOBAL_POSITION_INT:
			var globalPosition mavlink.GlobalPositionInt
			if err := globalPosition.Unpack(pkt); err == nil {
				vehicle.lat = globalPosition.Lat / 100
				vehicle.lon = globalPosition.Lon / 100
				vehicle.alt = globalPosition.Alt / 200
				vehicle.gpsTimeStamp = time.Now()

			}
		case mavlink.MSG_ID_GPS2_RAW:
			var gps2Raw mavlink.Gps2Raw
			if err := gps2Raw.Unpack(pkt); err == nil {
				vehicle.vel = uint8(float32(gps2Raw.Vel) / 100 * 3.6 / 1.852) //unit 0.01m/s convert to knot
				vehicle.cog = gps2Raw.Cog / 10                               // unit 0.01degree convert to 0.1degree
			}
		}
	}
}
