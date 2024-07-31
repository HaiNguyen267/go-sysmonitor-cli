package submenu

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/hainguyen267/go-sysmonitor-cli/utils"

)

type HostSubmenu struct {
	infoMap map[string]string
}

func NewHostSubmenu() *HostSubmenu {
	var infoMapForHost = make(map[string]string)
	infoMapForHost["Host Name"] = "The host name is the unique name assigned to a computer on a network. It is used to identify the machine within the network."
	infoMapForHost["Operating System"] = " The operating system is the software that manages the computer hardware and software resources, providing common services for computer programs."
	infoMapForHost["Platform"] = ": The platform provides detailed information about the operating system, including the specific version and edition."
	infoMapForHost["Kernel Version"] = "The kernel version gives the version number that indicates the specific build and version of the Windows kernel."
	infoMapForHost["Kernel Architecture"] = "The kernel architecture specifies the architecture of the processor the kernel is running on. x86_64 indicates a 64-bit processor architecture, which means the system can handle 64-bit instructions and memory addresses"

	return &HostSubmenu{
		infoMap: infoMapForHost,
	}
}




func (h HostSubmenu) Name() string {
	return "Host information"
}


func (h HostSubmenu) Execute() error {
	hostInfo, err := host.Info()

	if err != nil {
		fmt.Println("Error occured, please try again!")
		return err
	}
	utils.PrintHeading("HOST INFORMATION")
	fmt.Println("Host Name:", hostInfo.Hostname)
	fmt.Println("Operating System:", hostInfo.OS)
	fmt.Println("Platform:", hostInfo.Platform)
	fmt.Println("Kernel Version:", hostInfo.KernelVersion)
	fmt.Println("Kernel Architecture:", hostInfo.KernelArch)
	fmt.Println("Recent boot time:", utils.ConvertEpochSecondToDateTime(int64(hostInfo.BootTime)))
	fmt.Println("Uptime till now:", utils.FormatTimeFromSeconds(int64(hostInfo.Uptime)))


	showExplainOrBackSubmenu(&h.infoMap)
	return nil
}
