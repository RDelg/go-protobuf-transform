package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/protobuf/jsonpb" // Import from your generated code directory
	// Import from your generated code directory
	// Import your generated code package
	"example.com/protobuf-transformation/pkg"
)

func applyFuncToValue(obj map[string]interface{}, path []string, f func(interface{}) interface{}) {
	if len(path) == 1 {
		key := path[0]
		if value, ok := obj[key]; ok {
			obj[key] = f(value)
		}
	} else {
		key := path[0]
		if child, ok := obj[key].(map[string]interface{}); ok {
			applyFuncToValue(child, path[1:], f)
		} else if children, ok := obj[key].([]interface{}); ok {
			for _, child := range children {
				if child, ok := child.(map[string]interface{}); ok {
					applyFuncToValue(child, path[1:], f)
				}
			}
		}
	}
}

func main() {
	// Read the JSON string from a file or wherever you receive it
	jsonString := `{"object": {"foo": "bar", "array": [{"first": {"second": "str"}} ] } }`
	pathsToTransform := []string{"object.array.first.second", "object.foo"}

	// Create a new instance of your protobuf message type
	message := &pkg.SomeMessage{}

	// Unmarshal the JSON into the protobuf message
	if err := jsonpb.UnmarshalString(jsonString, message); err != nil {
		fmt.Println("Error:", err)
		return
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(message)
	json.Unmarshal(inrec, &inInterface)

	// apply transformation to every path in pathsToTransform
	for _, path := range pathsToTransform {
		applyFuncToValue(inInterface, strings.Split(path, "."), func(value interface{}) interface{} {
			if strValue, ok := value.(string); ok {
				return fmt.Sprintf("obfuscated(%v)", strValue)
			} else {
				fmt.Println("Not a string", value)
			}
			return value
		})
	}

	// Print the modified object
	jsonBytes, err := json.Marshal(inInterface)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

}
