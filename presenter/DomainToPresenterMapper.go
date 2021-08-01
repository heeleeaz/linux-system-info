package presenter

import "domain"

func mapDiskPartitionDomainToPresenterModel(input domain.DiskPartitionModel) DiskPartitionPresenterModel {
	return DiskPartitionPresenterModel{
		Name:           input.Name,
		Size:           input.Size,
		FileSystemType: input.FileSystemType,
		PartitionType:  input.PartitionType,
		MountPoint:     input.MountPoint,
		Model:          input.Model,
		State:          input.State,
	}
}
