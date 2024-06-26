process=$(lsof -ti :80 2>/dev/null)
if [[ ! -z "$process" ]]; then
    echo "Killing previous server"
    kill -9 $process
else
    echo "No server to kill"
fi

JENKINS_NODE_COOKIE=dontKillMe ./Jenkins &
