SRC=./cmd
BINARY=imgcat
BINARY_SRC=$(SRC)/imgcat
SOURCES=$(shell find $(BINARY_SRC) -name "*.go")

.DEFAULT_GOAL := all

$(BINARY): $(SOURCES)
	go build $(BINARY_SRC)

all: $(BINARY) test

test:
	./$(BINARY) -url "https://images.pexels.com/photos/2558605/pexels-photo-2558605.jpeg?cs=srgb&dl=pexels-anel-rossouw-2558605.jpg&fm=jpg" -height 25
