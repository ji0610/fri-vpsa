package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Meritev struct {
	vrsta    string
	vrednost float32
}

// 5 - 20 sec => 10 - 40 reps

func temp(ch chan Meritev) {
	for range rand.IntN(30) + 10 {
		rndTemp := rand.Float32()*60 - 20
		ch <- Meritev{vrsta: "temperatura", vrednost: rndTemp}
		time.Sleep(500 * time.Millisecond)
	}

}

func vlaga(ch chan Meritev) {
	for range rand.IntN(30) + 10 {
		rndVlaga := rand.Float32() * 100
		ch <- Meritev{vrsta: "vlaga", vrednost: rndVlaga}
		time.Sleep(500 * time.Millisecond)
	}

}

func tlak(ch chan Meritev) {
	for range rand.IntN(30) + 10 {
		rndTlak := rand.Float32()*100 + 950
		ch <- Meritev{vrsta: "tlak", vrednost: rndTlak}
		time.Sleep(500 * time.Millisecond)
	}

}

func in(input chan bool) {
	fmt.Scanln()
	input <- true
}

func main() {

	ch := make(chan Meritev, 3)
	input := make(chan bool)

	go temp(ch)
	go vlaga(ch)
	go tlak(ch)

	go in(input)

	var enota = map[string]string{
		"temperatura": "˚C",
		"vlaga":       "%",
		"tlak":        "hPa",
	}

	for {
		select {
		case m := <-ch:
			fmt.Printf("%s: %.2f %s\n", m.vrsta, m.vrednost, enota[m.vrsta])
		case <-time.After(5 * time.Second):
			fmt.Println("Sistem je neodziven.")
			return
		case <-input:
			fmt.Println("Prekinitev.")
			return
		}

	}
}
