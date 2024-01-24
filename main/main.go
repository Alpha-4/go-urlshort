package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"urlshort"

	"gopkg.in/yaml.v3"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	/*
			yaml := `
		- path: /urlshort
		  url: https://github.com/gophercises/urlshort
		- path: /urlshort-final
		  url: https://github.com/gophercises/urlshort/tree/solution
		`
	*/
	// Provide the path to the file you want to read
	filePath := "mappings.yaml"

	// Read the entire file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var file yml

	if err := yaml.Unmarshal(content, &file); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%+v\n", string(content))
	//fmt.Printf("%+v\n", file.Paths)

	yamlHandler := urlshort.YAMLHandler(file.Paths, mapHandler)

	if err != nil {
		fmt.Println("Error fetching your response", err)
		return
	}

	fmt.Println("Starting the server on http://localhost:8080/")
	//http.ListenAndServe(":8080", mapHandler)
	http.ListenAndServe(":8080", yamlHandler)
}

type yml struct {
	Paths map[string]string `yaml:",inline"`
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", hello)
	mux.HandleFunc("/", err)
	return mux
}

func err(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "no maping found")
}
