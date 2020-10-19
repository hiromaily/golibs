package tmpl_test

import (
	"fmt"
	"os"
	"testing"
	tt "text/template"

	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/tmpl"
)

//TODO
//http://stackoverflow.com/questions/11467731/is-it-possible-to-have-nested-templates-in-go-using-the-standard-library-googl

type TeachrInfo struct {
	ID      int
	Name    string
	Country string
}

type Teachers struct {
	URL      string
	Teachers []TeachrInfo
}

//type Site struct {
//	Name  string
//	Pages []int
//}

var (
	//test1
	tmplTeachers = `
	//1. use  dot
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
	//test2
	tmplTeachers2 = `
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
	//test3
	tmplVariable = `
	{{$test := "test"}}
	{{if eq $test "test"}}
		{{$test}}
		{{printf "This is %s" $test}}
	{{end}}
`

	//	//test4
	//	tmplSite = `
	//	{{range .Pages}}
	//		<li><a href="{{.}}">{{.}}</a></li>
	//	{{end}}
	//`

	//	//test5
	//	innerOuter = `
	//	{{with .Inner}}
	//	  Outer: {{$.OuterValue}}
	//	  Inner: {{.InnerValue}}
	//	{{end}}
	//`

	//test6
	tmplTeachers3 = `
	//1. range
	{{range .Teachers}}
	- Id:{{.Id | plus10}}
	- Name:{{.Name}}
	- Country:{{.Country}}
	{{end}}
	-----------------------------------
`
	//test7
	tmplTeachers4 = `
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
)

// For table test
var tmpleTests = []struct {
	tmplName   string
	tmplString string
}{
	{"tmplTeachers", tmplTeachers},
	{"tmplTeachers2", tmplTeachers2},
	{"tmplVariable", tmplVariable},
	//{"tmplSite", tmplSite},
	//{"innerOuter", innerOuter},
}

var tmpleTests2 = []struct {
	tmplName   string
	tmplString string
}{
	{"tmplTeachers3", tmplTeachers3},
	{"tmplTeachers4", tmplTeachers4},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[TMPL]")
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

func getTeathers() *Teachers {
	teachers := &Teachers{URL: "http://google.com",
		Teachers: []TeachrInfo{{ID: 1, Name: "Harry", Country: "Japan"},
			{ID: 2, Name: "Harry", Country: "Japan"},
			{ID: 3, Name: "Taro", Country: "UK"},
			{ID: 4, Name: "", Country: "Germany"},
			{ID: 5, Name: "Saburo", Country: "America"}}}

	return teachers
}

//-----------------------------------------------------------------------------
// Check
//-----------------------------------------------------------------------------
func TestTextTemplate(t *testing.T) {
	//t.Skip("skipping TestTextTemplate")

	teachers := getTeathers()

	//normal
	for i, td := range tmpleTests {
		fmt.Printf("[%d] name:%s\n", i+1, td.tmplName)

		tmpl, err := tt.New(td.tmplName).Parse(td.tmplString)
		if err != nil {
			t.Fatalf("[%d] template.New(%s) error: %s", i+1, td.tmplName, err)
		}
		err = tmpl.Execute(os.Stdout, teachers)
		if err != nil {
			t.Fatalf("[%d] template.Execute(%s) error: %s", i+1, td.tmplName, err)
		}
	}

	//with func
	for i, td := range tmpleTests2 {
		tmpl, err := tt.New("td.tmplName").Funcs(tt.FuncMap{"plus10": plus10}).Parse(td.tmplString)
		if err != nil {
			t.Fatalf("[%d] template.New(%s) error: %s", i+1, td.tmplName, err)
		}
		err = tmpl.Execute(os.Stdout, teachers)
		if err != nil {
			t.Fatalf("[%d] template.Execute(%s) error: %s", i+1, td.tmplName, err)
		}
	}
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestFilePathParser(t *testing.T) {
	siteInfo := getTeathers()

	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/tmpl/file/sample1.tmpl"
	result, err := FilePathParser(path, siteInfo)
	if err != nil {
		t.Fatalf("[01]FilePathParser() error: %s", err)
	}
	lg.Debug(result)
}

func TestFileTemplate(t *testing.T) {
	siteInfo := getTeathers()

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

func TestTmplPlusI18n(t *testing.T) {

}
