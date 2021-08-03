package presenter

import (
	"data"
)

type SystemInformationPresenter struct{}

var repository = data.NewSystemRepository()

func NewSystemInformationPresenter() SystemInformationPresenter {
	return SystemInformationPresenter{}
}

func (presenter SystemInformationPresenter) MemorySize() (MemorySizePresenterModel, error) {
	if size, err := repository.MemorySize(); err != nil {
		return MemorySizePresenterModel{}, err
	} else {
		return MemorySizePresenterModel{size.Size}, nil
	}
}

func (presenter SystemInformationPresenter) CpuCount() (CpuCountPresenterModel, error) {
	if count, err := repository.CpuCount(); err != nil {
		return CpuCountPresenterModel{}, err
	} else {
		return CpuCountPresenterModel{count.CpuCount}, nil
	}
}

func (presenter SystemInformationPresenter) GpuCount() (GpuCountPresenterModel, error) {
	if count, err := repository.GpuCount(); err != nil {
		return GpuCountPresenterModel{}, err
	} else {
		return GpuCountPresenterModel{count.GpuCount}, nil
	}
}

func (presenter SystemInformationPresenter) DiskPartition() (DiskPartitionPresenterModel, error) {
	if partitions, err := repository.DiskPartition(); err != nil {
		return DiskPartitionPresenterModel{}, err
	} else {
		partitionItems := make([]DiskPartitionItemPresenterModel, len(partitions))

		for i, partition := range partitions {
			partitionItems[i] = mapDiskPartitionDomainToPresenterModel(partition)
		}

		return DiskPartitionPresenterModel{Partitions: partitionItems}, nil
	}
}
