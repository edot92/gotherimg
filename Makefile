NAME = static

assets: clean
	go-bindata -pkg statik  -debug -o engine//static//$(NAME).go static/...
clean:
	rm -fr engine//static//$(NAME).go

.PHONY: all assets clean