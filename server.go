package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var root = flag.String("root", "", "root directory")
var dir string
type fileName []os.FileInfo

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

func getDir() string {
	gr := os.Getenv("GOROOT")
	return filepath.Join(gr, "doc")
}

func serve(n net.Conn) {
	defer n.Close()
	r := bufio.NewReader(n)
	w := bufio.NewWriter(n)
	defer w.Flush()

	l := r.ReadSlice('\n')
	line := string(l)
	line = filepath.Clean(line)
	fileName := filepath.Join(dir, line)
	f := os.Stat(fileName)

	switch {
	case f.IsDirectory():
		e := os.Open(fileName)
		fg := e.Readdir(-1)
		sort.Sort(fileName(fg))

    	for _, f := range fg {
		fmt.Fprintf(w, "%s%s\r\n",suffix(&f),strings.Join([]string{
				f.Name,line + "/" + f.Name,"127.0.0.1","70",}, "\t"))
		}

	case f.IsRegular():
		e := os.Open(fileName)
		io.Copy(w, e)

	default:
		log.Printf("bad file")
	}
}

func main() {
	flag.Parse()
	dir = getDir()
	l := net.Listen("tcp", ":70")
	for {
		go serve(l.Accept())
	}
}
