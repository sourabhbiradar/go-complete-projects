# gRPC + Docker + MongoDB
# blog project

install docker

`cmd` : docker --version 
`cmd` : docker-compose --version
`cmd` : sudo systemctl start docker
`cmd` : systemctl status docker

`docker-compose` 
.yml files
mongo
mongo_express  // UI for mongoDB
`cmd` : docker-compose up

mongoDB connection string:
mongodb://admin:pasAdmin@mongo:27017/

install mongo-driver for GO
`cmd` : go get mongodb.org/mongo-drive/mongo

`mongo.Collection` : mongo client >> mongoDB >> mongo.Collection

run make serverrun & make clientrun , then
`visit` : localhost:8081 >> admin : pass >> click `blog`





