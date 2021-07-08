package main

import (
	"tmp/serviceb"
	"tmp/storage"
)

func main() {
	stg := &storage.CommonStorage{}

	svb := serviceb.New(
		serviceb.NewCommonStorage(stg),
	)

	svb.DoSome(1)
}
