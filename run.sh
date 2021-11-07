export SERVER_PORT=1323
export MONGO_HOST=mongo
export MONGO_PORT=27017
export MONGO_SCHEMA=mongodb
export MONGO_DATABASE_NAME=nbasimulator
export GAME_COLLECTION=game
export TEAM_COLLECTION=team
export PLAYER_COLLECTION=player
export PLAYER_GAME_INFO_COLLECTION=player_game_info

go get github.com/labstack/echo/v4
go get go.mongodb.org/mongo-driver/mongo
go get github.com/go-resty/resty/v2
go mod tidy

docker run --name mongodb -d -p 27017:27017 mongo

go run src/main.go