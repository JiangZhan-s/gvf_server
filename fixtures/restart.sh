docker stop $(docker ps -aq)
echo "stop all docker"
docker rm $(docker ps -aq)
docker container stop $(docker container ls -aq)

docker volume prune
