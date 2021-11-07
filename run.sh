export SERVER_PORT=1323
export MONGO_HOST=localhost
export MONGO_PORT=27017
export MONGO_SCHEMA=mongodb
export MONGO_DATABASE_NAME=nbasimulator
export TEAM_COLLECTION=team
export PLAYER_COLLECTION=player

go get github.com/labstack/echo/v4
go get go.mongodb.org/mongo-driver/mongo
go mod tidy

docker run --name mongodb -d -p 27017:27017 mongo

go run src/main.go