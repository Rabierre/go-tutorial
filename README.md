# How To Install
## Linux
    mkdir ~/download
    cd ~/download
    wget https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz

    tar -C /usr/local -xzf go1.6.2.linux-amd64.tar.gz

    export PATH=$PATH:/usr/local/go/bin

## Other OS
https://golang.org/doc/install

## Test go install
    go


### Should print this
Go is a tool for managing Go source code.

Usage:

	go command [arguments]

The commands are:

	build       compile packages and dependencies
	clean       remove object files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         run go tool fix on packages
	fmt         run gofmt on package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

	c           calling between Go and C
	buildmode   description of build modes
	filetype    file types
	gopath      GOPATH environment variable
	environment environment variables
	importpath  import path syntax
	packages    description of package lists
	testflag    description of testing flags
	testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
