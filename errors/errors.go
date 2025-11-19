package errors

import "fmt"

type InvalidKeyError struct {
	Key string
}

func (i InvalidKeyError) Error() string {
	return fmt.Sprintf("invalid key: %v", i.Key)
}

func (i *InvalidKeyError) Is(target error) bool {
	_, ok := target.(*InvalidKeyError)
	return ok
}

func (i *InvalidKeyError) As(t any) bool {
	_, ok := t.(*InvalidKeyError)
	return ok
}
