GOBUILD=go build
GOTEST=go test

# alias for go build & go test cmds

all: clean stop build_server
	./ecommerce

# cmd : make all --> clean, stop, build server & run executable 

clean:
	rm -f ./ecommerce

# force (-f) clean / delete previous built binary files

stop:
	pkill ecommerce || true

# kills / stops all running servers in background (if any)

build_server:
	$(GOBUILD) -v .

# go build -v current dir --> builds all .go files

test:
	cd helper && $(GOTEST) -v .

# cd into helper dir & run all test files


