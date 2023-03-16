package ffmpeg

import (
	"fmt"
	"strconv"
	"strings"
)

type KVArgs map[string]interface{}

func (kvArgs KVArgs) toArgs() (args []string) {
	args = make([]string, 0, len(kvArgs)*2)

	// Extract input to make sure it gets appended as last argument
	inputKey := "i"
	input := ""
	if i, exists := kvArgs[inputKey]; exists {
		input = valueToString(i)
		delete(kvArgs, inputKey)
	}

	for key := range kvArgs {
		args = append(args, fmt.Sprintf("-%s", key))
		args = append(args, valueToString(kvArgs[key]))
	}

	if input != "" {
		args = append(args, fmt.Sprintf("-%s", inputKey), input)
	}

	return args
}

func (kvArgs KVArgs) toFilter() (filter string) {
	args := make([]string, 0, len(kvArgs))
	for key := range kvArgs {
		b := strings.Builder{}
		// Convert each KV pair into the format key=value
		b.WriteString(key)
		b.WriteString("=")
		b.WriteString(valueToString(kvArgs[key]))
		args = append(args, b.String())
	}
	// Concatenate all KV pairs into key=value,key=value,... format
	return strings.Join(args, ",")
}

func valueToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 32)
	default:
		return fmt.Sprintf("%v", v)
	}
}
