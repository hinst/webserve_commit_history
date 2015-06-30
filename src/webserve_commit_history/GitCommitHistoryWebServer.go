package webserve_commit_history

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

type GitCommitHistoryWebServer struct {
	History  *GitCommitHistory
	Port     int
	Server   *http.Server
	Listener net.Listener
	log      Log
}

func (this *GitCommitHistoryWebServer) Init() {
	this.log.TypeName = "GitCommitHistoryWebServer"
}

func (this *GitCommitHistoryWebServer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	writeString := func(theString string) {
		responseWriter.Write([]byte(theString))
	}
	writeString("lel")
}

func (this *GitCommitHistoryWebServer) createListener() {
	var listener, listenResult = net.Listen("tcp", ":"+strconv.Itoa(this.Port))
	if listenResult == nil {
		this.Listener = listener
	} else {
		this.log.Write("Error: " + listenResult.Error())
	}
}

func (this *GitCommitHistoryWebServer) Start() {
	this.createListener()
	this.Server = &http.Server{}
	this.Server.Handler = this
	var serveResult = this.Server.Serve(this.Listener)
	if serveResult != nil {
		fmt.Println(serveResult)
	}
}
