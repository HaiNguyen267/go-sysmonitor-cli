package submenu

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/process"
)

type ProcessSubmenu struct {}

func NewProcessSubmenu() *ProcessSubmenu {
	return &ProcessSubmenu{}
}

func (p ProcessSubmenu) Execute() error {
	processes, err := process.Processes()

	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}
	fmt.Printf("There are %v processes currently running.", len(processes))

	return nil
}

func (p ProcessSubmenu) Name() string {
	return "Number of running processes"
}


