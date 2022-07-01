package main

import (
	"fmt"
)

type StorageType int

type Storage interface {
	Store(data string)
}

type MemoryStorage struct{}

type S3Storage struct{}

const (
	Memory StorageType = iota
	S3
)

func (s *MemoryStorage) Store(data string) {
	fmt.Printf("store data %s to memory\n", data)
}

func (s *S3Storage) Store(data string) {
	fmt.Printf("store data %s to S3\n", data)
}

func NewStorage(storageType StorageType) Storage {
	switch storageType {
	case Memory:
		return &MemoryStorage{}
	case S3:
		return &S3Storage{}
	default:
		panic("this storage type is not supported")
	}
}

func main() {
	data := "my password"

	s1 := NewStorage(S3)
	s1.Store(data)

	s2 := NewStorage(Memory)
	s2.Store(data)
}
