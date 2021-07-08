package provider

import (
	"tmp/servicea"
	"tmp/serviceb"
	"tmp/servicec"
	"tmp/storage"
)

func New(storage *storage.CommonStorage) *Provider {
	return &Provider{storage: storage}
}

// Provider предоставляет функционал получения специализированных сервисных хранилищ.
// Используется там, где вызываются конструкторы сервисов.
type Provider struct {
	storage *storage.CommonStorage
}

func (p *Provider) ServiceAStorage() servicea.CommonStorager {
	return &serviceAStorage{CommonStorage: p.storage}
}

type serviceAStorage struct{ *storage.CommonStorage }

func (s *serviceAStorage) WithConsistency() servicea.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

type serviceBStorage struct{ *storage.CommonStorage }

func (s *serviceBStorage) WithConsistency() serviceb.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

func (p *Provider) ServiceBStorage() serviceb.CommonStorager {
	return &serviceBStorage{CommonStorage: p.storage}
}

type serviceCStorage struct{ *storage.CommonStorage }

func (s *serviceCStorage) WithConsistency() servicec.ConsistentStorager {
	return s.CommonStorage.WithConsistency()
}

func (p *Provider) ServiceCStorage() servicec.CommonStorager {
	return &serviceCStorage{CommonStorage: p.storage}
}
