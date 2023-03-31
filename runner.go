package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func skipNumber() (int, int) {
	path := "data/" + kv
	var num = 0
	var suc = 0
	var failed = 0
	fileForEachLineCreate(path, func(line string) {
		if num == 0 {
			suc, _ = strconv.Atoi(line)
		} else if num == 1 {
			failed, _ = strconv.Atoi(line)
		}
		num += 1
	})
	return suc, failed
}

func runner() {
	var suc, failed = skipNumber()
	var counter = 0
	stop := false
	path := "data/origin/" + name
	//path := "data/origin/test(unlabeled).csv"
	c := make(chan os.Signal, 1)
	st := make(chan bool, 1)
	ch, pro := makeWorker(st)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		st <- true
		stop = true
		go func() {
			time.Sleep(15 * time.Second)
			fmt.Println("\r- Stop the world")
			pro <- Progress{typ: 3}
		}()
	}()
	go progressListener(pro, suc, failed)
	fileForEachLine(path, func(line string) {
		if !stop {
			counter += 1
			if counter > suc+failed {
				url, typ := splitRow(line)
				pro <- Progress{typ: 0}
				ch <- Data{
					url: url,
					typ: typ,
				}
			}
		}
	})
}

func progressListener(progress chan Progress, suc int, failed int) {
	all := 0
	for {
		event := <-progress
		switch event.typ {
		case 0:
			all += 1
		case 1:
			suc += 1
		case 2:
			failed += 1
		default:
			stopTheWorld(suc, failed)
		}
		_, _ = fmt.Fprintf(os.Stdout, "Suc / Failed / All [%d/%d/%d]\n", suc, failed, all)
	}
}
