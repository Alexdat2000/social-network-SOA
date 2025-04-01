package content

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	pb "soa/gateway/content_grpc"
	"strconv"
	"time"
)

func HandleEntry(w http.ResponseWriter, r *http.Request, users, cont string) {
	if r.Method == http.MethodGet {
		handleGet(w, r, users)
	} else if r.Method == http.MethodPost {
		handlePost(w, r, users)
	} else if r.Method == http.MethodPut {
		handlePut(w, r, users)
	} else if r.Method == http.MethodDelete {
		handleDelete(w, r, users)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleList(w http.ResponseWriter, r *http.Request, users, cont string) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ans, err := c.GetPosts(ctx, &pb.GetPostsRequest{
		Page: uint32(page),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}
