package main

type ZipData struct {
	data, status []uint64
	vehicle      *Vehicle
}

func (v ZipData) zip() uint64 {
	result := v.data[0] // one element at least
	for i := 1; i < len(v.data); i++ {
		result = result*v.status[i] + v.data[i]
	}
	return result
}

type ZipDataInterface interface {
	dataUpdate()
	zip() uint64
	Frame() Frame
}

type ZipDataA struct {
	ZipData
}

func NewZipDataA(vehicle *Vehicle) *ZipDataA {
	frame := new(ZipDataA)
	frame.data = make([]uint64, 6)
	frame.status = make([]uint64, 6)
	frame.status[0] = 3600
	frame.status[1] = 3600
	frame.status[2] = 1800
	frame.status[3] = 3600
	frame.status[4] = 180
	frame.status[5] = 200
	frame.vehicle = vehicle
	return frame
}

func (v ZipDataA) dataUpdate() {
	v.data[0] = uint64((v.vehicle.roll + 180) / 0.1)
	v.data[1] = uint64((v.vehicle.yaw + 180) / 0.1)
	v.data[2] = uint64((v.vehicle.pitch + 90) / 0.1)
	v.data[3] = uint64(v.vehicle.cog)
	v.data[4] = uint64(v.vehicle.vel)
	v.data[5] = uint64(v.vehicle.heart)
}

func (v ZipDataA) Frame() (frame Frame) {
	v.dataUpdate()
	frame.frameHeader = 0x2441
	frame.data = v.zip()
	return
}

type ZipDataB struct {
	ZipData
}

func NewZipDataB(vehicle *Vehicle) *ZipDataB {
	frame := new(ZipDataB)
	frame.data = make([]uint64, 3)
	frame.status = make([]uint64, 3)
	frame.status[0] = 26000
	frame.status[1] = 18000000
	frame.status[2] = 36000000
	frame.vehicle=vehicle
	return frame
}

func (v ZipDataB) dataUpdate() {
	v.data[0] = uint64(v.vehicle.alt)
	v.data[1] = uint64(v.vehicle.lat)
	v.data[2] = uint64(v.vehicle.lon)
}

func (v ZipDataB) Frame() (frame Frame) {
	v.dataUpdate()
	frame.frameHeader = 0x2424
	frame.data = v.zip()
	return
}
