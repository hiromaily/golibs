package pdf

import (
	"os"
	"testing"
)

func BenchmarkCreatePDF_URL(b *testing.B) {
	//2.547,477,387 ns/op

	url := "https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf"
	output := "from_url.pdf"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createPDFFromURL(url, output, createConfig())
	}
	b.StopTimer()
}

func BenchmarkCreatePDF_File(b *testing.B) {
	//1.458,145,617 ns/op

	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/pdf/cmd/"
	filePath := path + "testfiles/tables/index.html"
	output := "from_file.pdf"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createPDFFromURL(filePath, output, createConfig())
	}
	b.StopTimer()
}
