package main

import (
	"io/ioutil"
    "strings"
)

// Looks into the "/dev" directory and returns all the files that maybe serial ports.
func listSerialPorts() (list []string) {
	files, _ := ioutil.ReadDir("/dev/")
	for _, f := range files {
        if strings.Index(f.Name(), "usbserial") >= 0 && strings.Index(f.Name(), "usbmodem") >= 0 && strings.Index(f.Name(), "ttyUSB") >= 0 {
            list = append(list, f.Name())
        }
	}
	return list
}
