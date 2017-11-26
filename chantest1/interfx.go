package main

import (
	"fmt"
	"reflect"
)

func icheck(x interface{}) {
	fmt.Printf("X is %v\n", x)

	if _, ok := x.(interface {
		String() string
	}); !ok {
		fmt.Printf("x can't String()\n")
	} else {
		fmt.Printf("x can String() - %s\n", reflect.TypeOf(x).String())
	}

}

type myType struct {
	id int
}

func (m myType) String() string {
	return "HI myType"
}

func interfaceRunit() {
	x := 44
	icheck(x)
	y := "Hi There"
	icheck(y)
	z := myType{id: 3}
	icheck(z)
}
