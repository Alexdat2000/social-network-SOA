package content

import (
	"context"
	"log"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"time"
)

func HandleLike(w http.ResponseWriter, r *http.Request, users string) {
	log.Printf("Received LIKE")
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

	ans, err := c.LikePost(ctx, &pb.UserPostRequest{
		User:   name,
		PostId: postId,
	})
	if err == nil && ans != nil && ans.GetSuccessful() == true {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
