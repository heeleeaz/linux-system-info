package data

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiskPartition(t *testing.T) {
	cmdout :=
		"NAME=\"sda\" SIZE=\"60G\" FSTYPE=\"\" TYPE=\"disk\" MOUNTPOINT=\"\" MODEL=\"PersistentDisk  \" STATE=\"running\"" +
			"\nNAME=\"sda1\" SIZE=\"512M\" FSTYPE=\"ext3\" TYPE=\"part\" MOUNTPOINT=\"/boot\" MODEL=\"\" STATE=\"\""

	expected := []DiskPartitionDataModel{
		map[string]string{
			"NAME": "sda", "SIZE": "60G", "FSTYPE": "", "TYPE": "disk", "MOUNTPOINT": "", "MODEL": "PersistentDisk",
			"STATE": "running",
		},
		map[string]string{
			"FSTYPE": "ext3", "MODEL": "", "NAME": "sda1", "SIZE": "512M", "STATE": "", "TYPE": "part",
		},
	}

	commandExecutor = MockCommandExecutor{cmdout, nil}
	if output, err := DiskPartition(); err != nil {
		if fmt.Sprintf("%T", err) != fmt.Sprintf("%T", expected) {
			t.Errorf("expected: %T, actual: %T", expected, err)
		}
	} else {
		if !reflect.DeepEqual(expected, output) {
			t.Errorf("expected: %s\nactual: %s", expected, output)
		}
	}
}
