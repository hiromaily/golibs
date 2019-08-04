package builder

import (
	"fmt"
	"io"
	"net/url"
	"strings"
	"text/template"

	"github.com/hiromaily/golibs/web/server/models"
)

type HTMLBuilder interface {
	Build(w io.Writer, creative *models.CreativeDetail, datString string) error
}

type htmlBuilder struct {
	random   string
	clickURL string
}

type htmlAdMarkupData struct {
	DivID string
	HTML  string
}

func NewHTMLBuilder(random, clickURL string) HTMLBuilder {
	return &htmlBuilder{
		random:   random,
		clickURL: clickURL,
	}
}

var (
	htmlTemplateText = `<div id="{{.DivID}}">
{{.HTML}}
</div>`
	htmlTemplate = template.Must(template.New("html").Parse(htmlTemplateText))
)

func (h *htmlBuilder) Build(w io.Writer, creative *models.CreativeDetail, datString string) error {
	lpURL := url.QueryEscape(creative.ClickURL)
	clickURL := fmt.Sprintf("%s/?dat=%s&lp=%s", h.clickURL, datString, lpURL)
	clickURLEsc := url.QueryEscape(clickURL)

	macros := map[string]string{
		"{CLICK_URL}":     clickURL,
		"{CLICK_URL_ESC}": clickURLEsc,
		"{RANDOM}":        h.random,
	}
	html := creative.HTML
	for k, v := range macros {
		html = strings.Replace(html, k, v, -1)
	}

	data := htmlAdMarkupData{
		DivID: h.random,
		HTML:  html,
	}

	if err := htmlTemplate.Execute(w, data); err != nil {
		return err
	}

	return nil
}
