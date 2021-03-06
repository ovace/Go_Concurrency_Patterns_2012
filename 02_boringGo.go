package main

import (
	"fmt"
	"math/rand"
	"time"
)

//////////////////////////////////////////////
/// aliases
//////////////////////////////////////////////

var println = fmt.Println

//////////////////////////////////////////////
/// functions
//////////////////////////////////////////////

func boring(msg string) {
	for i := 0; ; i++ {
		println(msg, i)
		ms := time.Duration(rand.Intn(1e3))
		time.Sleep(ms * time.Millisecond)
	}
}

func main() {
	go boring("boring!")
	println("I'm listening")
	time.Sleep(2 * time.Second)
	println("You're boring; I'm leaving")
}
