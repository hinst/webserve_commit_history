package webserve_commit_history

import (
	"fmt"
)

type Log struct {
	TypeName string
}

func (this *Log) Write(text string) {
	fmt.Println(this.TypeName + ": " + text)
}
