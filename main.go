package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/rickroll": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
	}
	mapVersion := mapHandler(pathsToUrls, mux)

	// yaml := `
	// 	- path: "/rickroll"
	// 	  url: "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	// `
	// yamlVersion, err := yamlHandler([]byte(yaml), mux)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapVersion)

	// fmt.Println("Starting the server on :8090")
	// http.ListenAndServe(":8090", yamlVersion)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
