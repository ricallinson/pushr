package main

import (
	"bufio"
	"flag"
	"github.com/tarm/goserial"
	"log"
	"strings"
	// "fmt"
	"io"
	"os"
)

func openPort(p string) io.ReadWriteCloser {
	c := &serial.Config{Name: p, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Print(err)
		return nil
	}
	return s
}

func sendMsg(s io.ReadWriteCloser, msg []byte) int {
	n, err := s.Write(msg)
	if err != nil {
		log.Print(err)
		return 0
	}
	return n
}

func main() {

	var p = flag.String("p", "", "the USB port to use")
	var l = flag.Bool("l", false, "list all avliable serial ports")
	flag.Parse()

	if *l {
		log.Printf("%v", listSerialPorts())
		return
	}

	var list []string

	if len(*p) > 0 {
		list = []string{*p}
	} else {
		list = listSerialPorts()
	}

	if len(list) == 0 {
		log.Print("No serial ports found.\n")
	}

	ports := []io.ReadWriteCloser{}

	// Try and open all ports.
	for _, port := range list {
		ports = append(ports, openPort(port))
	}

	if len(ports) == 0 {
		log.Print("No serial ports could be opened.\n")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		for _, port := range ports {
			sendMsg(port, []byte(strings.TrimSpace(msg)))
		}
	}
}
