package presenter

import "domain"

func mapDiskPartitionDomainToPresenterModel(input domain.DiskPartitionModel) DiskPartitionItemPresenterModel {
	return DiskPartitionItemPresenterModel{
		Name:           input.Name,
		Size:           input.Size,
		FileSystemType: input.FileSystemType,
		PartitionType:  input.PartitionType,
		MountPoint:     input.MountPoint,
		Model:          input.Model,
		State:          input.State,
	}
}
