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
	typeBody := header + "type "
	typeBody = typeBody + strings.Title(typeName) + " struct {"
	typeBody = typeBody + footer + fmt.Sprintf("\n")
	typeBody = typeBody + fmt.Sprintf("\n")
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	status := InitStatus
	tmp := ""
	type2 := ""
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		tmp = fmt.Sprintf("%v", t)
		if fmt.Sprintf("%T", t) == "string" && status == InitStatus {
			typeBody = typeBody + fmt.Sprintf("\t") +
				strings.Title(tmp) + " "
			status = WaitStatus
			type2 = tmp
		} else if status == WaitStatus {
			typeBody = typeBody + fmt.Sprintf("%T", t) + " `json:\"" + type2 + "\"`" + fmt.Sprintf("\n")
			status = InitStatus
		}
	}
	typeBody = typeBody + "}"
	return typeBody, nil
}

const (
	InitStatus string = "init"
	WaitStatus string = "wait"
)

var header string = `
//go:generate ginger $GOFILE
package main

//@ginger
`
var footer string = `
	Ginger_Created int32 ` + "`" + `json:"ginger_created"` + "`" + `
	Ginger_Id int32 ` + "`" + `json:"ginger_id" gorm:"primary_key"` + "`" + ``
