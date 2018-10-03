package timesrs_test

import "."

type storageMock struct {
	fnc func(entry *timesrs.StorageEntry) error
}

func NewStorageMock(fnc func(entry *timesrs.StorageEntry) error) *storageMock {
	return &storageMock{
		fnc: fnc,
	}
}

func (storage *storageMock) Store(entry *timesrs.StorageEntry) (error) {
	return storage.fnc(entry)
}
