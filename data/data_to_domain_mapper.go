package data

import "domain"

const SizeKey = "SIZE"
const NameKey = "NAME"
const FSTypeKey = "FSTYPE"
const TypeKey = "TYPE"
const MountKey = "MOUNTPOINT"
const ModelKey = "MODEL"
const StateKey = "STATE"

func (dataModel DiskPartitionDataModel) ToDomain() domain.DiskPartitionModel {
	return domain.DiskPartitionModel{
		Name:           dataModel[NameKey],
		Size:           dataModel[SizeKey],
		FileSystemType: dataModel[FSTypeKey],
		PartitionType:  dataModel[TypeKey],
		MountPoint:     dataModel[MountKey],
		Model:          dataModel[ModelKey],
		State:          dataModel[StateKey],
	}
}
