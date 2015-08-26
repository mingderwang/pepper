/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */
package main

import (
	"fmt"
	"github.com/mingderwang/pepper/jsonToGo"
)

func ExampleGen() {
	var jsonStream string = `
{"message": "Hello World", "size": 32, "number": 1.234}
`
	output, _ := jsonToGo.Gen(jsonStream, "test")
	fmt.Println(output)
	//Output:
////go:generate ginger $GOFILE
//package main
//
////@ginger
//type Test struct {
//	Ginger_Created int32 `json:"ginger_created"`
//	Ginger_Id int32 `json:"ginger_id" gorm:"primary_key"` 
//
//	Message string `json:"message"`
//	Size float64 `json:"size"`
//	Number float64 `json:"number"`
//}
}
