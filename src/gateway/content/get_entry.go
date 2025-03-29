package content

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"time"
)

func handleGet(w http.ResponseWriter, r *http.Request, users string) {
	log.Printf("Received GET")
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ans, err := c.Get(ctx, &pb.UserPostRequest{
		User:   name,
		PostId: postId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil && err.Error() == "no access" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}
