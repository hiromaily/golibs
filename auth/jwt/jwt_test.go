package jwt_test

import (
	"os"
	"testing"
	"time"

	. "github.com/hiromaily/golibs/auth/jwt"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
func setup() {
	tu.InitializeTest("[JWT]")
	//
	path := os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/testdata/key/"

	priKey := path + "jwt_rsa"
	pubKey := path + "jwt_rsa.pub"
	err := InitKeys(priKey, pubKey)
	if err != nil {
		panic(err)
	}
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
// function
//-----------------------------------------------------------------------------
func createToken(t *testing.T, mode uint8) {
	InitEncrypted(mode)

	ti := time.Now().Add(time.Second * 2).Unix()
	token, err := CreateBasicToken(ti, "client123", "harry")
	if err != nil {
		t.Errorf("[%d] CreateBasicToken() error: %s", mode, err)
	}
	lg.Debugf("[%d]token: %s", mode, token)

	err = JudgeJWTWithClaim(token, "client123", "harry")
	if err != nil {
		//verification error
		t.Errorf("[%d] 1.JudgeJWT() error: %s", mode, err)
	}

	//sleep
	time.Sleep(time.Second * 3)

	err = JudgeJWTWithClaim(token, "client123", "harry")
	if err == nil {
		t.Errorf("[%d] 2.JudgeJWT() error has to be set: %s", mode, "Token is expired")
	}
}

func createUserToken(t *testing.T, mode uint8) {
	InitEncrypted(mode)

	ti := time.Now().Add(time.Second * 2).Unix()
	token, err := CreateToken(ti, "client123", "harry", "option555")
	if err != nil {
		t.Errorf("[%d] CreateToken() error: %s", mode, err)
	}
	lg.Debugf("[%d]token: %s", mode, token)

	err = JudgeJWTWithCustomClaim(token, "client123", "harry", "option555")
	if err != nil {
		t.Errorf("[%d] 1.JudgeJWTWithClaim() error: %s", mode, err)
	}

	// set different name
	err = JudgeJWTWithCustomClaim(token, "client123", "harry", "option777")
	if err == nil {
		t.Errorf("[%d] 2.JudgeJWTWithClaim() error has to be set: %s", mode, "Option is invalid")
	}

	time.Sleep(time.Second * 3)

	err = JudgeJWTWithCustomClaim(token, "client123", "harry", "option555")
	if err == nil {
		t.Errorf("[%d] 3.JudgeJWTWithClaim() error has to be set: %s", mode, "Token is expired")
	}
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestCreateToken(t *testing.T) {
	//t.SkipNow()
	createToken(t, HMAC)
	createToken(t, RSA)
}

func TestCreateUserToken(t *testing.T) {
	//t.SkipNow()
	createUserToken(t, HMAC)
	createUserToken(t, RSA)
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkJwt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
