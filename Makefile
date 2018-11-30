GO ?= go
EXECUTABLE := go-waveform

all: build

build: $(EXECUTABLE)

$(EXECUTABLE):
	$(GO) build -v -o bin/$@ ./cmd
