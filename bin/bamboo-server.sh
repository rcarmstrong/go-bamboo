#!/bin/sh

. "$(dirname "$0")/checkenv.sh"

which docker > /dev/null
if [ $? -ne 0 ]; then
    echo "Test Bamboo server requires Docker. Please install Docker."
    exit 1
fi

docker inspect local-bamboo > /dev/null 2>&1
if [ $? == 1 ];then
  docker run -d -p 8085:8085 -v ${PWD}/data/bamboo:/var/atlassian/bamboo --name local-bamboo cptactionhank/atlassian-bamboo 
  echo "Be sure to set up Bamboo at http://localhost:8085/"
fi