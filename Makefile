GIT_TAG=$(shell git describe --tags --abbrev=0)
GIT_HASH=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date '+%F-%H:%M:%S')

all: info help

info:
	@echo "[leave it us make]\nbuild information : ${GIT_TAG} - ${GIT_HASH} (${BUILD_DATE})"

help:
	@echo "명령어 목록 : build, info, serve, clean, build-frontend, build-backend"

build-backend: backend
	cd backend && go build -v -x -o ../output/ -ldflags "-X main.gitTag=${GIT_TAG} -X main.gitHash=${GIT_HASH} -X main.buildDate=${BUILD_DATE}"

debug-backend: backend
	cd backend && go build -v -x -o ../output/ -ldflags "-X main.gitTag=${GIT_TAG} -X main.gitHash=${GIT_HASH} -X main.buildDate=${BUILD_DATE} -X main.displayDebug=true"

build-frontend: frontend/src
	cd frontend && npm run build
	cp -r frontend/dist output/static

serve: clean build-frontend debug-backend
	cd output && ./backend
	
clean:
	rm -rf ./output
	mkdir output

build: info clean build-backend build-frontend

install:
	cd frontend && npm update

