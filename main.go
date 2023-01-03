package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Handler struct {
	dirRoot      string
	defaultIndex string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mimeTypes := map[string]string{
		".html": "text/html",
		".js":   "text/javascript",
		".ico":  "image/vnd.microsoft.icon",
		".css":  "text/css",
	}

	urlPath := req.URL.Path
	if urlPath == "/" {
		urlPath = h.defaultIndex
	}
	finalPath := filepath.Join(h.dirRoot, urlPath)
	fmt.Println(req.Method, finalPath)
	f, err := os.Open(finalPath)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("ERROR:", err.Error())
		return
	}
	defer f.Close()

	ext := filepath.Ext(finalPath)

	w.Header().Set("content-type", mimeTypes[ext])
	w.Header().Set("cache-control", "no-cache")

	_, err = io.Copy(w, f)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("ERROR:", err.Error())
		return
	}
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "directory to serve HTTP")

	var port int
	flag.IntVar(&port, "port", 8080, "port to serve on")

	var defaultIndex string
	flag.StringVar(&defaultIndex, "default-index", "index.html", "default filename for an index HTML file")

	var certFile string
	flag.StringVar(&certFile, "certfile", "", "Filepath to server file for TLS.")

	var keyFile string
	flag.StringVar(&keyFile, "keyfile", "", "Filepath to key file for TLS.")

	flag.Parse()
	fmt.Println("--dir:", dir)
	fmt.Println("--port:", port)
	fmt.Println("--default-index:", defaultIndex)
	fmt.Println("--certfile:", certFile)
	fmt.Println("--keyfile:", keyFile)

	var err error
	if keyFile != "" && certFile != "" {
		err = http.ListenAndServeTLS(
			fmt.Sprintf(":%d", port),
			certFile,
			keyFile,
			Handler{
				dir,
				defaultIndex,
			})
	} else if keyFile != "" || certFile != "" {
		err = fmt.Errorf("Both --certfile and --keyfile must be set together, or unset together. (--certfile=%q, --keyfile=%q)", certFile, keyFile)
	} else {
		err = http.ListenAndServe(
			fmt.Sprintf(":%d", port),
			Handler{
				dir,
				defaultIndex,
			})
	}
	log.Fatal(err)
}
