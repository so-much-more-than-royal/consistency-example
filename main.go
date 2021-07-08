package main

import (
	"tmp/serviceb"
	"tmp/storage"
)

func main() {
	commonStorage := &storage.CommonStorage{}

	serviceB := serviceb.New(
		serviceb.NewCommonStorage(commonStorage),
	)

	serviceB.DoSome(1)
}
