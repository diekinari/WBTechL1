package taskFourteen

import (
	"fmt"
	"reflect"
)

func isChannel(ch interface{}) bool {
	return reflect.TypeOf(ch).Kind() == reflect.Chan
}

func RecognizeType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Got value of type int")
	case string:
		fmt.Println("Got value of type string")
	case bool:
		fmt.Println("Got value of type bool")
	// Здесь мы могли проверить набор из каналов разных типов, но предпочтем более общее решение
	default:
		if isChannel(value) {
			fmt.Println("Got value of type channel")
		} else {
			fmt.Println("Unknown type")
		}

	}
}
