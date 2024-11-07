package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		c := getCommands()
		_, ok := c[text]

		if ok {
			c[text].callback()
		} else {
			fmt.Println("error: unknown command")
		}
	}
}
