package main

import "time"

type Vehicle struct {
	roll           float32
	yaw            float32
	pitch          float32
	cog            uint16
	vel            uint8
	heart          uint8
	alt            int32
	lat            int32
	lon            int32
	gpsTimeStamp   time.Time
	heartTimeStamp time.Time
	ahrsTimeStamp  time.Time
}
