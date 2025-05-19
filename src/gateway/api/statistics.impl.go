package api

import "net/http"

// Statistics

func (s Server) GetPostsPostIdStats(w http.ResponseWriter, r *http.Request, postId int) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetPostsPostIdStatsDaily(w http.ResponseWriter, r *http.Request, postId int, params GetPostsPostIdStatsDailyParams) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetPostsTop10(w http.ResponseWriter, r *http.Request, params GetPostsTop10Params) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetUsersTop10(w http.ResponseWriter, r *http.Request, params GetUsersTop10Params) {
	//TODO implement me
	panic("implement me")
}
