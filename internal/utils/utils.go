package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStructure(b any) {
	res, err := json.MarshalIndent(b, "\n", "\t")
	if err != nil {
		fmt.Println("Error print structure: ", err)
	}
	fmt.Println(string(res))
}
