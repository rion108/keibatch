package repository

import "metis/infrastructure/dto"

type ResultIterator interface {
	HasNext() bool
	Next() (dto.ResultDTOs, error)
}
