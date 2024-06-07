package repositoryImpl

import (
	"bufio"
	"io"
	"metis/domain/repository"
	"metis/infrastructure/dto"
	"strings"
)

type resultIteratorImpl struct {
	readCloser io.ReadCloser
	scanner    *bufio.Scanner
	csvHeader  string
	limit      int
	hasNext    bool
}

func NewResultIteratorImpl(readerCloser io.ReadCloser, limit int) repository.ResultIterator {
	scanner := bufio.NewScanner(readerCloser)
	defer readerCloser.Close()

	return &resultIteratorImpl{
		readCloser: readerCloser,
		scanner:    scanner,
		csvHeader:  scanner.Text(),
		limit:      limit,
		hasNext:    scanner.Scan(),
	}
}

func (r *resultIteratorImpl) HasNext() bool {
	return r.hasNext
}

func (r *resultIteratorImpl) Next() (dto.ResultDTOs, error) {
	var rows strings.Builder
	rows.WriteString(r.csvHeader)
	for i := 0; i < r.limit; i++ {
		if !r.HasNext() {
			break
		}
		rows.WriteString("\n")
		rows.WriteString(r.nextRow())
	}
	return dto.NewResultDTOs(rows.String())
}

func (r *resultIteratorImpl) nextRow() string {
	row := r.scanner.Text()
	r.hasNext = r.scanner.Scan()
	return row
}
