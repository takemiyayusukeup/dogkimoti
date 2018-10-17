// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"google.golang.org/appengine"

	"google.golang.org/appengine/taskqueue"
)

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/mltask", mltask)
	http.HandleFunc("/ml", ml)
	appengine.Main()
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test!")
}


func mltask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "mltask!")

	//var projectid = "dogkimoti"
	//var modelid = "ICN4592178324928500759"
	//var url = "https://beta-dot-custom-vision.appspot.com/v1beta1/projects/" + projectid + "/locations/us-central1/models/" + modelid + ":predict "

	// jsonをデコードする
	//input, err := json.Marshal(params)
	
	fmt.Fprintln(w, "params")
	fmt.Fprintln(w, params)
	
	ctx := appengine.NewContext(r)
	task := taskqueue.NewPOSTTask("/ml", map[string][]string{"data": {"aaa"}})
	if _, err := taskqueue.Add(ctx, task, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintln(w, "err")
			fmt.Fprintln(w, err)
			return
	}

	fmt.Fprintln(w, "taskqueue.Add")
}

type Maps map[string]interface{}

func ml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ml!")

	var projectid = "dogkimoti"
	var modelid = "ICN4592178324928500759"
	var url = "https://beta-dot-custom-vision.appspot.com/v1beta1/projects/" + projectid + "/locations/us-central1/models/" + modelid + ":predict "
	data := r.FormValue("data")

	params := Maps{
		"payload" : Maps{
			"image" : Maps{ "imageBytes" : data },
		},
	}
    input, err := json.Marshal(params)
	if err != nil {
		fmt.Fprintln(w, "json.Marshal err")
		fmt.Fprintln(w, err)
		return
	}
	// ctx := appengine.NewContext(r)
	// task := NewJsonPOSTTask(url, params)
	// if _, err := taskqueue.Add(ctx, task, ""); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		fmt.Fprintln(w, "err")
	// 		fmt.Fprintln(w, err)
	// 		return
	// }
    response, err := http.Post(url, "application/json", bytes.NewBuffer(input))
	if err != nil {
		fmt.Fprintln(w, "http.Post err")
		fmt.Fprintln(w, err)
		return
	}

    fmt.Fprintln(w, "response")
    fmt.Fprintln(w, response)
	fmt.Fprintln(w, "ml end")
}

// func NewJsonPOSTTask(path string, params Maps) *taskqueue.Task {
//     h := make(http.Header)
//     h.Set("Content-Type", "application/json")
//     data, _ := json.Marshal(params) // TODO エラー捨ててるYO.
//     return &taskqueue.Task{
//         Path:    path,
//         Payload: []byte(data),
//         Header:  h,
//         Method:  "POST",
//     }
// }