package errors

import "fmt"

type InvalidKeyError struct {
	Key string
}

func (i *InvalidKeyError) Error() string {
	return fmt.Sprintf("invalid key: %v", i.Key)
}

func (i *InvalidKeyError) Is(target error) bool {
	converted, ok := target.(*InvalidKeyError)
	if !ok {
		return false
	}

	return converted.Key == i.Key
}

func (i *InvalidKeyError) As(t any) bool {
	_, ok := t.(*InvalidKeyError)
	return ok
}
