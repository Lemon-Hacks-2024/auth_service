package service

import "auth_service/internal/repository"

type StoreAddressesService struct {
	storeAddressesRepo repository.StoreAddresses
}

func NewStoreAddressesService(storeAddressesRepo repository.StoreAddresses) *StoreAddressesService {
	return &StoreAddressesService{storeAddressesRepo: storeAddressesRepo}
}

func (s *StoreAddressesService) AddStoreAddress(address string) error {

	return s.storeAddressesRepo.AddStoreAddress(address)
}
