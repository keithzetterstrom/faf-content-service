package models

import "github.com/keithzetterstrom/faf-content-service/internal/pkg/models"

func ModelToData(data *models.Data) Data {
	return Data{
		ID: data.ID,
		A:  data.A,
	}
}
