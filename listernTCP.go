package main

import (
	"net"
)

func listernTCP(hostAndPort string, vehicle *Vehicle) {
	listener := initServer(hostAndPort)

    for {
        conn, err := listener.Accept()
        checkError(err, "Accept: ")
        go connectionHandler(conn,vehicle)
    }
}

func initServer(hostAndPort string) *net.TCPListener {
    serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
    checkError(err, "Resolving address:port failed: '" + hostAndPort + "'")
    listener, err := net.ListenTCP("tcp", serverAddr)
    checkError(err, "ListenTCP: ")
    println("Listening to: ", listener.Addr().String())
    return listener
}


func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}

