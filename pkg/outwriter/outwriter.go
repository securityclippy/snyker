package outwriter

import (
	"encoding/json"
	"fmt"
)

type OutWriter struct {
}

func (o *OutWriter) Write(output string, input interface{}) {
	switch output {
	case "json":
		fmt.Println(o.jsonOut(input))
	}
}

func (o *OutWriter) jsonOut(i interface{}) string {
	js, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return ""
	}
	return string(js)
}
