package provider

import (
	"tmp/servicea"
	"tmp/serviceb"
	"tmp/servicec"
	"tmp/storage"
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
	storage storage.CommonStorage
}

func (p *provider) ServiceAStorage() servicea.CommonStorager {
	return &serviceAStorage{CommonStorage: p.storage}
}

type serviceAStorage struct{ storage.CommonStorage }

func (s *serviceAStorage) WithConsistency() servicea.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

type serviceBStorage struct{ storage.CommonStorage }

func (s *serviceBStorage) WithConsistency() serviceb.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

func (p *provider) ServiceBStorage() serviceb.CommonStorager {
	return &serviceBStorage{CommonStorage: p.storage}
}

type serviceCStorage struct{ storage.CommonStorage }

func (s *serviceCStorage) WithConsistency() servicec.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

func (p *provider) ServiceCStorage() servicec.CommonStorager {
	return &serviceCStorage{CommonStorage: p.storage}
}

type Consistencer interface {
	Commit()
	Rollback()
}
