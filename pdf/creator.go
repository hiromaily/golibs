package pdf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	wk "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	tm "github.com/hiromaily/golibs/time"
)

// Creator is to create
type Creator interface {
	GetConf() *ConfigPDF
	NewGenerator() *wk.PDFGenerator
	CreatePage(*wk.PDFGenerator) bool
	CreatePDFBuffer(*wk.PDFGenerator) *wk.PDFGenerator
}

// CreateFromURL is URL object corresponding with Creator interface
type CreateFromURL struct {
	URL string
	*ConfigPDF
}

// ConfigPDF is configuration for PDF
type ConfigPDF struct {
	PageSize    string
	Orientation string
	Dpi         uint
	Grayscale   bool
}

// SetOptions is to set PDF configuration
func (pc *ConfigPDF) SetOptions(pdfg *wk.PDFGenerator) {
	if pc.Dpi != 0 {
		pdfg.Dpi.Set(pc.Dpi)
	}
	if pc.Orientation != "" {
		//pdfg.Orientation.Set(wk.OrientationPortrait) //OrientationLandscape
		pdfg.Orientation.Set(pc.Orientation)
	}
	if pc.PageSize != "" {
		//pdfg.PageSize.Set(wk.PageSizeA4)
		pdfg.PageSize.Set(pc.PageSize)
	}
	pdfg.Grayscale.Set(pc.Grayscale)
	//pdfg.NoCollate.Set(false)
}

// GetConf is to return ConfigPDF
func (cr *CreateFromURL) GetConf() *ConfigPDF {
	return cr.ConfigPDF
}

// NewGenerator is to generate initial PDFGenerator object
func (cr *CreateFromURL) NewGenerator() *wk.PDFGenerator {
	defer tm.Track(time.Now(), "CreateFromURL.NewGenerator()") //78.5µs

	pdfg, err := wk.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	return pdfg
}

// CreatePage is to create page from URL
func (cr *CreateFromURL) CreatePage(pdfg *wk.PDFGenerator) bool {
	fmt.Println("Create by URL")
	defer tm.Track(time.Now(), "CreateFromURL.CreatePage()") //29.527µs

	page := wk.NewPage(cr.URL)

	// Set options for this page
	//setFooters(page)

	// Add to document
	pdfg.AddPage(page)

	return true
}

// CreatePDFBuffer is to create pdf buffer from url generator
func (cr *CreateFromURL) CreatePDFBuffer(pdfg *wk.PDFGenerator) *wk.PDFGenerator {
	defer tm.Track(time.Now(), "CreateFromURL.CreatePDFBuffer()") //2.35s

	err := pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// debug size
	fmt.Printf(" PDF size %d bytes\n", pdfg.Buffer().Len())

	return pdfg
}

// CreateFromFile is File object corresponding with Creator interface
type CreateFromFile struct {
	FilePath string
	File     *os.File
	*ConfigPDF
}

// GetConf is to return ConfigPDF
func (cr *CreateFromFile) GetConf() *ConfigPDF {
	return cr.ConfigPDF
}

// NewGenerator is to generate initial PDFGenerator object
func (cr *CreateFromFile) NewGenerator() *wk.PDFGenerator {
	defer tm.Track(time.Now(), "CreateFromFile.NewGenerator()") //29.407µs

	pdfg := wk.NewPDFPreparer()
	return pdfg
}

// CreatePage is to create page from Filepath
func (cr *CreateFromFile) CreatePage(pdfg *wk.PDFGenerator) bool {
	fmt.Println("Create by file path")
	defer tm.Track(time.Now(), "CreateFromFile.CreatePage()") //89.361µs

	// add a reader page as well
	var err error
	cr.File, err = os.Open(cr.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wk.NewPageReader(cr.File))

	return true
}

// CreatePDFBuffer is to create pdf buffer from filepath generator
func (cr *CreateFromFile) CreatePDFBuffer(pdfg *wk.PDFGenerator) *wk.PDFGenerator {
	defer tm.Track(time.Now(), "CreateFromFile.CreatePDFBuffer()") //1.47s

	// The html string is also saved as base64 string in the JSON file
	defer func() {
		err := cr.File.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

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

	// debug size
	fmt.Printf(" PDF size %d bytes\n", pdfgFromJSON.Buffer().Len())

	return pdfgFromJSON
}
