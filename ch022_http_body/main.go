package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, World!")

	contentType := r.Header.Get("Content-Type")
	fmt.Println("Content-Type:", contentType)

	switch {
	case contentType == "application/json":
		handleJSON(w, r)
	case contentType == "application/x-www-form-urlencoded":
		handleFormURLEncoded(w, r)
	case strings.Contains(contentType, "multipart/form-data"):
		handleFileUpload(w, r)
	default:
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
	}
}

// 处理JSON数据
func handleJSON(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received JSON: %v", data)
}

// 处理表单数据  `application/x-www-form-urlencoded`
func handleFormURLEncoded(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Received form data: %+v", r.Form)
}

// 处理文件上传 `multipart/form-data`
func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 * 2^20 bytes，限制文件大小为10MB
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "File uploaded successfully")
}
