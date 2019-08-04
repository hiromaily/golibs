package db

import (
	"github.com/hiromaily/golibs/db/mysql"
	"github.com/hiromaily/golibs/web/server/models"
)

type creativeDetailRepository struct {
	db *mysql.MS
}

func NewCreativeRepository(db *mysql.MS) models.CreativeDetailRepository {
	return &creativeDetailRepository{
		db: db,
	}
}

func (c *creativeDetailRepository) Fetch(creativeID string) (*models.CreativeDetail, error) {
	//TODO: WIP
	return nil, nil
}
