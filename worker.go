package main

import (
	"net/http"
	"time"
)

func makeWorker(stop chan bool) (chan Data, chan Progress) {
	ch := make(chan Data, 1024)
	progress := make(chan Progress, 1)
	for i := 0; i < 1024; i++ {
		go worker(ch, progress, stop)
	}
	return ch, progress
}

func worker(ch chan Data, pro chan Progress, stop chan bool) {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}
	var st = false
	go func() {
		st = <-stop
	}()
	for {
		if !st {
			task := <-ch
			valid := isValid(client, task.url)
			if valid {
				pro <- Progress{typ: 1}
			} else {
				pro <- Progress{typ: 2}
			}
			go write(task, valid)
		} else {
			break
		}
	}
}

func isValid(client *http.Client, link string) bool {
	url := "http://" + link
	_, err := client.Get(url)
	if err != nil {
		//if strings.Contains(err.Error(), "forcibly") {
		//	//println(link, "=>", "[X]", err.Error())
		//	return isValid(client, link)
		//}
		return false
	}
	//println(link, "=>", "[√]")
	return true
}

func write(task Data, valid bool) {
	var v = "0"
	if valid {
		v = "1"
	}
	writer(task.url + "," + task.typ + "," + v)
}
