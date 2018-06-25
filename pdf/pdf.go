package pdf

import (
	"fmt"
	"log"

	"bytes"
	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"os"
	"strings"
)

// https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf

//func setFooters(page *wk.Page) {
//	//page.FooterRight.Set("[page]")
//	//page.FooterFontSize.Set(10)
//	//page.Zoom.Set(95.50)
//}

// NewPDFGenerator is to create PDF
func NewPDFGenerator(cr Creator, outputPath string) {
	// Create new PDF generator
	pdfg := cr.NewGenerator()

	// Set global options
	cr.GetConf().SetOptions(pdfg)
	//setOptions(pdfg, cr)

	// Create a new input page form what depending on creator parameter
	cr.CreatePage(pdfg)

	// Create PDF document in internal buffer
	pdfg = cr.CreatePDFBuffer(pdfg)

	// Write buffer contents to file on disk
	err := pdfg.WriteFile(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

// NewPDFGeneratorFromJSON is only experimental
func NewPDFGeneratorFromJSON() {
	//It works

	const html = `<!doctype html><html><head><title>WKHTMLTOPDF TEST</title></head><body>HELLO PDF</body></html>`

	// Client code
	pdfg := wk.NewPDFPreparer()
	pdfg.AddPage(wk.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// The JSON can be saved, uploaded, etc.

	// Server code, create a new PDF generator from JSON, also looks for the wkhtmltopdf executable
	pdfgFromJSON, err := wk.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Use the PDF
	fmt.Printf("PDF size %d bytes\n", pdfgFromJSON.Buffer().Len())

	//
	err = pdfgFromJSON.WriteFile("./simplesample3.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")

}

// NewPDFGeneratorFromJSON2 is only experimental
func NewPDFGeneratorFromJSON2() {
	//It works

	// Client code
	pdfg := wk.NewPDFPreparer()

	// add a reader page as well
	//TODO:file path should be absolute path
	file, err := os.Open("./testfiles/html5.html")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	pdfg.AddPage(wk.NewPageReader(file))

	pdfg.Dpi.Set(600)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// The JSON can be saved, uploaded, etc.

	// Server code, create a new PDF generator from JSON, also looks for the wkhtmltopdf executable
	pdfgFromJSON, err := wk.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Use the PDF
	fmt.Printf("PDF size %d bytes\n", pdfgFromJSON.Buffer().Len())

	//
	err = pdfgFromJSON.WriteFile("./simplesample4.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")

}
