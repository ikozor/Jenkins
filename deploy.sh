kill $(lsof -ti :80)
JENKINS_NODE_COOKIE=dontKillMe ./Jenkins &
