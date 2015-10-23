package server

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

func (s Server) GetProfile(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	profileReply, err := s.profileClient.GetProfile(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(profileReply)

	w.Write(body)
}
