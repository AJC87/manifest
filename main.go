package main

import (
	"fmt"
	"os"
)

func checkArguments() {
	if len(os.Args) < 2 {
		fmt.Println("Argument not provided, exiting...")
		os.Exit(1)
	}
}

func main() {
	checkArguments();

	arguments := os.Args[1:];

	environment := arguments[0];

	fmt.Println("Welcome to manifest!");
	fmt.Println("Selected environment is: " + environment);
}
