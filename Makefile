sources = main.go version.go
name = redis-copy
dist/$(name).exe: $(sources)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o dist/$(name).exe

dist/$(name)-osx: $(sources)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/$(name)-osx

dist/$(name)-linux: $(sources)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags '-s' -o dist/$(name)-linux

.PHONY: build tag release clean
build: dist/$(name).exe dist/$(name)-osx dist/$(name)-linux

tag:
	./tag.sh $(VERSION)

release: tag build
	./release.sh $(name) $(VERSION) dist/*

clean :
	-rm -r dist
