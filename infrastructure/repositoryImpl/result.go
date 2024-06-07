package repositoryImpl

import (
	"metis/domain/repository"
	"os"
)

type resultRepositoryImpl struct {
	filePaths []string
	limit     int
}

func NewResultRepositoryImpl(filePaths []string, limit int) repository.ResultRepository {
	return &resultRepositoryImpl{
		filePaths: filePaths,
		limit:     limit,
	}
}

func (r *resultRepositoryImpl) GetResultIterator(filePath string) (repository.ResultIterator, error) {
	readCloser, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return NewResultIteratorImpl(readCloser, r.limit), nil
}

func (r *resultRepositoryImpl) GetFilePaths() []string {
	return r.filePaths
}
