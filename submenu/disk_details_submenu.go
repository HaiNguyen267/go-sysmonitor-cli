package submenu

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/hainguyen267/go-sysmonitor-cli/utils"
)

type DiskSubmenu struct {
	infoMap map[string]string
}


func NewDiskSubmenu() *DiskSubmenu {

	var infoMapForDisk = make(map[string]string)

	infoMapForDisk["Partition Device"] = "The device name of the partition, typically representing the storage device or logical volume"
	infoMapForDisk["Partition Mountpoint"] = "The directory where the partition is mounted, which is the point in the directory tree where you can access the contents of the partition"
	infoMapForDisk["Partition Filesystem Type"] = "The type of the filesystem used on the partition, such as NTFS, FAT32, ext4, etc."


	return &DiskSubmenu{
		infoMap: infoMapForDisk,
	}
}


func (d DiskSubmenu) Execute() error {
	parittions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Error occured, please try again.")
		return err
	}

	utils.PrintHeading("DISK PARTITIONS AND USAGE")
	for _, parition := range parittions {
		printDiskPartitionAndUsage(&parition)
	}

	showExplainOrBackSubmenu(&d.infoMap)
	return nil
}

func (d DiskSubmenu) Name() string {
	return "Disk details"
}




func printDiskPartitionAndUsage(partition *disk.PartitionStat) {
	fmt.Println("Partion Device:", partition.Device)
	fmt.Println("Partion Mountpoint:", partition.Mountpoint)
	fmt.Println("Partion Filesystem type:", partition.Fstype)
	usage, err := disk.Usage(partition.Device)
	if err != nil {
		fmt.Println("Error occured, please try again.")
	}
	fmt.Println("Total space:", utils.FormatBytes(usage.Total)) 
	fmt.Println("Used space:", utils.FormatBytes(usage.Used)) 
	fmt.Println("Free space:", utils.FormatBytes(usage.Free))
	fmt.Printf("Space used percent: %.2f%%\n", usage.UsedPercent) 
	fmt.Println("") // blank line 
}
