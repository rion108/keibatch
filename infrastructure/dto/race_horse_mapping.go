package dto

import (
	"metis/domain/types"
)

type RaceHorseMappingDto struct {
	ID        int             `gorm:"column:id"`
	HorseID   types.HorseID   `gorm:"column:horse_id"`
	HorseName types.HorseName `gorm:"column:horse_name"`
	RaceID    types.RaceID    `gorm:"column:race_id"`
}

type RaceHorseMappingDTOs []*RaceHorseMappingDto
