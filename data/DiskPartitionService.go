package data

import (
	"fmt"
	"regexp"
	"strings"
)

const FetchDiskPartitionCommand = "lsblk -P -o NAME,SIZE,FSTYPE,TYPE,MOUNTPOINT,MODEL,STATE"

const DiskPartitionKeyValueItemSplitRegex = "\\w{1,20}=\"\\w*\\W*\""

type DiskPartitionDataModel map[string]string

func DiskPartition() ([]DiskPartitionDataModel, error) {
	if out, err := executeCommand(FetchDiskPartitionCommand); err != nil {
		return nil, err
	} else {
		return formatDiskPartition(out)
	}
}

func formatDiskPartition(output string) ([]DiskPartitionDataModel, error) {
	partitions := strings.Split(output, "\n")

	if len(partitions) <= 0 {
		return nil, &FormatOutputError{fmt.Sprintf("error formatting: %s", output), nil}
	}

	regexCompile := regexp.MustCompile(DiskPartitionKeyValueItemSplitRegex)

	partitionModel := make([]DiskPartitionDataModel, len(partitions))
	for i, partition := range partitions {
		partitionInformation := regexCompile.FindAllStringSubmatch(partition, -1)

		partitionInformationItem := make(DiskPartitionDataModel)
		for _, equalSeperated := range partitionInformation {
			if kv, _ := keyValueFromEqualSeperatedString(equalSeperated[0]); kv != nil {
				partitionInformationItem[kv[0]] = kv[1]
			}
		}

		partitionModel[i] = partitionInformationItem
	}

	return partitionModel, nil
}

func keyValueFromEqualSeperatedString(input string) ([]string, error) {
	keyValue := strings.Split(input, "=")

	if len(keyValue) == 2 {
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(strings.ReplaceAll(keyValue[1], "\"", ""))

		return []string{key, value}, nil
	}

	return nil, &FormatOutputError{fmt.Sprintf("error formatting: %s", input), nil}
}
