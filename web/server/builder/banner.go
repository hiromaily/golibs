package builder

import (
	"html/template"
	"io"

	"github.com/hiromaily/golibs/web/server/models"
)

type BannerBuilder interface {
	Build(w io.Writer, creative *models.CreativeDetail, datString string) error
}

type bannerBuilder struct {
	random   string
	clickURL string
}

type bannerData struct {
	ClickURL string
	DivID    string
	Dat      string
	ImgURL   string
}

func NewBannerBuilder(random, clickURL string) BannerBuilder {
	// return bannerBuilder struct
	return &bannerBuilder{
		random:   random,
		clickURL: clickURL,
	}
}

var (
	bannerTemplateText = `<div id="{{.DivID}}">
<a target="_blank" href="{{.ClickURL}}?dat={{.Dat}}">
<img src="{{.ImgURL}}"></img>
</a>
</div>`
	bannerTemplate = template.Must(template.New("banner").Parse(bannerTemplateText))
)

func (b *bannerBuilder) Build(w io.Writer, creative *models.CreativeDetail, datString string) error {
	data := bannerData{
		ClickURL: b.clickURL,
		DivID:    b.random,
		Dat:      datString,
		ImgURL:   creative.FilePath,
	}

	if err := bannerTemplate.Execute(w, data); err != nil {
		return err
	}

	return nil
}
