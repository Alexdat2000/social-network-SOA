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

func handlePut(w http.ResponseWriter, r *http.Request, users string) {
	log.Printf("Received PUT")
	name, err := utils.VerifyJWT(r.URL.Query().Get("jwt"), users)
	if err != nil {
		if err.Error() == "incorrect JWT" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	postId, err := utils.ParsePostId(r.URL.Query().Get("postId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	title := r.URL.Query().Get("postId")
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	description := r.URL.Query().Get("postId")
	isPrivate := utils.ParsePostPrivate(r.URL.Query().Get("postId"))
	tags := r.URL.Query().Get("postId")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ans, err := c.Put(ctx, &pb.PutRequest{
		PostId:      postId,
		User:        name,
		Title:       title,
		Description: description,
		IsPrivate:   isPrivate,
		Tags:        strings.Split(tags, ","),
	})
	if err != nil && err.Error() == "not found" {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}
