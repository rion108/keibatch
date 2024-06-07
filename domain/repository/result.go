package repository

type ResultRepository interface {
	GetResultIterator(filePath string) (ResultIterator, error)
	GetFilePaths() []string
}
