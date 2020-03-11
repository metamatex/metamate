package config

import (
	"encoding/json"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"gopkg.in/yaml.v2"
	"net/http"
)

func GetJsonConfigHandleFunc(c types.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Fields", "application/json")

		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetYamlConfigHandleFunc(c types.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Fields", "text/x-yaml")

		err := yaml.NewEncoder(w).Encode(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
