package main

import "fmt"

type InvalidKeyError struct {
	Key string
}

func (i *InvalidKeyError) Error() string {
	return fmt.Sprintf("invalid key: %v", i.Key)
}
