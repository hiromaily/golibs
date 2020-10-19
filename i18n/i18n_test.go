package i18n_test

//import (
//	"os"
//	"strings"
//	"testing"
//
//	. "github.com/hiromaily/golibs/i18n"
//	lg "github.com/hiromaily/golibs/log"
//	tu "github.com/hiromaily/golibs/testutil"
//)
//
////-----------------------------------------------------------------------------
//// Test Framework
////-----------------------------------------------------------------------------
//
//func setup() {
//	tu.InitializeTest("[I18n]")
//}
//
//func teardown() {
//}
//
//func TestMain(m *testing.M) {
//	setup()
//
//	code := m.Run()
//
//	teardown()
//
//	os.Exit(code)
//}
//
////-----------------------------------------------------------------------------
//// function
////-----------------------------------------------------------------------------
//
////-----------------------------------------------------------------------------
//// Test
////-----------------------------------------------------------------------------
//func TestGetTranslations(t *testing.T) {
//	trans := GetTranslations("message1")
//	for _, lang := range Languages {
//		if trans[lang] == "" {
//			t.Errorf("GetTranslations couldn't get value :lang[%s]", lang)
//		}
//	}
//}
//
//func TestGetTranslationsWithNumberArgument(t *testing.T) {
//	t.Skip("TestGetTranslationsWithNumberArgument")
//	for _, lang := range Languages {
//		text := T("message3", 35).String(lang)
//		lg.Debug(lang, text)
//		if strings.Contains(text, "<no value>") {
//			t.Errorf("arguments was not replaced: %s", text)
//		}
//	}
//}
//
//func TestGetTranslationsWithStringArgument(t *testing.T) {
//	t.Skip("TestGetTranslationsWithStringArgument")
//	for _, lang := range Languages {
//		text := T("message2", "Mark").String(lang)
//		lg.Debug(lang, text)
//		if strings.Contains(text, "<no value>") {
//			t.Errorf("arguments was not replaced: %s", text)
//		}
//	}
//}
//
//func TestGetTranslationsWithStringArgumentWithMap(t *testing.T) {
//	//t.Skip("TestGetTranslationsWithStringArgumentWithMap")
//	mapData := map[string]interface{}{"Name": "Mark"}
//	for _, lang := range Languages {
//		text := T("message2", mapData).String(lang)
//		lg.Debug(lang, text)
//		if strings.Contains(text, "<no value>") {
//			t.Errorf("arguments was not replaced: %s", text)
//		}
//	}
//}
//
////-----------------------------------------------------------------------------
//// Benchmark
////-----------------------------------------------------------------------------
//func BenchmarkI18n(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		//
//		//_ = CallSomething()
//		//
//	}
//	b.StopTimer()
//}
