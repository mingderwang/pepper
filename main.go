/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */
package main

import (
	"fmt"
	"github.com/mingderwang/pepper/jsonToGo"
)

func main() {
	var jsonStream string = `
{"Message": "Hello", "Number": 1.234}
`
	output, _ := jsonToGo.Gen(jsonStream, "taipeiCity")
	fmt.Println(output)
}
