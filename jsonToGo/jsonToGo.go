/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */
package jsonToGo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func Gen(jsonStream string, typeName string) (string, error) {
	typeBody := ""
	typeName = upperFirstChar(typeName)
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
	return typeName + typeBody, nil
}
