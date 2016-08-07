package tmpl

import (
	"bytes"
	"text/template"
)

//template for string
func StrTempParser(temp string, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	tpl := template.Must(template.New("tpl").Funcs(funcMap).Parse(temp))

	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}

	return parseResult.String(), nil
}

//template for file
func FilePathParser(path string, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	tpl := template.Must(template.ParseFiles(path))
	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}
	return parseResult.String(), nil
}

func FileTempParser(tpl *template.Template, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}
	return parseResult.String(), nil
}
