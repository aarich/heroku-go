BIN_DIR=$(shell pwd)/bin

build:
	go build -o $(BIN_DIR)/heroku-go

local:
	heroku local

deploy:
	git push heroku master

clean:
	rm -rf $(BIN_DIR)
