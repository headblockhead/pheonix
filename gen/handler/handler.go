package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"time"

	gen "github.com/headblockhead/phoenix"
	"github.com/nfnt/resize"
)

type Response struct {
	Frames            [][]byte
	Objection         []byte
	Seed              int
	ObjectionLocation int
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var resp Response
	var queries = r.URL.Query()
	seed := int(time.Now().UTC().UnixNano())
	if queries.Get("seed") != "" {
		var err error
		seed, err = strconv.Atoi(queries.Get("seed"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid seed", http.StatusBadRequest)
			return
		}
	}
	if queries.Get("fullQuality") == "" {
		http.Error(w, "Missing fullQuality query", http.StatusBadRequest)
		return
	}
	resp.Seed = seed
	frames, objection, ObjectionLocation, err := gen.Generate(seed)
	resp.ObjectionLocation = ObjectionLocation
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to generate: %v", err), http.StatusInternalServerError)
		return
	}
	if objection != nil {
		o := new(bytes.Buffer)
		err = jpeg.Encode(o, objection, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Printf("failed to encode the objection: %v\n", err)
		}
		resp.Objection = o.Bytes()
	}
	if queries.Get("fullQuality") == "false" {
		for i := 0; i < len(frames); i++ {
			newImage := resize.Resize(220, 140, frames[i], resize.Lanczos3)
			frames[i] = newImage
		}
	}
	resp.Frames = make([][]byte, len(frames))
	for i := 0; i < len(frames); i++ {
		b := new(bytes.Buffer)
		err = jpeg.Encode(b, frames[i], &jpeg.Options{Quality: 100})
		if err != nil {
			log.Printf("failed to encode frame %d: %v\n", i, err)
		}
		resp.Frames[i] = b.Bytes()
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	enc.Encode(resp)
}
