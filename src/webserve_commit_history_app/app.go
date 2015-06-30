package main

import (
	"fmt"
	"webserve_commit_history"
)

var _ = webserve_commit_history.Foo

func main() {
	var history webserve_commit_history.GitCommitHistory
	history.Directory = "/home/pi/hinst_static_website"
	fmt.Println(history.GetText())
}
