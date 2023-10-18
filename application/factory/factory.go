package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/otaviopontes/codepix-go/application/usecase"
	"github.com/otaviopontes/codepix-go/infra/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	return usecase.TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixRepository:         pixRepository,
	}

}
