package pdf

import (
	"os"
	"testing"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func createConfig() *ConfigPDF {
	config := ConfigPDF{
		PageSize:    "",
		Orientation: wk.OrientationPortrait,
		Dpi:         300,
		Grayscale:   true,
	}

	return &config
}

func createPDFFromURL(url, output string, config *ConfigPDF) {
	cr := CreateFromURL{
		URL:       url,
		ConfigPDF: config,
	}
	NewPDFGenerator(&cr, output)
}

func createPDFFromFile(filePath, output string, config *ConfigPDF) {
	cr := CreateFromFile{
		FilePath:  filePath,
		ConfigPDF: config,
	}
	NewPDFGenerator(&cr, output)
}

// Though prefix is Test, actually it's just Example
func TestNewPDFGenerator_FromURL(t *testing.T) {
	t.SkipNow()
	//2.35s
	url := "https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf"
	output := "from_url.pdf"

	createPDFFromURL(url, output, createConfig())
}

// Though prefix is Test, actually it's just Example
func TestNewPDFGenerator_FromFile(t *testing.T) {
	t.SkipNow()
	//1.47s
	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/pdf/cmd/"
	filePath := path + "testfiles/tables/index.html"
	output := "from_file.pdf"

	createPDFFromFile(filePath, output, createConfig())
}
