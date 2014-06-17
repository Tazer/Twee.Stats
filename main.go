// Twee.Stats.Api project main.go
package main

import (
	"encoding/json"
	"github.com/dpapathanasiou/go-api"
	"net/http"
)

type Message struct {
	Text string
}

func helloWorldJSON(w http.ResponseWriter, r *http.Request) string {
	// while we're not using r in this example, the http.Request object
	// has several attributes which help inform what the exact reply will be
	// (see http://golang.org/pkg/net/http/#Request for the full list of attributes,
	// as well as the weather-api.go example file in this package to get an idea of what's possible)
	m := Message{"Hello World"}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err) // no, not really
	}

	return string(b)
}

func main() {
	handlers := map[string]func(http.ResponseWriter, *http.Request){}
	handlers["/hello/"] = func(w http.ResponseWriter, r *http.Request) {
		api.Respond("application/json", "utf-8", helloWorldJSON)(w, r)
	}

	api.NewLocalServer(9001, api.DefaultServerReadTimeout, handlers)
}
