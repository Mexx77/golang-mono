package main

import (
	"fmt"
	"github.com/Mexx77/golang-mono/pkg/generic_data_types/set"
	"github.com/Mexx77/golang-mono/pkg/lib1"
	"github.com/Mexx77/golang-mono/services/app1/internal/some_package"
)

type Helloer interface {
	Hello(name string) string
}

func main() {
	var v Helloer = lib1.Helloer{}
	fmt.Println(v.Hello("John Doe"))

	stringSet := set.Make[string]()
	someService := some_package.New(stringSet)
	someService.Work()
}
