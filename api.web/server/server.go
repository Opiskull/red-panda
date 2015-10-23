package server

import profile "github.com/opiskull/red-panda/service.profile/client"

type Server struct {
	profileClient *profile.Client
}

func NewServer(profileClient *profile.Client) Server {
	return Server{
		profileClient: profileClient,
	}
}
