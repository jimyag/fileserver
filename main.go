/*
Serve is a very simple static file server in go
Usage:

	-p="8100": port to serve on
	-d=".":    the directory of static files to host

Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	version = ""
	commit  = ""
	date    = ""
	builtBy = ""
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 && (args[0] == "version" || args[0] == "-v") {
		str := ""
		if version != "" {
			str += "Version: " + version
		}
		if commit != "" {
			str += "\nCommit: " + commit
		}
		if date != "" {
			str += "\nDate: " + date
		}
		if builtBy != "" {
			str += "\nBuilt by: " + builtBy
		}
		if str != "" {
			log.Println(str)
		}
		os.Exit(0)
	}
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", "./", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		http.ServeFile(w, r, *directory+r.URL.Path)
	}))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
