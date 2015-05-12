package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type reqConfig struct {
	code  int
	delay time.Duration
}

func newReqConfig() *reqConfig {
	return &reqConfig{code: 200, delay: 0 * time.Second}
}

var reqMapLock sync.RWMutex
var reqMap map[string]*reqConfig = map[string]*reqConfig{}

func getReqName(r *http.Request) string {
	return mux.Vars(r)[reqName]
}

func adminDelay(w http.ResponseWriter, r *http.Request) {
	rName := getReqName(r)
	if rName == "" {
		http.Error(w, "no request name", http.StatusBadRequest)
		return
	}

	delay, err := strconv.Atoi(mux.Vars(r)[delay])
	if err != nil {
		http.Error(w, "invalid delay", http.StatusBadRequest)
		return
	}

	reqMapLock.Lock()
	defer reqMapLock.Unlock()
	conf, ok := reqMap[rName]
	if !ok {
		conf = newReqConfig()
	}
	conf.delay = time.Duration(delay) * time.Second
	reqMap[rName] = conf

	w.WriteHeader(http.StatusOK)
}

func adminCode(w http.ResponseWriter, r *http.Request) {
	rName := getReqName(r)
	if rName == "" {
		http.Error(w, "no request name", http.StatusBadRequest)
		return
	}

	code, err := strconv.Atoi(mux.Vars(r)[code])
	if err != nil {
		http.Error(w, "invalid code", http.StatusBadRequest)
		return
	}

	reqMapLock.Lock()
	defer reqMapLock.Unlock()
	conf, ok := reqMap[rName]
	if !ok {
		conf = newReqConfig()
	}
	conf.code = code
	reqMap[rName] = conf

	w.WriteHeader(http.StatusOK)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rName := getReqName(r)
	if rName == "" {
		http.Error(w, "no request name", http.StatusBadRequest)
		return
	}

	reqMapLock.RLock()
	defer reqMapLock.RUnlock()
	conf, ok := reqMap[rName]
	if !ok {
		http.Error(w, "request name "+rName+" not found", http.StatusNotFound)
		return
	}
	time.Sleep(conf.delay)
	w.WriteHeader(conf.code)
}
