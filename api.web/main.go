package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/opiskull/red-panda/api.web/server"
	profile "github.com/opiskull/red-panda/service.profile/client"
)

func main() {
	var (
		port              = flag.String("port", "5000", "The server port")
		profileServerAddr = flag.String("profile_server_addr", "127.0.0.1:10001", "The Pofile server address in the format of host:port")
		//authServerAddr    = flag.String("auth_server_addr", "127.0.0.1:10001", "The Auth server address in the format of host:port")
		//geoServerAddr     = flag.String("geo_server_addr", "127.0.0.1:10002", "The Geo server address in the format of host:port")
		//rateServerAddr    = flag.String("rate_server_addr", "127.0.0.1:10004", "The Rate Code server address in the format of host:port")
	)
	flag.Parse()

	profileClient, err := profile.NewClient(*profileServerAddr)
	if err != nil {
		log.Fatal("ProfileClient error:", err)
	}
	defer profileClient.Close()

	server := server.NewServer(profileClient)

	http.HandleFunc("/", server.GetProfile)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
