package domain

type SystemRepository interface {
	GpuCount() (GpuCountModel, SystemError)
	CpuCount() (CpuCountModel, SystemError)
	MemorySize() (MemorySizeModel, SystemError)
	DiskPartition() (DiskPartitionModel, SystemError)
}
