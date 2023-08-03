package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestRate(t *testing.T) {
	ur := "http://localhost:8089/a"
	ur2 := "http://localhost:8089/tt"
	var work = func(ur string) {
		intn := rand.Intn(3)
		fmt.Printf("%s sleep %d s\n", ur, intn)
		time.Sleep(time.Duration(intn) * time.Second)
		_, _ = http.Get(ur)
	}
	go func() {
		for {
			work(ur)
		}
	}()
	go func() {
		for {
			work(ur2)
		}
	}()
	time.Sleep(6 * time.Minute)
}
