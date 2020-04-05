package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Address   *Address
}

type Address struct {
	City string
	Town string
	Sub  *SubAddr
}

type SubAddr struct {
	City string
	Town string
}

func TestZapLog(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//how about slice of pointer
	var (
		url   = "http://xxxxx.com"
		sub   = SubAddr{"city", "town"}
		addr  = Address{"city", "town", &sub}
		user  = User{10, "foo", "bar", &addr}
		addrs = []*Address{&addr}
	)

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
		zap.Any("user", user),
		zap.Any("user as address", &user),
		zap.Any("slice of address of address", addrs),
	)
}
