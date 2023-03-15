OUTPUT_FOLDER=bin
FILENAME=ipam

all: build
 
build:
	go build -o ${OUTPUT_FOLDER}/${FILENAME} main.go

install:
	cp ${OUTPUT_FOLDER}/${FILENAME} /usr/local/bin/${FILENAME}

clean:
	go clean
	rm -r ${OUTPUT_FOLDER}/
