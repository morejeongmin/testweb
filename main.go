package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/", new(myHandler))

	http.ListenAndServe(":8888", nil)
}

type myHandler struct {
	http.Handler
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defPath := req.URL.Path
	if defPath == "/" {
		defPath += "index.html"
	}
	localPath := "wwwroot" + defPath
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
