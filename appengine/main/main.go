// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"

	"google.golang.org/appengine"

	"google.golang.org/appengine/taskqueue"
)

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/ml", ml)
	appengine.Main()
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test!")
}

type M map[string]interface{}

func ml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ml!")

	var data = ""
	var projectid = "dogkimoti"
	var modelid = "ICN4592178324928500759"
	var url = "https://beta-dot-custom-vision.appspot.com/v1beta1/projects/" + projectid + "/locations/us-central1/models/" + modelid + ":predict "

	params := M{
		"payload" : M{
			"image" : M{ "imageBytes" : data },
		},
	}

	// jsonをデコードする
	//input, err := json.Marshal(params)
	
	fmt.Fprintln(w, "params")
	fmt.Fprintln(w, params)
	
	ctx := appengine.NewContext(r)
	task := NewJsonPOSTTask(url, params)
	if _, err := taskqueue.Add(ctx, task, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintln(w, "err")
			fmt.Fprintln(w, err)
			return
	}

	fmt.Fprintln(w, "taskqueue.Add")
}

func NewJsonPOSTTask(path string, params url.Values) *Task {
    h := make(http.Header)
    h.Set("Content-Type", "application/json")
    data, _ := json.Marshal(params) // TODO エラー捨ててるYO.
    return &Task{
        Path:    path,
        Payload: []byte(data),
        Header:  h,
        Method:  "POST",
    }
}