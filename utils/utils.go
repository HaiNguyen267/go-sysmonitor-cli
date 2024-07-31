package utils

import (
	"fmt"
	"strings"
	"math"
	"time"
)

func PrintHeading(content string) {
	var totalLength int = 60

	if len(content) >= totalLength {
		totalLength = len(content) + 10
	}

	var dashLenth = totalLength - len(content)
	var prefixDashLength int = dashLenth / 2
	var suffixLength int = dashLenth - prefixDashLength

	var prefixDash string = strings.Repeat("-", prefixDashLength)
	var suffixDash string = strings.Repeat("-", suffixLength)
	fmt.Printf("\n%v | %v | %v\n", prefixDash, content, suffixDash)

}



func ConvertEpochSecondToDateTime(secondsSinceEpoch int64) string {
	t := time.Unix(secondsSinceEpoch, 0)
	return t.Format(time.RFC3339)
}

func FormatTimeFromSeconds(seconds int64) string {
	if (seconds < 60) {
		return fmt.Sprintf("%v seconds", seconds)
	} else if (seconds < 3600) {
		minutes := int64(math.Ceil(float64(seconds / 60)))
		secondLeft := seconds - (minutes * 60)
		return fmt.Sprintf("%v minute(s) %v second(s)", minutes, secondLeft)
	} else {
		hours := int64(math.Ceil(float64(seconds / 3600)))
		minuteLeft := int64(math.Ceil(float64((seconds - hours * 3600) / 60)))
		secondLeft := seconds - (hours * 3600) - (minuteLeft * 60)
		return fmt.Sprintf("%v hour(s) %v minute(s) %v second(s)", hours, minuteLeft, secondLeft)
	}
}


func FormatBytes(byteNum uint64) string {

	var byteNumFloat float64 = float64(byteNum)
	var oneKB float64 = 1024
	var oneMB float64 = oneKB * 1024
	var oneGB float64 = oneMB * 1024
	var oneTB float64 = oneGB * 1024

	if (byteNumFloat < oneKB) {
		return fmt.Sprintf("%d bytes", byteNum)
	} else if (byteNumFloat < oneMB) {
		return fmt.Sprintf("%.2f KB", byteNumFloat / (oneKB))
	} else if (byteNumFloat < oneGB) {
		return fmt.Sprintf("%.2f MB", byteNumFloat / (oneMB))
	} else if (byteNumFloat < oneTB) {
		return fmt.Sprintf("%.2f GB", byteNumFloat / (oneGB))
	} else {
		return fmt.Sprintf("%.2f TB", byteNumFloat / (oneTB))
	}
}


