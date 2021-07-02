docker rm $(docker stop $(docker ps -aq))
echo y|docker network prune
echo y|docker volume prune

docker rmi $(docker images | grep dev-* | awk '{print $3}')

ps -aux|grep 'python3 manage.py'|grep -v grep|awk '{print$2}'|xargs kill -9



# ps -aux|grep 'python3 manage'|grep -v grep|awk '{print$2}'|xargs kill -9