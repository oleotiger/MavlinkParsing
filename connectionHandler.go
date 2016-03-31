package main

import (
	"log"
	"net"
)

var (  
	status=0
	readcount=0    
)

func TCPreadByte(port net.Conn) byte {
	buf := make([]byte, 1)
	log.Println("ReadTCP:", uint8(buf[0]))
	_, err := port.Read(buf)
	if err != nil {
		log.Panic(err)
	}
	//read, remoteAddr, err := port.ReadFromUDP(buf)
        //if err != nil {
         //   fmt.Println("read failed:", err)
        //}
        //fmt.Println(read, remoteAddr)
	return buf[0]
}


func connectionHandler(conn net.Conn, vehicle *Vehicle) {
	    connFrom := conn.RemoteAddr().String()
    println("Connection from: ", connFrom)
    //talktoclients(conn)
    for {
        var ibuf []byte = make([]byte,1)
        length, err := conn.Read(ibuf)
	    switch err {
	    case nil:
        	handleMsg(length, err, ibuf,  vehicle)

    	default:
        	goto DISCONNECT
    	}
    }
    DISCONNECT:
    err := conn.Close()
    println("Closed connection:" , connFrom)
    checkError(err, "Close:" )
}

func handleMsg(length int, err error, msg []byte, vehicle *Vehicle) {
	//log.Println("ReadTCP:", uint8(msg[0]))
	if status==1 {
		if msg[0] == 0xFD {
			status=2
		}else {
			status=1
		}
	}else if status==2 {
		if msg[0] == 0x11 {
			status=0
		} else {
			status=1
		}
	}else {
		readcount=readcount+1
		if readcount == 17 {
			if uint8(msg[0]) != 255 {
				vehicle.heart = uint8(msg[0])
			}else{
				vehicle.heart = 0
			}
			log.Println("ReadTCP:", uint8(msg[0]))
		}else if readcount == 27 {
			status =1
			readcount=0
		}
	}
		
}
