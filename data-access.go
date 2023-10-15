package main

import (
	"fmt"
)

var m = make(map[int]string)

func initDataMap() {
	m[1] = "Joe"
}

func getNameById(id int) (string, error) {
	if val, ok := m[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("no name found for %d", id)
}

func putNameById(name string, id int) (string, error) {

	m[id] = name
	return m[id], nil
}
