docker stop bot_container
docker rm bot_container
docker rmi bot
docker build . -t bot
docker run  --name bot_container -p 9096:9096 -d bot
docker logs -f bot_container