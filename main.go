package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}
	err = cfg.SetUser("madman")
	if err != nil {
		panic(err)
	}
	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", cfg)
}
