package config

import (
	"golang-transaction-experiment/repositories"
	"gorm.io/gorm"
)

type TransactionProvider interface {
	Transact(txFunc func(adapters *repositories.Adapters) error) error
}

type GenericTxProvider[T any] struct {
	db          *gorm.DB
	repoBuilder func(tx *gorm.DB) T
}

func NewGenericTxProvider[T any](db *gorm.DB, repoBuilder func(tx *gorm.DB) T) *GenericTxProvider[T] {
	return &GenericTxProvider[T]{
		db:          db,
		repoBuilder: repoBuilder,
	}
}

func (p *GenericTxProvider[T]) Transact(txFunc func(repos T) error) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		repos := p.repoBuilder(tx)
		return txFunc(repos)
	})
}
