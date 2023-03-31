package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func splitRow(data string) (string, string) {
	i := strings.LastIndex(data, ",")
	if i > 0 {
		return data[0:i], data[i+1:]
	} else {
		return data[:len(data)-1], ""
	}
}

func fileForEachLine(path string, callback func(string)) {
	var err error
	var rd *bufio.Reader
	var line string
	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	rd = bufio.NewReader(f)
	for {
		line, err = rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		callback(line[:len(line)-1])
	}
}
func fileForEachLineCreate(path string, callback func(string)) {
	var err error
	var rd *bufio.Reader
	var line string
	var f *os.File
	f, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd = bufio.NewReader(f)
	for {
		line, err = rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		callback(line[:len(line)-1])
	}
}
