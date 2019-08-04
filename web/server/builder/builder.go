package builder

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/hiromaily/golibs/web/server/models"
)

type Builder interface {
	Build(w io.Writer, password string, creative *models.CreativeDetail) error
}

type builder struct {
	bannerBuilder BannerBuilder //TODO: ここはinterface型じゃなくてもstructでもいいような, DI?
	htmlBuilder   HTMLBuilder
}

func NewBuilder(bannerBuilder BannerBuilder, htmlBuilder HTMLBuilder) Builder {
	// create builder struct
	return &builder{
		bannerBuilder: bannerBuilder,
		htmlBuilder:   htmlBuilder,
	}
}

func (b *builder) Build(w io.Writer, datString string, creative *models.CreativeDetail) error {
	if creative == nil {
		return errors.New("creative is nil")
	}

	switch creative.Type {
	case models.BannerCreativeType:
		return b.bannerBuilder.Build(w, creative, datString)
	case models.HTMLCreativeType:
		return b.htmlBuilder.Build(w, creative, datString)
	default:
		return fmt.Errorf("invalid creative type %v", creative.Type)
	}
}
