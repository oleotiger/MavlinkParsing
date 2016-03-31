package main

import "time"

type Frame struct {
	frameHeader uint16
	data        uint64
}

func frame2Channel(c chan Frame, period time.Duration, values ZipDataInterface) {
	for {
		c <- values.Frame()
		time.Sleep(period)
	}
}
