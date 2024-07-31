package submenu

import (
	"bufio"
	"fmt"
	"os"
)

type Submenu interface {
	Execute() error
	Name() string
}



var scanner = bufio.NewScanner(os.Stdin)

func showExplainOrBackSubmenu(infoMap *map[string]string) {
	fmt.Print("\n> Press 'e' to get explaination or press any key to continue: ")

	scanner.Scan()
	option := scanner.Text()
	if option == "e" {
		printInfoMap(infoMap)
	}
}

func printInfoMap(infoMap *map[string]string) {

	for key, value := range *infoMap {
		fmt.Printf(" +) %v: %v\n", key, value)
	}
}