package dto

import (
	"metis/domain/types"
	"time"
)

type RaceMappingDto struct {
	ID        int           `gorm:"autoIncrement column:id"`
	HorseID   types.HorseID `gorm:"column:horse_id"`
	RaceID    types.RaceID  `gorm:"column:race_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RaceMappingDTOs []RaceMappingDto

func NewRaceMappingDTOs(results ResultDTOs, horseMaps map[types.HorseName]types.HorseID) RaceMappingDTOs {
	var mappings RaceMappingDTOs
	for _, result := range results {
		mapping := NewRaceMappingDTO(result, horseMaps)
		mappings = append(mappings, mapping)
	}

	return mappings
}

func NewRaceMappingDTO(result *ResultDto, horseMaps map[types.HorseName]types.HorseID) RaceMappingDto {
	return RaceMappingDto{
		HorseID: horseMaps[result.HorseName],
		RaceID:  result.ID,
	}
}

func (*RaceMappingDto) TableName() string {
	return "race_mappings"
}
