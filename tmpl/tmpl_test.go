package tmpl_test

import (
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/tmpl"
	//ht "html/template"
	"os"
	"testing"
	tt "text/template"
)

type TeachrInfo struct {
	Id      int
	Name    string
	Country string
}

type SiteInfo struct {
	Url      string
	Teachers []TeachrInfo
}

type Site struct {
	Name  string
	Pages []int
}

var tmplTeachers string = `
	//1. just dot
	{{.}}

	//2. variable
	Url:{{.Url}}

	//3. range
	{{range .Teachers}}
	- Id:{{.Id}}
	- Name:{{.Name}}
	- Country:{{.Country}}
	{{end}}
	-----------------------------------
`

var tmplTeachers2 string = `
	//1. range
	{{range $index, $value := .Teachers}}
	- Index:{{$index}}
	- Value:{{$value}}
	 - Id:{{.Id}}
	- Name:{{.Name}}
	- Country:{{.Country}}
	{{end}}

	* A eq B -> A == B
	* A ne B -> A != B
	* A lt B -> A < B
	* A le B -> A <= B
	* A gt B -> A > B
	* A ge B -> A >= B
	-----------------------------------
`

var tmplTeachers3 string = `
	//1. range
	{{range .Teachers}}
	- Id:{{.Id | plus10}}
	- Name:{{.Name}}
	- Country:{{.Country}}
	{{end}}
	-----------------------------------
`

var tmplTeachers4 string = `
	//1. range
	{{range $index, $value := .Teachers}}
	{{if .Name}}
	- Id:{{.Id | plus10}}
	- Name:{{.Name}}
	- Country:{{.Country}}
	{{else}}
	Teacher's name is not found.
	{{end}}

	{{if eq $index 1}}
	only index = 1
	{{end}}

	{{if eq $index 3 4 }}
	only index = 3 or 4
	{{end}}
	{{end}}
	-----------------------------------
`

var tmplSite string = `
	{{range .Pages}}
		<li><a href="{{.}}">{{.}}</a></li>
	{{end}}
`

var innerOuter string = `
{{with .Inner}}
  Outer: {{$.OuterValue}}
  Inner: {{.InnerValue}}
{{end}}
`

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[TMPL]")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// functions
//-----------------------------------------------------------------------------
func plus10(num int) int {
	return num + 10
}

func getTestData() *SiteInfo {
	siteInfo := &SiteInfo{Url: "http://google.com",
		Teachers: []TeachrInfo{{Id: 1, Name: "Harry", Country: "Japan"},
			{Id: 2, Name: "Harry", Country: "Japan"},
			{Id: 3, Name: "Taro", Country: "UK"},
			{Id: 4, Name: "", Country: "Germany"},
			{Id: 5, Name: "Saburo", Country: "America"}}}

	return siteInfo
}

//-----------------------------------------------------------------------------
// Check
//-----------------------------------------------------------------------------
func TestTextTemplate(t *testing.T) {
	//t.Skip("skipping TestTextTemplate")

	siteInfo := getTestData()

	//
	tmpl, err := tt.New("siteInfoText").Parse(tmplTeachers)
	if err != nil {
		t.Fatalf("[01] template.New() error: %s", err)
	}
	err = tmpl.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("[02] template.Execute() error: %s", err)
	}

	//2
	tmpl2, err := tt.New("siteInfoText2").Parse(tmplTeachers2)
	if err != nil {
		t.Fatalf("[03] template.New() error: %s", err)
	}
	err = tmpl2.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("[04] template.Execute() error: %s", err)
	}

	//3
	tmpl3, err := tt.New("siteInfoText3").Funcs(tt.FuncMap{"plus10": plus10}).Parse(tmplTeachers3)
	if err != nil {
		t.Fatalf("[05] template.New() error: %s", err)
	}
	err = tmpl3.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("[06] template.Execute() error: %s", err)
	}

	//4
	tmpl4, err := tt.New("siteInfoText4").Funcs(tt.FuncMap{"plus10": plus10}).Parse(tmplTeachers4)
	if err != nil {
		t.Fatalf("[07] template.New() error: %s", err)
	}
	err = tmpl4.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("[08] template.Execute() error: %s", err)
	}
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestFilePathParser(t *testing.T) {
	siteInfo := getTestData()

	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/tmpl/file/sample1.tmpl"
	result, err := FilePathParser(path, siteInfo)
	if err != nil {
		t.Fatalf("[01]FilePathParser() error: %s", err)
	}
	lg.Debug(result)
}

func TestFileTemplate(t *testing.T) {
	siteInfo := getTestData()

	goPath := os.Getenv("GOPATH")
	tpl, err := tt.ParseFiles(goPath + "/src/github.com/hiromaily/golibs/tmpl/file/sample1.tmpl")
	if err != nil {
		t.Fatalf("[01] template.ParseFiles() error: %s", err)
	}

	result, err := FileTempParser(tpl, siteInfo)
	if err != nil {
		t.Fatalf("[02] FileTempParser() error: %s", err)
	}
	lg.Debug(result)
}
