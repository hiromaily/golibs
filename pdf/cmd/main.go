package main

import (
	"os"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/hiromaily/golibs/pdf"
)

func main() {

	// Create PDF from URL
	createPDFWithURL()

	// Create PDF from file
	createPDFWithFile()

	// version2 for json
	//pdf.NewPDFGeneratorFromJSON()

	// version3 for json
	//pdf.NewPDFGeneratorFromJSON2()
}

func createConfig() *pdf.ConfigPDF {
	config := pdf.ConfigPDF{
		PageSize:    "",
		Orientation: wk.OrientationPortrait,
		Dpi:         300,
		Grayscale:   true,
	}

	return &config
}

func createPDFWithURL() {
	// Create PDF from URL
	url := "https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf"
	output := "./output/godoc.pdf"

	cr := pdf.CreateFromURL{URL: url, ConfigPDF: createConfig()}
	pdf.NewPDFGenerator(&cr, output)
}

func createPDFWithFile() {
	// Create PDF from file
	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/pdf/cmd/"
	//filePath := "testfiles/html5.html"
	filePath := "testfiles/tables/index.html"
	output2 := "./output/html5.pdf"

	cr2 := pdf.CreateFromFile{FilePath: path + filePath, ConfigPDF: createConfig()}
	pdf.NewPDFGenerator(&cr2, output2)
}
