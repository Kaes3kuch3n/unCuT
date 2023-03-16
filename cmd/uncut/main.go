package main

import (
	"fmt"
	"os"
)

var commands = map[string]func(){
	"generate": GenerateLead,
	"extract":  Extract,
}

func keys[K comparable, V any](m map[K]V) (keys []K) {
	keys = make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("arg must be one of %v", keys(commands))
		os.Exit(1)
	}
	cmd, exists := commands[args[1]]
	if !exists {
		fmt.Printf("invalid command '%s'; arg must be one of %v", args[0], keys(commands))
		os.Exit(1)
	}

	cmd()
}
