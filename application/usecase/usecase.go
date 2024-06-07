package usecase

import (
	"metis/domain/repository"
)

type LakshmiUsecase interface {
	Update() error
}

type lakshmiUsecase struct {
	lakshmiRepo repository.LakshmiRepository
	resultRepo  repository.ResultRepository
}

func NewLakshmiUsecase(
	lakshmiRepo repository.LakshmiRepository,
	resultRepo repository.ResultRepository,
) LakshmiUsecase {

	return &lakshmiUsecase{
		lakshmiRepo: lakshmiRepo,
		resultRepo:  resultRepo,
	}
}

func (u *lakshmiUsecase) Update() error {

	filePaths := u.resultRepo.GetFilePaths()
	for _, filePath := range filePaths {
		err := u.lakshmiRepo.Transaction(func(tx repository.LakshmiRepository) error {

			iter, err := u.resultRepo.GetResultIterator(filePath)
			if err != nil {
				return err
			}
			for iter.HasNext() {
				results, err := iter.Next()
				if err != nil {
					return err
				}
				horses := results.ToHorses()
				err = tx.UpsertHorses(horses)
				if err != nil {
					return err
				}

				races := results.ToRaces()
				err = tx.UpsertRaces(races)
				if err != nil {
					return err
				}

				horseMaps, err := tx.GetPartialHorses(horses.GetHorseNames())
				if err != nil {
					return err
				}
				err = tx.CreateRaceMapping(results, horseMaps)
				if err != nil {
					return err
				}

				raceHorseMaps, err := tx.GetRaceHorseMappings(results)
				if err != nil {
					return err
				}
				err = tx.CreateRaceResult(results, raceHorseMaps)
				if err != nil {
					return err
				}

			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
