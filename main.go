package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


var imageMap *ConcurrentImageMultiMap = NewConcurrentImageMultiMap()

func health(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Server Running\n")
}

// func stubResponseJson() string {
// 	// format appears to be Array of string: formatted to be base64 image for <img src tag>
// 	html := "data:image/png;base64,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"
// 	return "{\"data\":[\"" + html + "\"]}"
// } 

// GET will return all images as a JSON object containing an array<string>
// {
//   data:[ "data:image/png;base64,<base64 data>" ]
// } 
//
// Each string is suitable for use as an html <img src={that string}> >
// as it contains the image format (PNG ONLY IN THIS CODE) plus 
// the base64 picture data
type ApiImagesResult struct {
	Data []string `json:"data"`
}

func fetchImages(w http.ResponseWriter, r *http.Request) {
    team := r.PathValue("team")

	images := imageMap.Get(team)

	result := ApiImagesResult{Data: []string{} }

	for _, image := range images {
		htmlFragment := "data:image/png;base64," + image.File 
		result.Data = append(result.Data, htmlFragment)
	}

	enc := json.NewEncoder(w)
	enc.Encode(result)
	fmt.Println("load", team, len(result.Data))
}

// POST will save a single PNG image (it must be PNG)
// The POST body is a JSON object with two fields:
// {
//    "file": "<base64 bytes of image data only, no mime type>"
//	  "fileName": "<soe unique identifier for the file, any format>"
// }
type ApiPostImage struct {
	File 	 	string `json:"file"`
	FileName	string `json:"fileName"`
}

func saveImage(w http.ResponseWriter, r *http.Request) {
    team := r.PathValue("team")
	body, _ := io.ReadAll(r.Body)


	req := ApiPostImage{}
	if err := json.Unmarshal(body, &req); err != nil {
		panic(err)
    }

	fmt.Println("save", team, req.FileName, len(req.File));

	imageMap.Put (team, Image(req))
}

func applyCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  w.Header().Add("Access-Control-Allow-Origin", "*")
	  w.Header().Add("Access-Control-Allow-Credentials", "true")
	  w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	  w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  
	  next(w, r)
	}
  }

  func corsPreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method != "OPTIONS" {
		return
	}

	http.Error(w, "No Content", http.StatusNoContent)
	
}
  
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("GET /health", health)
  mux.HandleFunc("/", applyCors(corsPreflight))
  mux.HandleFunc("GET /IL/teams/{team}/files", applyCors(fetchImages))
  mux.HandleFunc("POST /IL/teams/{team}/files", applyCors(saveImage))

  http.ListenAndServe("localhost:8090", mux)
}