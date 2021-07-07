package storage

import (
	"tmp/servicea"
	"tmp/serviceb"
	"tmp/servicec"
)

// Provider предоставляет функционал получения специализированных сервисных хранилищ.
// Используется там, где вызываются конструкторы сервисов.
type Provider interface {
	ServiceAStorage() servicea.CommonStorager
	ServiceBStorage() serviceb.CommonStorager
	ServiceCStorage() servicec.CommonStorager
}

var _ Provider = &provider{}

// provider является реализацией Provider.
type provider struct {
	storage commonStorage
}

func (p *provider) ServiceAStorage() servicea.CommonStorager {
	return &serviceAStorage{commonStorage: p.storage}
}

type serviceAStorage struct{ commonStorage }

func (s *serviceAStorage) WithConsistency() servicea.ConsistentStorager {
	return s.commonStorage.WithConsistency()
}

type serviceBStorage struct{ commonStorage }

func (s *serviceBStorage) WithConsistency() serviceb.ConsistentStorager {
	return s.commonStorage.WithConsistency()
}

func (p *provider) ServiceBStorage() serviceb.CommonStorager {
	return &serviceBStorage{commonStorage: p.storage}
}

type serviceCStorage struct{ commonStorage }

func (s *serviceCStorage) WithConsistency() servicec.ConsistentStorager {
	return s.commonStorage.WithConsistency()
}

func (p *provider) ServiceCStorage() servicec.CommonStorager {
	return &serviceCStorage{commonStorage: p.storage}
}
