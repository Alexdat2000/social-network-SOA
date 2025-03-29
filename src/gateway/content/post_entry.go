package content

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"strings"
	"time"
)

func handlePost(w http.ResponseWriter, r *http.Request, users string) {
	log.Printf("Received POST")
	name, err := utils.VerifyJWT(r.URL.Query().Get("jwt"), users)
	if err != nil {
		if err.Error() == "incorrect JWT" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("Error verifying JWT %v", err)
		return
	}
	title := r.URL.Query().Get("title")
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := r.URL.Query().Get("description")
	isPrivate := utils.ParsePostPrivate(r.URL.Query().Get("isPrivate"))
	tags := r.URL.Query().Get("tags")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ans, err := c.Post(ctx, &pb.PostRequest{
		User:        name,
		Title:       title,
		Description: description,
		IsPrivate:   isPrivate,
		Tags:        strings.Split(tags, ","),
	})
	if err != nil {
		log.Printf("Error POST grpc: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}
