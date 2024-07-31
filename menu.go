package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"bufio"
	"os"
	"github.com/hainguyen267/go-sysmonitor-cli/submenu"
	"github.com/hainguyen267/go-sysmonitor-cli/utils"

)

var scanner = bufio.NewScanner(os.Stdin)

type Menu struct {
	subMenuList []submenu.Submenu
}



func New() *Menu {
	return &Menu{
		subMenuList: []submenu.Submenu{
			submenu.NewHostSubmenu(),
			submenu.NewCpuSubmenu(),
			submenu.NewMemorySubmenu(),
			submenu.NewDiskSubmenu(),
			submenu.NewProcessSubmenu(),
		},
	}
}

func sayGoodBye() {
	// if now is after 10 PM or before 3 AM
	currentHour := time.Now().Hour()
	if currentHour >= 22 || currentHour <= 3{
		fmt.Println("Goodbye, good night!")
	} else {
		fmt.Println("Goodbye, have a nice day!")
	}
}


func (m *Menu) Execute() {
	var option int64

	for true {
		option = m.getUserInput()
		if option == 0 {
			break
		}

		m.subMenuList[option-1].Execute() // because user input option starting from 1
	}

}


func (m *Menu) getUserInput() int64 {
	var option int64
	var err error
	utils.PrintHeading("System Monitor")
	for index, submenu := range m.subMenuList {
		m.printSubmenuName(index, submenu.Name())
	}
	fmt.Println("0. Exit")

	fmt.Print("\n> Please enter your option: ")
	scanner.Scan()
	option, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)

	maxOption := int64(len(m.subMenuList))
	for err != nil || option < 0 || option > maxOption {
		fmt.Printf("Invalid option, please choose number from 0 to %v: ", maxOption)
		scanner.Scan()
		option, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
	}

	return option
}

func (m *Menu) printSubmenuName(index int, submenuName string) {
	fmt.Printf("%v. %v\n", index + 1, submenuName)// number starts from 1
}