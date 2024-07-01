package libs

import (
	"Chat/dto"
	"time"
)

func CreateDefaultValue() dto.TDefaultValue {

	return dto.TDefaultValue{
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		IsDeleted: false,
	}
}

func UpdateDefaultValue() dto.TDefaultValue {

	return dto.TDefaultValue{
		UpdatedAt: time.Now().String(),
	}
}
