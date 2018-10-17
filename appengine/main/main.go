// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"net/http"
	"fmt"
	"encoding/json"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/ml", ml)
	appengine.Main()
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test!")
}

func ml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ml!")

	var data = ""
	var projectid = "dogkimoti"
	var modelid = "ICN4592178324928500759"
	var url = "https://beta-dot-custom-vision.appspot.com/v1beta1/projects/" + projectid + "/locations/us-central1/models/" + modelid + ":predict "

	// params = 
	// payload {
	// 	image {
	// 		imageBytes: image
	// 	}
	// }
	//var params = map[string]interface{}
	var image = map[string]interface{}
	image = { "imageBytes" : data }
	params := { "image" : image }
	//params.payload.image.imageBytes = image

	// jsonをデコードする
	input, err := json.Marshal(params)
	
	fmt.Fprintln(w, "input")
	fmt.Fprintln(w, input)

	//params = {}
	//request = prediction_client.predict(modelid, payload, params)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(input))

	fmt.Fprintln(w, "response")
	fmt.Fprintln(w, response)
}