package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintPrettyJson(data interface{}) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
