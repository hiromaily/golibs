package errors_test

import (
	"testing"

	"fmt"
	lg "github.com/hiromaily/golibs/log"
	r "github.com/hiromaily/golibs/runtimes"
	"github.com/pkg/errors"
)

// "github.com/pkg/errors"
// これはwrapで複数のエラーを保持できるが、エラーの種類に応じて最終的にハンドリングしたい場合が
// ないのであれば使う必要はない

func first() error {
	err := second()
	if err != nil {
		//エラーの発生源のみ、loggerを仕込み、エラーメッセージも原因特定のために詳細を添える
		//それ以外は通常返すのみでいい
		return errors.Wrap(err, "failed to call second()")
	}

	return nil
}

func second() error {
	err := third()
	if err != nil {
		//エラーの発生源のみ、loggerを仕込み、エラーメッセージも原因特定のために詳細を添える
		//それ以外は通常返すのみでいい
		return errors.Wrap(err, "failed to call third()")
	}

	return nil
}

func third() error {
	//エラーメッセージ
	err := errors.Errorf("something error in %s()", r.CurrentFunc(1))

	//エラー発生源でloggerを仕込む
	//日時、goのファイル名、行数、エラーメッセージ
	lg.Error(err.Error())

	//stack trace
	//これは見づらい
	//log.Printf("%+v", err)
	lg.Stack()

	return err
}

func TestError(t *testing.T) {
	lg.InitializeLog(lg.DebugStatus, lg.DateTimeShortFile, "", "", "hiromaily")

	err := first()
	if err != nil {
		//最後のエラー処理はどうすればいい？
		fmt.Println(errors.Cause(err))
	}
}
