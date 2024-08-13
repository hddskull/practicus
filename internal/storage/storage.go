package storage

type Storage interface {
	UpdateMetric(name, mType, value string) error
}
