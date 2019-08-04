package redis

import (
	"github.com/hiromaily/golibs/db/redis"
	"github.com/hiromaily/golibs/web/server/models"
)

type creativeDetailRepository struct {
	rd *redis.RD
}

func NewCreativeRepository(rd *redis.RD) models.CreativeDetailRepository {
	return &creativeDetailRepository{
		rd: rd,
	}
}

func (c *creativeDetailRepository) Fetch(creativeID string) (*models.CreativeDetail, error) {
	//TODO: WIP
	return nil, nil
}
