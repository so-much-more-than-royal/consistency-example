package main

import (
	"tmp/provider"
	"tmp/serviceb"
	"tmp/storage"
)

func main() {
	commonStorage := &storage.CommonStorage{}
	storageProvider := provider.New(commonStorage)
	serviceB := serviceb.New(storageProvider.ServiceBStorage())
	serviceB.DoSome(1)
}
