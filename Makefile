all: local alpine linux

local:
	./build.sh

alpine:
	docker run -ti -v $(CURDIR):/$(CURDIR)/ --workdir /$(CURDIR)/ qnib/alpn-go-dev ./build.sh

linux:
	docker run -ti -v $(CURDIR):/$(CURDIR)/ --workdir /$(CURDIR)/ qnib/golang ./build.sh
  
