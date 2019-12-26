package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	root := flag.String("root", "", "Root of the static files to serve")
	flag.Parse()

	if root == nil || *root == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Starting file server in", *root)

	dir := http.Dir(*root)
	fmt.Println(http.ListenAndServe(":80", http.FileServer(dir)))
}
