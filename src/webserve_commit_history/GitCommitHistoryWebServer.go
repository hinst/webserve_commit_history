package webserve_commit_history

import (
	"net"
	"net/http"
	"strconv"
	"sync"
)

type GitCommitHistoryWebServer struct {
	History          *GitCommitHistory
	Port             int
	Server           *http.Server
	listener         net.Listener
	log              Log
	requestWaitGroup sync.WaitGroup
}

func (this *GitCommitHistoryWebServer) Init() {
	this.log.TypeName = "GitCommitHistoryWebServer"
}

func (this *GitCommitHistoryWebServer) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	this.requestWaitGroup.Add(1)
	writeString := func(theString string) {
		responseWriter.Write([]byte(theString))
	}
	writeString(this.History.GetText())
	responseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
	this.requestWaitGroup.Done()
}

func (this *GitCommitHistoryWebServer) createListener() {
	var listener, listenResult = net.Listen("tcp", ":"+strconv.Itoa(this.Port))
	if listenResult == nil {
		this.listener = listener
	} else {
		this.log.Write("Error: " + listenResult.Error())
	}
}

func (this *GitCommitHistoryWebServer) Start() {
	this.requestWaitGroup.Add(1)
	this.createListener()
	this.Server = &http.Server{}
	this.Server.Handler = this
	var serveResult = this.Server.Serve(this.listener)
	if serveResult != nil {
		this.log.Write(serveResult.Error())
	}
	this.requestWaitGroup.Done()
}

func (this *GitCommitHistoryWebServer) Stop() {
	this.listener.Close()
	this.requestWaitGroup.Wait()
}
