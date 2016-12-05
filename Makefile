GO=go
BIN=trajectory

trajectory:
	$(GO) build 
	cp main/main $(BIN)

clean:
	rm $(BIN)
