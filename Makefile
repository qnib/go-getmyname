all: local alpine linux

local:
	./build.sh

alpine:
	docker run -ti -v $(CURDIR):/$(basename $(pwd))/ --workdir /$(basename $(pwd))/ qnib/alpn-go-dev ./build.sh

linux:
	docker run -ti -v $(CURDIR):/$(basename $(pwd))/ --workdir /$(basename $(pwd))/ qnib/golang ./build.sh
  
