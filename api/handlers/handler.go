package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-rod/rod/lib/utils"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"scraper/helpers"
	"scraper/internal"
	"scraper/internal/models"
	"strconv"
	"time"
)

const PATH_PDF = "static/pdf/"
const PATH_IMG = "static/img/"

var counter = 0

func ImageHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	var response map[string]interface{}
	timestamp := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%v_%s.jpeg", counter, timestamp)
	counter += 1
	switch req.Method {
	case "GET":
		var startTime = time.Now()
		var url = req.URL.Query().Get("url")
		browser := internal.InitBrowser().MustPage(url)
		utils.Sleep(4)
		browser.MustScreenshot(PATH_IMG + filename)
		time.Sleep(3)
		defer func() {
			log.Printf("Closing browser...")
			browser.MustClose()
		}()
		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, helpers.GetImageFromPath(PATH_IMG+filename), nil); err != nil {
			log.Println("unable to encode image.")
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
		if _, err := w.Write(buffer.Bytes()); err != nil {
			log.Println("unable to write image.")
		}
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("Image Screenshot Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		return
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func PdfHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/pdf")
	var response map[string]interface{}
	timestamp := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%v_%s.pdf", counter, timestamp)
	counter += 1
	switch req.Method {
	case "GET":
		var startTime = time.Now()
		var url = req.URL.Query().Get("url")
		browser := internal.InitBrowser().MustPage(url)
		utils.Sleep(4)
		browser.MustPDF(PATH_PDF + filename)
		time.Sleep(3)
		defer func() {
			log.Printf("Closing browser...")
			browser.MustClose()
		}()
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.NotFound(w, req)
			return
		}
		w.Header().Set("Transfer-Encoding", "chunked")
		buf, err := ioutil.ReadFile(PATH_PDF + filename)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(buf)
		flusher.Flush()
		//f, err := os.Open(PATH_PDF + filename)
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(500)
		//	return
		//}
		//defer f.Close()
		//f.SetDeadline(time.Now().Add(time.Second * 30))
		//Set header
		//Stream to response
		//if _, err := io.Copy(w, f); err != nil {
		//	fmt.Println(err)
		//	w.WriteHeader(500)
		//}
		var endTime = time.Now()
		duration := endTime.Sub(startTime)
		log.Printf("PDF Took: %v NanoSeconds or %v MiliSeconds ", duration.Nanoseconds(), duration.Milliseconds())
		return
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		return
	case req.Method:
		response = models.CreateResponse("failure", "method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		return
	}
}
