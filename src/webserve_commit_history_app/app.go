package main

import (
	"webserve_commit_history"
)

var _ = webserve_commit_history.Foo

func main() {
	var server webserve_commit_history.GitCommitHistoryWebServer
	var log webserve_commit_history.Log
	server.Init()
	server.Port = 81
	var history webserve_commit_history.GitCommitHistory
	history.Init()
	history.Directory = "/home/pi/hinst_static_website"
	history.TextFormat = "<b>%an</b> @ <b>%ad</b>%n<br>%B<br><br>%n%n"
	server.History = &history
	webserve_commit_history.InstallShutdownReceiver(
		func() {
			log.Write("Now exiting app...")
			server.Stop()
			log.Write("Exited")
		})
	server.Start()
}
