package main

import (
	"os"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io"
	"log"
	"net/http"
	"fmt"
)


// this function seems has problem
func UploadFile(uri string, params map[string]string, paramName, path string, verbose bool) (string, bool){
	file, err := os.Open(path)
	if err != nil {
		return "file failed to open " + err.Error(), false
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return "error in write file " + err.Error(), false
	}
	_, err = io.Copy(part, file)
	if err != nil{
		return "error in copy file to part " + err.Error(), false
	}

	if params != nil {
		for key, val := range params {
			_ = writer.WriteField(key, val)
		}
	}

	err = writer.Close()
	if err != nil {
		return "error in write file " + err.Error(), false
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		log.Fatal(err)
	}
	// don't forget content-type, if not will lose boundary
	request.Header.Set("Content-Type", writer.FormDataContentType())


	client := &http.Client{}
	// do post and the response
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return "request error " + err.Error(), false
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()

		if verbose {
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Header)
			fmt.Println(body)
		}
		return body.String(), true
	}
}
