package dto

import (
	"metis/domain/types"
)

type PartialHorseDTO struct {
	ID   types.HorseID
	Name types.HorseName
}

type PartialHorseDTOs []PartialHorseDTO

func (ds PartialHorseDTOs) ToPartialHorseMaps() map[types.HorseName]types.HorseID {
	horseMaps := make(map[types.HorseName]types.HorseID)
	for _, d := range ds {
		horseMaps[d.Name] = d.ID
	}
	return horseMaps
}

func (*PartialHorseDTO) TableName() string {
	return "horses"
}
