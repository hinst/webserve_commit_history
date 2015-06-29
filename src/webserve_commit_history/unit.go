package webserve_commit_history

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var Foo = 0

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

func GetCommitHistoryAsText(directory string) string {
	var text = ""
	var gitDirectory = directory + "/.git"
	var _, gitStatResult = os.Stat(gitDirectory)
	if gitStatResult == nil {
		var command = exec.Command("/usr/bin/git", "log", "--pretty=format:%an @ %ad %n%B", "--max-count=10")
		command.Dir = directory
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
