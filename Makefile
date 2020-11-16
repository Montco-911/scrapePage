
docker-build:
	docker build  --no-cache -t quay.io/mchirico/scrape:test -f Dockerfile .

push:
	docker push quay.io/mchirico/scrape:test

build:
	go build -v .

run:
	docker run --name scrape --rm -it -p 3000:3000  quay.io/mchirico/scrape:test




