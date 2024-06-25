lsof -ti :80
status=$(echo $?)
if [[ "$status"=="0" ]]; then
    echo "Killing previous server"
    kill -9 $(lsof -ti :80)
else
    echo "No server to kill"

JENKINS_NODE_COOKIE=dontKillMe ./Jenkins &
fi
