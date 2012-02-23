all: dgrid

dgrid: *.go
	go tool 6g -o dgrid.6 *.go
	go tool 6l -o dgrid dgrid.6
