package main

import (
	"os"
	"strconv"
)

var writeFile *os.File

func init() {
	suc, fail := skipNumber()
	writePath := "data/" + folder + "/" + strconv.Itoa(suc+fail) + ".csv"
	writeFile, _ = os.OpenFile(writePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
}

func writer(data string) {
	_, err := writeFile.WriteString(data + "\n")
	if err != nil {
		println("Write Error", err.Error())
		return
	}
}
