package webserve_commit_history

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

type GitCommitHistory struct {
	Directory    string
	CountOfItems int
	TextFormat   string
}

const DefaultCountOfItems = 10
const DefaultTextFormat = "%an @ %ad %n%B"

func (this *GitCommitHistory) Init() {
	this.CountOfItems = DefaultCountOfItems
	this.TextFormat = DefaultTextFormat
}

func (this *GitCommitHistory) GetText() string {
	var text = ""
	var gitDirectory = this.Directory + "/.git"
	var _, gitStatResult = os.Stat(gitDirectory)
	if gitStatResult == nil {
		var command = exec.Command("/usr/bin/git", "log", "--pretty=format:"+this.TextFormat, "--max-count="+strconv.Itoa(this.CountOfItems))
		command.Dir = this.Directory
		var output, outputResult = command.StdoutPipe()
		if outputResult == nil {
			var startResult = command.Start()
			if startResult == nil {
				text = ReaderToString(output)
				command.Wait()
			} else {
				fmt.Println(startResult)
			}
		} else {
			text = "Error: could not obtain output"
		}
	} else {
		text = "Error: no " + gitDirectory
	}
	return text
}

func ReaderToString(reader io.Reader) string {
	const chunkLength = 1024
	var chunk []byte = make([]byte, chunkLength)
	var text bytes.Buffer
	for {
		var readLength, readResult = reader.Read(chunk)
		var chunkToWrite = chunk[:readLength]
		text.Write(chunkToWrite)
		if readResult == io.EOF {
			break
		}
	}
	return text.String()
}

var Foo = 0
