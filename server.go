package main

import (
	"github.com/taotetek/gogopher"
	"log"
	"net"
	"path/filepath"
	"flag"
)

const (
	host = "localhost"
)

var (
	port = flag.String("port", "8070", "port")
	root = flag.String("root", "/srv/gopher", "root")
)

func init() {
	flag.Parse()
}

func handle(c net.Conn) {

	b := make([]byte, 1024)
	p,_ := filepath.Abs("/srv/gopher")
	p = "gopher://" + p
	gd, _ := gogopher.NewGopherDir(p)

	for {
		r := c.Read(b)
		w := c.Write(gd.ToResponse())
		c.Close()
		break
	}


}

func serve(h string, p string, done chan int) {
	l := net.Listen("tcp", h + ":" + p)
	for {
		c := l.Accept()
		go handle(c)
	}
	done <- 1
}

func main() {
	done := make(chan int)
	go server(host, *port, done)
	<-done
}
