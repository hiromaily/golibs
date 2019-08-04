package handlers

import (
	"fmt"
	"net/http"
)

type HealthChecker struct {
	isRunning bool
}

func NewHealthChecker(isRunning bool) HealthChecker {
	return HealthChecker{
		isRunning: isRunning,
	}
}

func (hc HealthChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if hc.isRunning {
		fmt.Fprint(w, "ok")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "ng")
	}
}

func (hc HealthChecker) Switch(isRunning bool) {
	hc.isRunning = isRunning
}
