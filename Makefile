.DEFAULT_GOAL := all

imgcat: main.go
	go build -o $@

build: imgcat

all: build test

test:
	./imgcat -u "https://images.pexels.com/photos/2558605/pexels-photo-2558605.jpeg?cs=srgb&dl=pexels-anel-rossouw-2558605.jpg&fm=jpg" -r 25
