all: clean build

clean:
	rm -rf cmd/*

build:
	go build -o cmd/itcrowdipsum