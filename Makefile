sources = main.go

.PHONY: build release clean
build: dist/redis-copy.exe dist/redis-copy-osx dist/redis-copy-linux

dist/redis-copy.exe: $(sources)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o dist/redis-copy.exe

dist/redis-copy-osx: $(sources)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/redis-copy-osx

dist/redis-copy-linux: $(sources)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/redis-copy-linux

release: build
	./release.sh $(VERSION)

clean :
	-rm -r dist
