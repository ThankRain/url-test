package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
)

// const name = "test(unlabeled).csv"
// const folder = "result/test"
// const kv = "kv_test.txt"
const name = "train1.csv"
const folder = "result"
const kv = "kv.txt"

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("Hi，%s!\n", body)))
	return nil
}

func main() {
	//fc.StartHttp(HandleHttpRequest)
	//runner()
	//select {}
	//calculator()
	compatAll()
}

// 去除多余的逗号
func removeDot() {
	var suc = 0
	var failed = 0
	fileForEachLine("data/result/test/0.csv", func(line string) {
		url, r := splitRow(line)
		//println(r)
		writer(url + r)
	})
	println("Suc", suc, "Failed", failed)
}

func compatAll() {
	path := "data/result/train1"
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		fileForEachLine(path, func(line string) {
			writer(line)
		})
		return nil
	})
	if err != nil {
		println(err.Error())
		return
	}
}
