package data

import (
	"domain"
)

type SystemDataRepository struct{}

func NewSystemRepository() SystemDataRepository {
	return SystemDataRepository{}
}

func (repository SystemDataRepository) GpuCount() (domain.GpuCountModel, error) {
	if count, err := GpuCount(); err != nil {
		return domain.GpuCountModel{}, err
	} else {
		return domain.GpuCountModel{GpuCount: count}, nil
	}
}

func (repository SystemDataRepository) CpuCount() (domain.CpuCountModel, error) {
	if count, err := CpuCount(); err != nil {
		return domain.CpuCountModel{}, err
	} else {
		return domain.CpuCountModel{CpuCount: count}, nil
	}
}

func (repository SystemDataRepository) MemorySize() (domain.MemorySizeModel, error) {
	if out, err := MemorySize(); err != nil {
		return domain.MemorySizeModel{}, err
	} else {
		return domain.MemorySizeModel{Size: out}, nil
	}
}

func (repository SystemDataRepository) DiskPartition() (domain.DiskPartitionModel, error) {
	if out, err := DiskPartition(); err != nil {
		return domain.DiskPartitionModel{}, err
	} else {
		return out.ToDomain(), nil
	}
}
