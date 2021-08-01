package domain

type DiskPartitionModel struct {
	Name           string
	Size           string
	FileSystemType string
	PartitionType  string
	MountPoint     string
	Model          string
	State          string
}
