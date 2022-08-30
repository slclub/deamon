package server

import "github.com/slclub/deamon/server/docker"

// default register router
func RegisterRoute(serv *Server) {
	serv.AddServ("DockerServer", new(docker.DockerServer))
}
