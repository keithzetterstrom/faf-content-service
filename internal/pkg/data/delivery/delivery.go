package delivery

import (
	"github.com/keithzetterstrom/faf-content-service/internal/pkg/data"
)

type dataDelivery struct {
	dataUsecase data.Usecase
}

func New(dataUsecase data.Usecase) data.Delivery {
	return &dataDelivery{
		dataUsecase: dataUsecase,
	}
}
