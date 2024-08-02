package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Function to prettify the JSON object by removing empty lists and dictionaries
func prettify(v interface{}) interface{} {
	switch vv := v.(type) {
	case map[string]interface{}:
		for k, val := range vv {
			if newVal := prettify(val); newVal == nil {
				delete(vv, k)
			} else {
				vv[k] = newVal
			}
		}
		if len(vv) == 0 {
			return nil
		}
		return vv
	case []interface{}:
		newList := make([]interface{}, 0, len(vv))
		for _, item := range vv {
			if newItem := prettify(item); newItem != nil {
				newList = append(newList, newItem)
			}
		}
		if len(newList) == 0 {
			return nil
		}
		return newList
	default:
		return v
	}
}

func main() {
	var t int
	fmt.Scan(&t)

	results := make([]string, 0, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		var buf bytes.Buffer
		for j := 0; j < n; j++ {
			var line string
			fmt.Scanln(&line)
			buf.WriteString(line)
		}

		var data interface{}
		jsonDecoder := json.NewDecoder(&buf)
		jsonDecoder.UseNumber() // Use Number to preserve number types as json.Number

		if err := jsonDecoder.Decode(&data); err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "Error decoding JSON: %v\n", err)
			return
		}

		prettifiedData := prettify(data)
		result, err := json.Marshal(prettifiedData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			return
		}

		results = append(results, string(result))
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
