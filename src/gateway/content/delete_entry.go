package content

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"time"
)

func handleDelete(w http.ResponseWriter, r *http.Request, users string) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := c.Delete(ctx, &pb.UserPostRequest{
		User:   name,
		PostId: postId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}
