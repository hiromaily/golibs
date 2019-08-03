package i18n

import (
	"fmt"
	ht "html/template"
	"os"

	"github.com/nicksnyder/go-i18n/i18n"
	"gopkg.in/guregu/null.v3"
)

//https://github.com/nicksnyder/go-i18n

// Translation is translated data
type Translation struct {
	msgID string
	args  []interface{}
}

// Languages is en, nl, de, fr and so on.
var (
	Languages []string
	files     = []string{"en-us.all", "nl-nl.all", "de-de.all", "fr-fr.all"}
)

func init() {
	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/i18n/locale/"
	//wd, _ := os.Getwd()
	//if strings.HasSuffix(wd, "test") {
	//	wd = ".."
	//}
	for _, v := range files {
		i18n.MustLoadTranslationFile(fmt.Sprintf("%s%s.json", path, v))
	}
	Languages = i18n.LanguageTags()
}

// MustTfunc is
func MustTfunc(msgID string, args ...interface{}) func(lang string) string {
	return func(lang string) string {
		return i18n.MustTfunc(lang)(msgID, args...)
	}
}

// GetTranslations is to get translated message by ID
func GetTranslations(msgID string, args ...interface{}) map[string]string {
	res := make(map[string]string, len(Languages))
	for _, tag := range Languages {
		res[tag] = MustTfunc(msgID, args...)(tag)
	}
	return res
}

// T is to return Translation struct
func T(msgID string, args ...interface{}) Translation {
	return Translation{msgID, args}
}

// Map is to get map data
func (t Translation) Map() map[string]string {
	return GetTranslations(t.msgID)
}

// String is to get string data
func (t Translation) String(lang string) string {
	return MustTfunc(t.msgID, t.args...)(lang)
}

// GetFuncMap is for calling from template
func GetFuncMap(lang string) ht.FuncMap {
	funcMap := ht.FuncMap{
		"gettext": func(id string) string {
			return T(id).String(lang)
		},
		"gettext2": func(id string, args ...interface{}) ht.HTML {
			text := T(id, args...).String(lang)
			return ht.HTML(text)
		},
		"getAsMap": func(args ...interface{}) map[string]interface{} {
			//key: value
			maps := map[string]interface{}{}
			var val string
			for i := 0; i < len(args); i += 2 {
				key, _ := args[i].(string)
				//check for null.String type as possibility
				if v, ok := args[i+1].(null.String); ok {
					val = v.String
				} else {
					val, _ = args[i+1].(string)
				}
				maps[key] = val
			}
			return maps
		},
		"unescape": func(text string) ht.HTML { return ht.HTML(text) },
	}
	return funcMap
}
