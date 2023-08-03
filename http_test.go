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
	go func() {
		for i := 0; i < 200; i++ {
			intn := rand.Intn(3)
			fmt.Printf("sleep %d s\n", intn)
			time.Sleep(time.Duration(intn) * time.Second)
			_, _ = http.Get(ur)
		}
	}()
	for i := 0; i < 200; i++ {
		intn := rand.Intn(3)
		fmt.Printf("sleep %d s\n", intn)
		time.Sleep(time.Duration(intn) * time.Second)
		_, _ = http.Get(ur2)
	}
}
