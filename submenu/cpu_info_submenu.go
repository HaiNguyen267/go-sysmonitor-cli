package submenu

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/hainguyen267/go-sysmonitor-cli/utils"
)

type CpuSubmenu struct {
	infoMap map[string]string
}

func NewCpuSubmenu() *CpuSubmenu {
	var infoMapForCpu = make(map[string]string)

	infoMapForCpu["CPU Clock Speed"] = "The CPU clock speed, measured in megahertz (MHz) or gigahertz (GHz), indicates the speed at which a CPU can execute instructions. A higher clock speed means the CPU can process more instructions per second."
	infoMapForCpu["Physical CPU Cores"] = " Physical CPU cores refer to the actual, physical cores in the CPU. Each core can independently execute tasks, which allows for better multitasking and parallel processing."
	infoMapForCpu["Logical CPU Cores"] = "Logical CPU cores, also known as virtual cores. A single physical core to act like two separate cores to the operating system, effectively doubling the number of cores available for handling threads."


	return &CpuSubmenu{
		infoMap: infoMapForCpu,
	}
}


func (c CpuSubmenu) Execute() error {
	procInfo, err := cpu.Info()
	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}
	fmt.Println("This is your cpu details: ")
	utils.PrintHeading("CPU")
	for _, cpu := range procInfo {
		printCpu(cpu)
	}
	
	physicalCpuCount, _ := cpu.Counts(false)
	fmt.Println("Physical CPU cores: ",physicalCpuCount)

	logicalcpuCount, _ := cpu.Counts(true)
	fmt.Println("Logical CPU cores: ",logicalcpuCount)

	// back or explain menu
	showExplainOrBackSubmenu(&c.infoMap)
	return nil
}

func (c CpuSubmenu) Name() string {
	return "Cpu details"
}

func printCpu(cpu cpu.InfoStat) {
	fmt.Printf("CPU #%v\n", cpu.CPU + 1)
	fmt.Println("Model name:", cpu.ModelName)
	fmt.Printf("CPU clock speed: %v MHz\n", cpu.Mhz)
}
