build:
	go build -o dist/main -gcflags all=-N
	#exec dist/main -- up mongodb://manager:password@localhost:27017/property?authSource=admin file://./example/migration