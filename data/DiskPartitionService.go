package data

import (
	"regexp"
	"strings"
)

const FetchDiskPartitionCommand = "lsblk -P -o NAME,SIZE,FSTYPE,TYPE,MOUNTPOINT,MODEL,STATE"

const DiskPartitionKeyValueItemSplitRegex = `(?:\w{1,20}="(?:\S*\s*)+")`

type DiskPartitionDataModel map[string]string

func DiskPartition() (DiskPartitionDataModel, error) {
	if out, err := executeCommand(FetchDiskPartitionCommand); err != nil {
		return nil, err
	} else {
		return formatDiskPartition(out)
	}
}

func formatDiskPartition(output string) (DiskPartitionDataModel, *FormatOutputError) {
	lineSplit := strings.Split(output, "\n")
	regexCompile := regexp.MustCompile(DiskPartitionKeyValueItemSplitRegex)

	partitionModels := make(DiskPartitionDataModel)
	for _, line := range lineSplit {
		keyValue := regexCompile.Split(line, -1)
		partitionModels[keyValue[0]] = keyValue[1]
	}

	return partitionModels, nil
}
