package main

import (
	"fmt"
	"index/suffixarray"
	"os"
)

func sample() {
	index := suffixarray.New([]byte("banana$habana$anna$"))
	fmt.Printf("index of 'nna': %v\n", index.Lookup([]byte("nna"), -1))
	fmt.Printf("index of 'na': %v\n", index.Lookup([]byte("na"), -1))
}

func main() {
	// from https://www.gutenberg.org/ebooks/2701
	b, _ := os.ReadFile("2701.txt.utf-8")
	index := suffixarray.New(b)
	fmt.Printf("index of 'nna': %v\n", index.Lookup([]byte("nna"), -1))
	fmt.Printf("index of 'na': %v\n", index.Lookup([]byte("na"), -1))
}