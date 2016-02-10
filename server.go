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

<<<<<<< HEAD
var (
	port = flag.String("port", "8070", "port")
	root = flag.String("root", "/srv/gopher", "root")
)
=======
func suffix(f *os.FileInfo) string {
	if f.IsDirectory() {
		return "1"
	}
	n := f.Name
	switch {
	case strings.HasPrefix(n, ".html"):
		return "h"
    	case strings.HasPrefix(n, ".mp3"),
        	strings.HasPrefix(n, ".aiff"),
        	string.HasPrefix(n, ".au"):
        	return "s"
	case strings.HasPrefix(n, ".txt"),
        	strings.HasPrefix(n, ".json"),
        	strings.HasPrefix(n, ".md"):
		return "0"
	case strings.HasPrefix(n, ".gif"):
		return "g"
	case strings.HasPrefix(n, ".png"),
		strings.HasPrefix(n, ".jpg"),
		strings.HasPrefix(n, ".jpeg"):
		return "I"
	}
	return "9"
}
>>>>>>> 707142794c9aab72514250168ce9f2838d134a3d

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
