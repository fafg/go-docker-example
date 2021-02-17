#!/usr/bin/env bash

if [[ "$( docker container inspect -f '{{.State.Running}}' sonarqube )" == "false" ]];
then
  docker start sonarqube
elif [[ "$( docker container inspect -f '{{.State.Running}}' sonarqube )" == "true" ]];
then
  docker ps
else
  docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube
fi

echo "waiting for sonarqube starts..."
wget -q -O - "$@" http://localhost:9000 | awk '/STARTING/{ print $0 }'

STATUS="$(wget -q -O - "$@" http://localhost:9000 | awk '/UP/{ print $0 }')"

while [ -z "$STATUS" ]
do
	sleep 2s
	STATUS=$(wget -q -O - "$@" http://localhost:9000 | awk '/UP/{ print $0 }')
	echo -e "."
done

echo -e "\\n$STATUS"