package tmpl

import (
	"bytes"
	"encoding/json"
	//ht "html/template"
	tt "text/template"
)

// StrTempParser is to get string of template for string
// It's just example
func StrTempParser(temp string, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	funcMap := tt.FuncMap{
		"add": func(a, b int) int { return a + b },
		"dump": func(in interface{}) string {
			jsonData, err := json.MarshalIndent(in, "", "  ")
			if err != nil {
				return err.Error()
			}
			return string(jsonData)
		},
	}

	tpl := tt.Must(tt.New("tpl").Funcs(funcMap).Parse(temp))

	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}

	return parseResult.String(), nil
}

// FilePathParser is to get string from template file
func FilePathParser(path string, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	tpl := tt.Must(tt.ParseFiles(path))
	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}
	return parseResult.String(), nil
}

// FileTempParser is to get string from template.Template
func FileTempParser(tpl *tt.Template, params interface{}) (string, error) {
	var parseResult bytes.Buffer

	if err := tpl.Execute(&parseResult, params); err != nil {
		return "", err
	}
	return parseResult.String(), nil
}
