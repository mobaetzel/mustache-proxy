package src

import (
	"encoding/json"
	"github.com/cbroglie/mustache"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createRequestHandler(allowedTargets []string)  func(w http.ResponseWriter, r* http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		srcUrl := r.URL.Query().Get("src")

		parsedSrcUrl, err := url.Parse(srcUrl)
		if err != nil {
			panic(err)
		}

		shouldBlock := true
		for _, allowedTarget := range allowedTargets {
			if parsedSrcUrl.Host == allowedTarget {
				shouldBlock = false
				break
			}
		}

		if shouldBlock {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		rawData := r.URL.Query().Get("data")

		var data map[string]interface{}
		err = json.Unmarshal([]byte(rawData), &data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}

		resp, err := http.Get(srcUrl)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("broken target"))
			return
		}

		result, err := mustache.Render(string(body), data)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("not acceptable"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}