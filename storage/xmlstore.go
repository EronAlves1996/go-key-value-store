package storage

import (
	"encoding/xml"
	"fmt"
	"os"
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
	inMemoryStorage
	file string
}

var _ Storage = &xmlStore{}

func (x *xmlStore) save() error {
	f, err := os.OpenFile(x.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := xml.NewEncoder(f)

	if err := xmlMap(x.data).MarshalXML(encoder); err != nil {
		return err
	}

	return nil
}

// Delete implements Storage.
func (x *xmlStore) Delete(key string) (string, error) {
	value, err := x.inMemoryStorage.Get(key)

	if err != nil {
		return "", err
	}

	if err := x.save(); err != nil {
		x.data[key] = value
		return "", err
	}

	return value, nil
}

// Get implements Storage.
func (x *xmlStore) Get(key string) (string, error) {
	return x.inMemoryStorage.Get(key)
}

// Set implements Storage.
func (x *xmlStore) Set(key string, value string) error {

	// Trying to make keyset atomically
	oldValue, found := x.inMemoryStorage.Get(key)

	var err error
	defer func() {
		if err == nil {
			return
		}

		if found == nil {
			x.inMemoryStorage.Delete(key)
			return
		}

		x.inMemoryStorage.Set(key, oldValue)
	}()

	err = x.inMemoryStorage.Set(key, value)
	if err != nil {
		return fmt.Errorf("unable to save key %q: %w", key, err)
	}

	err = x.save()
	if err != nil {
		return fmt.Errorf("unable to save key %q: %w", key, err)
	}

	return nil
}
