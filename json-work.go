package main

import (
	"github.com/tidwall/gjson"
	"log"
	"os"
	"strconv"
	// "strings"
)

func jsonWork(path string) map[int][]string {
	var tableData = map[int][]string{}

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	titleTable := gjson.Get(string(jsonFile), "title")
	for i := 0; i < len(titleTable.Array()); i++ {
		tableData[0] = append(tableData[0], titleTable.Array()[i].String())

	}
	rowTableAll := gjson.Get(string(jsonFile), "row|@pretty")

	for i := 0; i < len(rowTableAll.Map()); i++ {
		array := rowTableAll.Map()[strconv.Itoa(i+1)]
		for j := 0; j < len(array.Array()); j++ {
			tableData[i+1] = append(tableData[i+1], array.Array()[j].String())
		}

	}
	return tableData
}
