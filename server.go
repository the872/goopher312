package main 

import (
        "fmt"
	"net"
	"os"
)

const (
      CONN_HOST = "localhost"
      CONN_PORT = "8070"
      CONN_TYPE = "tcp"
)

func main() { 

     //listening for incoming connections
     l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
     if err != nil {
     	fmt.Println("Listening error:", err.Error())
	os.Exit(0)
     }

     defer l.Close()
     fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
     for { 
     	 //listen for incoming connection
	 conn, err := l.Accept()
	 if err != nil {
	    fmt.Println("Error accepting: ", err.Error())
	    os.Exit(0)
	 }

	 //request handling
	 go request(conn)
     }
}

func request(conn net.Conn) {
     //data buffer
     buff := make([]byte, 1024)
     
     //read incoming connection
     reqLen, err := conn.Read(buff)
     if err != nil {
     	fmt.Println("Error reading: ", err.Error())
     }
     
     //response 
     conn.Write ([]byte("Message received!"))

     conn.Close()
}