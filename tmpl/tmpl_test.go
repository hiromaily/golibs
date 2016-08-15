package tmpl_test

import (
	//. "github.com/hiromaily/golibs/tmpl"
	lg "github.com/hiromaily/golibs/log"
	o "github.com/hiromaily/golibs/os"
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

var (
	benchFlg bool = false
)

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

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	lg.InitializeLog(lg.DEBUG_STATUS, lg.LOG_OFF_COUNT, 0, "[TMPL_TEST]", "/var/log/go/test.log")
	if o.FindParam("-test.bench") {
		lg.Debug("This is bench test.")
		benchFlg = true
	}
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

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestTextTemplate(t *testing.T) {
	funcName := "TestTextTemplate"
	//t.Skip("skipping TestTextTemplate")
	siteInfo := SiteInfo{Url: "http://google.com",
		Teachers: []TeachrInfo{{Id: 1, Name: "Harry", Country: "Japan"},
			{Id: 2, Name: "Harry", Country: "Japan"},
			{Id: 3, Name: "Taro", Country: "UK"},
			{Id: 4, Name: "", Country: "Germany"},
			{Id: 5, Name: "Saburo", Country: "America"}}}

	//
	tmpl, err := tt.New("siteInfoText").Parse(tmplTeachers)
	if err != nil {
		t.Fatalf("%s[01] error: %s", funcName, err)
	}
	err = tmpl.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("%s[02] error: %s", funcName, err)
	}

	//2
	tmpl2, err := tt.New("siteInfoText2").Parse(tmplTeachers2)
	if err != nil {
		t.Fatalf("%s[03] error: %s", funcName, err)
	}
	err = tmpl2.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("%s[04] error: %s", funcName, err)
	}

	//3
	tmpl3, err := tt.New("siteInfoText3").Funcs(tt.FuncMap{"plus10": plus10}).Parse(tmplTeachers3)
	if err != nil {
		t.Fatalf("%s[05] error: %s", funcName, err)
	}
	err = tmpl3.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("%s[05] error: %s", funcName, err)
	}

	//4
	tmpl4, err := tt.New("siteInfoText4").Funcs(tt.FuncMap{"plus10": plus10}).Parse(tmplTeachers4)
	if err != nil {
		t.Fatalf("%s[06] error: %s", funcName, err)
	}
	err = tmpl4.Execute(os.Stdout, siteInfo)
	if err != nil {
		t.Fatalf("%s[07] error: %s", funcName, err)
	}
}