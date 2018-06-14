package main

import (
	"fmt"
	"testing"
	"time"
)

//Testx ...
func TestX(t *testing.T) {
	fmt.Println("Text")
}

func BenchmarkX(b *testing.B) {
	fmt.Println("X")
	time.Sleep(time.Second * 1)
}
func BenchmarkY(b *testing.B) {
	fmt.Println("Y")
	time.Sleep(time.Second * 2)
}

func BenchmarkZ(b *testing.B) {
	fmt.Println("Z")
	time.Sleep(time.Second * 2)
}
