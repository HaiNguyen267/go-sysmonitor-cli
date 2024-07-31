package submenu

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/hainguyen267/go-sysmonitor-cli/utils"
)

type MemorySubmenu struct {
	infoMap map[string]string
}


func NewMemorySubmenu() *MemorySubmenu {
	var infoMap = make(map[string]string)

	infoMap["Swap Device"] = "A swap device is a storage device or partition that is used to extend the amount of virtual memory available to a system. When the system's physical memory (RAM) is fully utilized, the operating system can move inactive pages of memory to the swap device to free up RAM for other tasks."
	infoMap["Swap Memory"] = "Swap memory, also known as virtual memory, is a portion of the storage space that is used as an extension of the system's physical memory (RAM). Swap memory allows the system to handle more processes and data than would fit in the physical RAM alone."

	return &MemorySubmenu{
		infoMap: infoMap,
	}
}

func (m MemorySubmenu) Execute() error {
	virtualMem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}
	
	swapDevices, err := mem.SwapDevices()
	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}

	swagMem, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}

	utils.PrintHeading("MEMORY")
	fmt.Println("Total Memory:",utils.FormatBytes(virtualMem.Total))
	fmt.Println("Used Memory:",utils.FormatBytes(virtualMem.Used))
	fmt.Println("Free Memory:",utils.FormatBytes(virtualMem.Free))
	fmt.Printf("Memory used percent: %v%%\n",virtualMem.UsedPercent)
	utils.PrintHeading("SWAP DEVICE AND SWAP MEMORY")

	for index, device := range swapDevices {
		printSwapDevice(index, device)
	}
	fmt.Println("Total swap memory:", utils.FormatBytes(swagMem.Total))
	fmt.Println("Used swap memory:", utils.FormatBytes(swagMem.Used))
	fmt.Println("Free wwap memory:", utils.FormatBytes(swagMem.Free))
	fmt.Printf("Swap memory used percent: %.2f%%\n", swagMem.UsedPercent)

	showExplainOrBackSubmenu(&m.infoMap)
	return nil
}

func (m MemorySubmenu) Name() string {
	return "Memory details"
}

func printSwapDevice(deviceNo int, device *mem.SwapDevice) {
	fmt.Println("Device #", deviceNo + 1)
	fmt.Println("Device name:", device.Name)
	fmt.Println("Device space used:", utils.FormatBytes(device.UsedBytes))
	fmt.Println("Device space Free:", utils.FormatBytes(device.FreeBytes))
}
