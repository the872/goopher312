# prerequisite: GOROOT and GOARCH must be defined
include GOPATH/src/github.com/the872/goopher312

# name of the package (library) being built
TARG=goopher312

# source files in package
GOFILES=\
	server.go

package: _obj/$(TARG).a


# create a Go package file (.a)
_obj/$(TARG).a: _go_.$O
	@mkdir -p _obj/$(dir)
	rm -f _obj/$(TARG).a
	gopack grc $@ _go_.$O

install:
  - go get github.com/taotetek/gogopher

serve:	
	server.go

compile:
	go build server.go

