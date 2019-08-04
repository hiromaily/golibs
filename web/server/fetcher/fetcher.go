package fetcher

import (
	"github.com/pkg/errors"

	"github.com/hiromaily/golibs/web/server/models"
)

type Fetcher interface {
	Fetch() (*models.CreativeDetail, error)
}

// ----------------------------------------------------------------------------
// Dat
// ----------------------------------------------------------------------------

type datFetcher struct {
	dat          string
	creativeRepo models.CreativeDetailRepository
}

func NewDatFetcher(dat string, creativeRepo models.CreativeDetailRepository) Fetcher {
	return &datFetcher{
		dat:          dat,
		creativeRepo: creativeRepo,
	}
}

func (f *datFetcher) Fetch() (*models.CreativeDetail, error) {

	creative, err := f.creativeRepo.Fetch(f.dat)
	if err != nil {
		return nil, errors.Errorf("can't fetch creative: %s", f.dat)
	}

	return creative, nil
}

// ----------------------------------------------------------------------------
// Rad
// ----------------------------------------------------------------------------

type radFetcher struct {
	rad          string
	creativeRepo models.CreativeDetailRepository
}

func NewRadFetcher(rad string, creativeRepo models.CreativeDetailRepository) Fetcher {
	return &radFetcher{
		rad:          rad,
		creativeRepo: creativeRepo,
	}
}

func (f *radFetcher) Fetch() (*models.CreativeDetail, error) {

	creative, err := f.creativeRepo.Fetch(f.rad)
	if err != nil {
		return nil, errors.Errorf("can't fetch creative: %s", f.rad)
	}

	return creative, nil
}
