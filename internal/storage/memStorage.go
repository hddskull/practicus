package storage

import "strconv"

type MemStorage struct {
	counters map[string]float64
	gauges   map[string]int64
}

var _ Storage = NewMemStorage()

// UpdateMetric implements Storage.
func (m MemStorage) UpdateMetric(name, mType, value string) error {
	err := m.update(name, mType, value)
	if err != nil {
		return err
	}
	return nil

}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		counters: make(map[string]float64),
		gauges:   make(map[string]int64),
	}
}

func (m MemStorage) update(name, mType, value string) error {
	if mType == "counter" {
		c, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}

		m.counters[name] = c

	} else {
		g, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		prev := m.gauges[name]
		m.gauges[name] = prev + g
	}

	return nil
}
