package gosubmoduleexample

import "errors"

var ErrNotFound = errors.New("not found")

type Data struct {
	ValA string
	ValB string
}

type Logger func(Data)

type Store struct {
	logger  Logger
	dataMap map[string]Data
}

func NewStore(logger Logger) *Store {
	return &Store{
		logger:  logger,
		dataMap: make(map[string]Data),
	}
}

func (c *Store) StoreData(a string, b string) error {
	newData := Data{
		ValA: a,
		ValB: b,
	}
	c.dataMap[a] = newData
	c.logger(newData)

	return nil
}

func (c *Store) RetrieveData(a string) (Data, error) {
	if data, ok := c.dataMap[a]; ok {
		return data, nil
	}

	return Data{}, ErrNotFound
}
