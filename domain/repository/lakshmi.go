package repository

import (
	"metis/domain/model"
	"metis/domain/types"
	"metis/infrastructure/dto"
)

type LakshmiRepository interface {
	Transaction(f func(lakshmiRepository LakshmiRepository) error) error
	GetPartialHorses(horseNames types.HorseNames) (map[types.HorseName]types.HorseID, error)
	UpsertHorses(horses model.Horses) error
	UpsertRaces(races model.Races) error
	CreateRaceMapping(results dto.ResultDTOs, horseMaps map[types.HorseName]types.HorseID) error
	GetRaceHorseMappings(results dto.ResultDTOs) (dto.RaceHorseMappingDTOs, error)
	CreateRaceResult(results dto.ResultDTOs, raceHorseMappings dto.RaceHorseMappingDTOs) error
}
