package convert

import (
	"encoding/json"
	"fmt"
)

func ConvertByteToStruct[T any](data []byte)*T{
	var payload T
	err := json.Unmarshal(data,&payload)
	if err != nil{
		fmt.Println(err)
	}
	return &payload
}