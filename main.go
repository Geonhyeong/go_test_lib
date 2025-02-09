package main

import (
	"C"
	v4 "main/v4"

	myItem "github.com/Geonhyeong/go_test_lib/proto/item"
)

//export StarForceWrapper
func StarForceWrapper(myItem *myItem.Item) {
	v4.StarForce(myItem)
}

func main() {}
