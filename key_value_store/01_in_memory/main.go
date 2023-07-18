package main

import (
	"fmt"
	"time"
)

type KeyValue struct {
	key string
	value string
}

func writer(channel chan KeyValue, key_value KeyValue) {
	fmt.Println(time.Now(), "writing", key_value)
	channel <- key_value
	fmt.Println(time.Now(), "wrote", key_value)
}

func reader(channel chan KeyValue) {
	key_value := <- channel
	fmt.Println(time.Now(), "read", key_value)
}

func main() {
	channel := make(chan KeyValue)

	// write some values
	go writer(channel, KeyValue{"omer", "rocks"})
	go writer(channel, KeyValue{"go", "socks"})
	go writer(channel, KeyValue{"goldi", "locks"})
	
	// read written values
	go reader(channel)
	go reader(channel)
	go reader(channel)

	time.Sleep(1 * time.Second)
}