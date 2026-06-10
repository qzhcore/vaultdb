package main

import (
	"fmt"
	"vaultdb/internal/storage"
)

func main() {
	engine, err := storage.NewEngine("vault.log")
	if err != nil {
		panic(err)
	}

	err = engine.Set("studio", "Darkk Games")
	if err != nil {
		panic(err)
	}

	val, exists := engine.Get("studio")
	if exists {
		fmt.Println("Retrieved:", val)
	}
}
