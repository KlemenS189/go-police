package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/KlemenS189/go-police/types"
	"github.com/julienschmidt/httprouter"
)

type OutputMode string

var (
	JSON OutputMode = "json"
	RAW  OutputMode = "raw"
)

var (
	port               string
	endpoint           string
	passThroughPercent float64
	writeToFile        bool
	fileName           string
	outputMode         OutputMode

	file *os.File
)

func main() {
	var err error
	flag.StringVar(&port, "port", "8000", "Port for violation server")
	flag.StringVar(&endpoint, "endpoint", "/csp/violation-report/", "Endpoint for reporting csp violations")
	flag.Float64Var(&passThroughPercent, "passThroughPercent", 1.0, "Percent of requests to be logged")

	flag.BoolVar(&writeToFile, "to-file", false, "Switch if output should be written to a file")
	flag.StringVar(&fileName, "filename", "csp_violation_report.txt", "File for saving csp violations. "+
		"Ignored if to-file is set to false. Datetime is appended to beginning")

	flag.Func("output-mode", "Write as JSON or plain text", func(s string) error {
		if OutputMode(s) != JSON && OutputMode(s) != RAW {
			return errors.New("Must be 'json' or 'raw'")
		}
		outputMode = OutputMode(s)
		return nil
	})
	flag.Parse()

	if writeToFile {
		now := time.Now().UTC().Format("20060102150405")
		fileName = fmt.Sprintf("%s-%s", now, fileName)
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Println("Error with opening a file")
			log.Fatal(err)
		}
	}

	router := initRouter()
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func violationHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !canPassThrough() {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var report types.CspReport
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&report)
	if err != nil {
		log.Println(err)
		http.Error(w, "Wrong csp report data", http.StatusBadRequest)
		return
	}
	var output string
	if outputMode == JSON {
		output = report.Json()
	} else {
		output = report.Raw()
	}

	if writeToFile {
		_, err := file.WriteString(fmt.Sprintf("%s\n", output))
		if err != nil {
			log.Println(err)
		}
	} else {
		// Log output to STDOUT
		log.Println(output)
	}
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST(endpoint, violationHandler)
	return router
}

func canPassThrough() bool {
	if passThroughPercent == 1.0 {
		return true
	}
	return rand.Float64() < passThroughPercent
}
