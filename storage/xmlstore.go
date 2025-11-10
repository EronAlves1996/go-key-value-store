package storage

import (
	"encoding/xml"
	"errors"
	"os"
	"sync"
)

type xmlMap map[string]string

func (m xmlMap) MarshalXML(e *xml.Encoder) error {
	start := xml.StartElement{
		Name: xml.Name{
			Local: "map",
		},
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		name := xml.StartElement{
			Name: xml.Name{Local: k},
		}

		err := e.EncodeElement(v, name)
		if err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

type xmlStore struct {
	mu   *sync.RWMutex
	data xmlMap
	file string
}

var _ Storage = &xmlStore{}

func (x *xmlStore) lock() func() {
	x.mu.Lock()
	return func() {
		x.mu.Unlock()
	}
}

func (x *xmlStore) rlock() func() {
	x.mu.RLock()
	return func() {
		x.mu.RUnlock()
	}
}

func (x *xmlStore) save() error {
	f, err := os.OpenFile(x.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := xml.NewEncoder(f)

	if err := x.data.MarshalXML(encoder); err != nil {
		return err
	}

	return nil
}

// Delete implements Storage.
func (x *xmlStore) Delete(key string) (string, error) {
	defer x.lock()()
	value, err := x.internalGet(key)

	if err != nil {
		return "", err
	}

	if err := x.save(); err != nil {
		x.data[key] = value
		return "", err
	}

	return value, nil
}

func (x *xmlStore) internalGet(key string) (string, error) {
	value, ok := x.data[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Get implements Storage.
func (x *xmlStore) Get(key string) (string, error) {
	defer x.rlock()()
	return x.internalGet(key)
}

// Set implements Storage.
func (x *xmlStore) Set(key string, value string) error {
	defer x.lock()()

	// Trying to make keyset atomically
	oldValue, found := x.data[key]

	var err error
	defer func() {
		if err == nil {
			return
		}

		if !found {
			delete(x.data, key)
			return
		}

		x.data[key] = oldValue
	}()

	x.data[key] = value

	err = x.save()
	if err != nil {
		return errors.New("unable to save key")
	}

	return nil
}
