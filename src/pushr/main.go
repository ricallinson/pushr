package main

import (
	"bufio"
	"flag"
	"github.com/tarm/goserial"
	"log"
	"strings"
	"io"
	"os"
)

func listPorts(port string) (list []string) {
    if len(port) > 0 {
        list = []string{port}
    } else {
        list = listSerialPorts()
    }
    return list
}

func openPort(p string) io.ReadWriteCloser {
	c := &serial.Config{Name: p, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Print(err)
		return nil
	}
	return s
}

func openPorts(list []string) (ports []io.ReadWriteCloser) {
    for _, port := range list {
        ports = append(ports, openPort(port))
    }
    return ports
}

func send(s io.ReadWriteCloser, msg []byte) int {
	n, err := s.Write(msg)
	if err != nil {
		log.Print(err)
		return 0
	}
	return n
}

func listen(ports []io.ReadWriteCloser) {
    reader := bufio.NewReader(os.Stdin)
    for {
        msg, _ := reader.ReadString('\n')
        for _, port := range ports {
            send(port, []byte(strings.TrimSpace(msg)))
        }
    }
}

func main() {

	var p = flag.String("p", "", "the USB port to use")
	var l = flag.Bool("l", false, "list all avliable serial ports")
	flag.Parse()

	if *l {
		log.Printf("%v", listSerialPorts())
		return
	}

	list := listPorts(*p)

    if len(list) == 0 {
        log.Print("No serial ports found.\n")
        return
    }

	ports := openPorts(list)

    if len(ports) == 0 {
        log.Print("No serial ports could be opened.\n")
        return
    }

	listen(ports)
}
