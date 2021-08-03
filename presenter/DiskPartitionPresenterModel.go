package presenter

type DiskPartitionPresenterModel struct {
	Partitions []DiskPartitionItemPresenterModel `json:"partitions"`
}

type DiskPartitionItemPresenterModel struct {
	Name           string `json:"name"`
	Size           string `json:"size"`
	FileSystemType string `json:"file_system_type"`
	PartitionType  string `json:"partition_type"`
	MountPoint     string `json:"mount_point"`
	Model          string `json:"model"`
	State          string `json:"state"`
}
