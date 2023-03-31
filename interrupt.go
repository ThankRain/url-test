package main

import (
	"os"
	"strconv"
)

func stopTheWorld(suc int, failed int) {
	path := "data/" + kv
	kvFile, _ := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	kvFile.WriteString(strconv.Itoa(suc) + "\n")
	kvFile.WriteString(strconv.Itoa(failed) + "\n")
	kvFile.Close()
	os.Exit(0)
}
