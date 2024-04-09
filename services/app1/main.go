package main

import (
	"fmt"
	lib1 "github.com/Mexx77/golang-mono/libs/lib1"
)

type Helloer interface {
	Hello(name string) string
}

func main() {
	var v Helloer = lib1.Helloer{}
	fmt.Println(v.Hello("John Doe"))
}
