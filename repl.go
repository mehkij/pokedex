package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cnfg *config) {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		slice := strings.Fields(text)

		c := getCommands()
		_, ok := c[slice[0]]

		if ok {
			c[slice[0]].callback(cnfg, slice)
		} else {
			fmt.Println("error: unknown command")
		}
	}
}
