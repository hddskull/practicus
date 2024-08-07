package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/hddskull/practicus/internal/storage"
)

const (
	counter = "counter"
	gauge   = "gauge"
)

func UpdateMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusUnauthorized)
		return
	}

	// resp := "aboba\n"

	mType, err := getMetricType(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// resp += fmt.Sprintln(mType)

	mName, err := getMetricName(r.URL.Path)
	// _, err = getMetricName(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// resp += fmt.Sprintln(mName)

	value, err := validateValue(mType, r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := storage.NewMemStorage()

	s.UpdateMetric(mName, mType, value)

	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte(resp))
}

func getMetricType(p string) (string, error) {
	components := strings.Split(p, "/")

	if len(components) < 3 {
		return "", errors.New("invalid path")
	}

	mType := components[2]

	if mType == "counter" || mType == "gauge" {
		return mType, nil
	} else {
		return "", errors.New("invalid path")
	}
}

func getMetricName(p string) (string, error) {
	p = p[:len(p)-1]
	components := strings.Split(p, "/")

	if len(components) < 4 {
		return "", errors.New("invalid path")
	}

	mName := components[3]

	return mName, nil
}

func validateValue(t, p string) (string, error) {
	slice := strings.Split(p, "/")
	value := slice[len(slice)-1]
	value = strings.ReplaceAll(value, " ", "")

	var err error

	switch t {
	case counter:
		_, err = strconv.Atoi(value)

	case gauge:
		_, err = strconv.ParseFloat(value, 64)

	default:
		return "", errors.New("invalid metric type")
	}

	if err != nil {
		return "", errors.New("invalid metric type")
	}

	return value, nil
}
