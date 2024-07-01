package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	odam := Person{"Fakhriya", 17, "USA"}

	// динамические операции на основе типов.
	t := reflect.TypeOf(odam)

	// Iterate over fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := reflect.ValueOf(odam).Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}