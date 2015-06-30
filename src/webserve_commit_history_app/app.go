package main

import (
	"webserve_commit_history"
)

var _ = webserve_commit_history.Foo

func main() {
	var server webserve_commit_history.GitCommitHistoryWebServer
	server.Port = 81
	var history webserve_commit_history.GitCommitHistory
	history.Init()
	history.Directory = "/home/pi/hinst_static_website"
	server.History = &history
	server.Start()
}
