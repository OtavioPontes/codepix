package usecase

import (
	"errors"

	"github.com/otaviopontes/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}
	pixKey, err := model.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}
	pixKeyRegistered, err := p.PixKeyRepository.RegisterKey(pixKey)
	if err != nil {
		return nil, err
	}
	if pixKeyRegistered.ID == "" {
		return nil, errors.New("unable to register key at the moment")
	}
	return pixKeyRegistered, err
}

func (p *PixUseCase) FindKey(key, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
