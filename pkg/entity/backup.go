package entity

import "context"

type Backup = E

type FilterBackup = Filter

type StoreBackup interface {
	InsertBackup(context.Context, Backup) error
	FetchBackup(context.Context, FilterBackup) (Backup, error)
	FetchManyBackup(context.Context, FilterBackup) ([]Backup, []byte, error)
	DeleteBackup(context.Context, FilterBackup) error
}
