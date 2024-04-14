package repository

import "github.com/jmoiron/sqlx"

type StoreAddressesRepo struct {
	db *sqlx.DB
}

func NewStoreAddressesRepo(db *sqlx.DB) *StoreAddressesRepo {
	return &StoreAddressesRepo{db: db}
}

func (r *StoreAddressesRepo) AddStoreAddress(address string) error {

	query := `INSERT INTO store_addresses (address) values ($1)`
	_, err := r.db.Exec(query, address)

	if err != nil {
		return err
	}

	return nil
}
