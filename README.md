
![Go](https://github.com/Montco-911/scrapePage/workflows/Go/badge.svg)
# scrapePage


## Docker
```
docker run --rm -it quay.io/mchirico/scrape

# For non daemon
docker  run --rm -it --name activinc -a stdout -a stderr  quay.io/mchirico/scrape > activInc

docker  run --rm -it --name activinc  -d  quay.io/mchirico/scrape
docker logs  activinc

docker attach activinc

# To detach the tty without exiting the shell,
# use the escape sequence Ctrl-p + Ctrl-q

docker exec -ti  kind-worker /bin/bash
```







