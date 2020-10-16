package main

import (
	"fmt"
	"reflect"
)
// Basic
type Sample struct {
	Name string
}

// Pro
type SamplePro struct {
	Name string `required:"true" max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	// t.NumField adalah jumlah field
	for i := 0; i < t.NumField(); i++ {
		// Field yang ke berdasarkan i
		field := t.Field(i)
		// pengecekan apakah terdapat tag required yang true
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()

			if value == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	//Biasa
	sample := Sample{"Adam"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	//Lanjutan
	samplePro := SamplePro{"Adam Lesmana"}
	var sampleTypePro reflect.Type = reflect.TypeOf(samplePro)

	fmt.Println(sampleTypePro.NumField())
	fmt.Println(sampleTypePro.Field(0).Name)
	fmt.Println(sampleTypePro.Field(0).Tag.Get("required"))
	fmt.Println(sampleTypePro.Field(0).Tag.Get("max"))
	//ubah slice samplePro Name
	//samplePro.Name = ""
	fmt.Println(samplePro)
	fmt.Println(isValid(samplePro))
}