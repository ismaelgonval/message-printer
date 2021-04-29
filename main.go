package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
)

func Printer(w http.ResponseWriter, req *http.Request) {
	color.HiBlack(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	defer color.HiBlack("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

	color.HiBlue("====== HEADERS ======")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", color.HiCyanString(name), h)
		}
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	color.HiBlue("\n======= BODY =======")
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "  ")
	if error != nil {
		color.Red("JSON parse error: %v", error)
		return
	}
	color.HiWhite("%v", string(prettyJSON.Bytes()))
}

func main() {
	port := ":8090"
	fmt.Println("Listening on port ", port)

	http.HandleFunc("/", Printer)

	http.ListenAndServe(port, nil)
}
