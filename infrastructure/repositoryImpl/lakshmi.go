package repositoryImpl

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"metis/domain/model"
	"metis/domain/repository"
	"metis/domain/types"
	"metis/infrastructure/dto"
)

type lakshmiRepositoryImpl struct {
	db *gorm.DB
}

func NewLakshmiRepositoryImpl(db *gorm.DB) repository.LakshmiRepository {
	return &lakshmiRepositoryImpl{db: db}
}

func (r *lakshmiRepositoryImpl) Transaction(f func(lakshmiRepository repository.LakshmiRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return f(&lakshmiRepositoryImpl{db: tx})
	})
}
func (r *lakshmiRepositoryImpl) UpsertHorses(horses model.Horses) error {
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(dto.NewHorses(horses))
	return result.Error
}

func (r *lakshmiRepositoryImpl) GetPartialHorses(horseNames types.HorseNames) (map[types.HorseName]types.HorseID, error) {
	var partialHorses dto.PartialHorseDTOs
	err := r.db.
		Select("id,name").
		Where("name IN ?", horseNames).
		Find(&partialHorses).Error
	if err != nil {
		return nil, err
	}
	return partialHorses.ToPartialHorseMaps(), nil
}

// UpsertRaces レース情報を更新する
func (r *lakshmiRepositoryImpl) UpsertRaces(races model.Races) error {
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(dto.NewRaces(races))
	return result.Error
}

func (r *lakshmiRepositoryImpl) CreateRaceMapping(results dto.ResultDTOs, horseMaps map[types.HorseName]types.HorseID) error {
	raceMappings := dto.NewRaceMappingDTOs(results, horseMaps)
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&raceMappings)
	return result.Error
}

func (r *lakshmiRepositoryImpl) GetRaceHorseMappings(results dto.ResultDTOs) (dto.RaceHorseMappingDTOs, error) {
	var raceHorseMappings dto.RaceHorseMappingDTOs
	//レコード分selectが走るので、処理は見直しが必要
	for _, result := range results {
		var raceHorseMapping *dto.RaceHorseMappingDto
		err := r.db.Table("race_mappings").Select("race_mappings.id as id , horse_id, race_id, horses.name as horse_name").
			Joins("INNER JOIN horses ON race_mappings.horse_id = horses.id").
			Where("race_id = ? AND horses.name = ?", result.ID, result.HorseName).
			Find(&raceHorseMapping).Error
		if err != nil {
			return nil, err
		}
		raceHorseMappings = append(raceHorseMappings, raceHorseMapping)
	}
	return raceHorseMappings, nil
}

func (r *lakshmiRepositoryImpl) CreateRaceResult(results dto.ResultDTOs, raceHorseMappings dto.RaceHorseMappingDTOs) error {
	result := r.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(dto.NewRaceResultsDTOs(results, raceHorseMappings))
	return result.Error
}
