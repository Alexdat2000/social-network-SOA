package api

import (
	"io"
	"net/http"
)

func redirectToUserService(w http.ResponseWriter, r *http.Request) {
	targetURL := "http://users:8080" + r.URL.RequestURI()
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header = r.Header.Clone()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

// Users

func (s Server) PostUsersLogin(w http.ResponseWriter, r *http.Request) {
	redirectToUserService(w, r)
}

func (s Server) PostUsers(w http.ResponseWriter, r *http.Request) {
	redirectToUserService(w, r)
}

func (s Server) PatchUsers(w http.ResponseWriter, r *http.Request) {
	redirectToUserService(w, r)
}

func (s Server) GetUsersUsername(w http.ResponseWriter, r *http.Request, username string) {
	redirectToUserService(w, r)
}
