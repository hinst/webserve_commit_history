#!/bin/bash
appHome=/home/pi/webserve_commit_history/src/webserve_commit_history_app
if [ "$1" = "start" ]
then
	cd "$appHome"
	nohup ./webserve_commit_history_app >normal.log 2>error.log &
	echo $! > app.pid
fi
if [ "$1" == "stop" ]
then
	cd "$appHome"
	pid=$(<app.pid)
	kill -SIGINT $pid
fi


