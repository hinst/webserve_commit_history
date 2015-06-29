package main

import (
	"fmt"
	"webserve_commit_history"
)

var _ = webserve_commit_history.Foo

func main() {
	fmt.Println(webserve_commit_history.GetCommitHistoryAsText("/home/pi/hinst_static_website"))
}
