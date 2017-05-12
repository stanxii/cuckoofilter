package main

import (
	"fmt"

	"github.com/stanxii/cuckoofilter"
)

func main() {
	cf := cuckoofilter.New()

	cf.Insert([]byte("http://anglia-live.com"))

	b := cf.Lookup([]byte("Hello"))
	fmt.Printf("cf.Lookup Hello === %v\n", b)

	b2 := cf.Lookup([]byte("http://anglia-live.com"))
	fmt.Printf("cf.Lookup anglia === %v\n", b2)


	cf.Count()

	bd := cf.Delete([]byte("world"))
	fmt.Printf("cf.Lookup world === %v\n", bd)

	bdel := cf.Delete([]byte("http://anglia-live.com"))
	fmt.Printf("cf.Del anglia === %v\n", bdel)

	bquery := cf.Lookup([]byte("http://anglia-live.com"))
	fmt.Printf("cf.Lookup anglia after del === %v\n", bquery)
}

