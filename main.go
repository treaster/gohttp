package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "directory to serve HTTP")

	var port int
	flag.IntVar(&port, "port", 8080, "port to serve on")
	flag.Parse()

	fmt.Println("dir:", dir)
	fmt.Println("port:", port)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			http.FileServer(http.Dir(dir))))
}
