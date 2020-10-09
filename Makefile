GIT_TAG=$(shell git describe --tags --abbrev=0)
GIT_HASH=$(shell git rev-parse HEAD)
BUILD_DATE=$(shell date '+%F-%H:%M:%S')

all: info help

info:
	@echo "[leave it us make]\nbuild information : ${GIT_TAG} - ${GIT_HASH} (${BUILD_DATE})"

help:
	@echo "명령어 목록 : build, info, serve, clean, build-frontend, build-backend"

build-backend: backend
	cd backend && go build -v -x -o ../output/

build-frontend: frontend/src
	cd frontend && npm run build
	cp -r frontend/dist output/static

serve: 
	cd output && ./backend

clean:
	rm -rf ./output
	mkdir output

build: info clean build-backend build-frontend

