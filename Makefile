.DEFAULT_GOAL := all

DIST=./bin
SRC=./cmd
BINARY=$(DIST)/imgcat
BINARY_SRC=$(SRC)/imgcat
SOURCE=$(shell find $(BINARY_SRC) -name "*.go")

$(BINARY): $(SOURCE)
	go build -trimpath -o $(BINARY) $(BINARY_SRC)

all: $(BINARY) test

test:
	$(BINARY) -url "https://images.pexels.com/photos/2558605/pexels-photo-2558605.jpeg?cs=srgb&dl=pexels-anel-rossouw-2558605.jpg&fm=jpg" -height 25