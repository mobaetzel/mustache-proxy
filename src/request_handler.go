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
		rawData := r.URL.Query().Get("data")

		parsedSrcUrl, err := url.Parse(srcUrl)
		if err != nil {
			panic(err)
		}

		shouldBlock := true
		for _, allowedTarget := range allowedTargets {
			if parsedSrcUrl.Host == allowedTarget {
				shouldBlock = false
				continue
			}
		}

		if shouldBlock {
			w.WriteHeader(401)
			w.Write([]byte("unauthotized"))
			return
		}

		var data map[string]interface{}
		err = json.Unmarshal([]byte(rawData), &data)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("bad request"))
			return
		}

		resp, err := http.Get(srcUrl)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("not found"))
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(424)
			w.Write([]byte("failed dependency"))
			return
		}

		result, err := mustache.Render(string(body), data)
		if err != nil {
			w.WriteHeader(406)
			w.Write([]byte("not acceptable"))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(result))
	}
}