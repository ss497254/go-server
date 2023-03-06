package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type T interface{}

func CheckErr(x T, err error) T {
	if err != nil {
		fmt.Println("Closing from CheckErr")
		log.Fatal(err)
	}

	return x
}

func JSON_Stringify(v any) string {
	res, err := json.Marshal(v)

	if err != nil {
		fmt.Println("Unable to parse JSON")
		return "{}"
	}

	return string(res)
}
