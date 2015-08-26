/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */
package main

import (
	"github.com/mingderwang/pepper/jsonToGo"
)

func main() {
	var jsonStream string = `
{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
`
	josnToGo.jsonToGo(jsonStream, "test")
}
