package main

import (
	"fmt"
	"log"
)

const (
	ok = true
	no = false
)

// Person
type Person struct {
	nama  string
	tahun int
}

// Umur comment
type Umur interface {
	getUmur() int
	getTest(x int) int
}

func (p *Person) getUmur() int {
	return 2018 - p.tahun
}

func getTest(x int) int {
	return 2018 - x
}

func main() {
	orang := Person{nama: "dian", tahun: 40}
	fmt.Println(orang.nama)
	fmt.Println(orang.getUmur)
	fmt.Println(getTest(orang.tahun))
	fmt.Println(ok)

	orang2 := &orang
	fmt.Println(orang2.nama)
	orang.nama = "efraim"
	orang2 = nil
	if orang2 == nil {
		log.Println("nil")
	} else {
		fmt.Println(orang2.nama)
	}
}
