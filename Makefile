BIN := p2

insta:
	gh extension install shuntaka9576/gh-p2
remove:
	gh extension remove $(BIN)
build:
	goreleaser release --snapshot --rm-dist --skip-publish

.PHONY: build
