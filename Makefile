
.PHONY: debug run

run:
	cd test; go run testrun.go

debug:
	go tool cgo *.go
