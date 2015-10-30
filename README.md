# GoWebApp

A basic MVC web app in Golang, but with all of the dependencies vendored using [GoDep](https://github.com/tools/godep).

### Usage
```
cd ~/go/src
git clone https://github.com/cleesmith/gowebapp.git
cd gowebapp
go build
maybe, edit config/config.json to change various app settings
sudo ./gowebapp
	- why sudo? coz the default is port 80, so edit config/config.json to change it
```

### ToDo's
1. switch from [Foundation](http://foundation.zurb.com/) to [Bootstrap](http://getbootstrap.com/)
1. use [BoltDB](https://github.com/boltdb/bolt) instead of [MySQL](https://www.mysql.com/) or [SQLite](https://www.sqlite.org/)
1. reduce the app's footprint to the bare minimum
	* remove the unused dependencies, well, for my typical usage anyways
		* such as the user registration and recaptcha stuff
		* typically, my apps define an **admin** user who manages other users
	* currently on a Mac the binary file size is 16,370,184
		* which is not too big considering that is everything -- an all-in-one easy to deploy file
1. cross compile doesn't work on a Mac because of the included C file for SQLite
	* ```GOOS=linux GOARCH=amd64 go build``` ... yields this error message:
		```shared/database/database.go:9:2: C source files not allowed when not using cgo or SWIG: sqlite3-binding.c```
	* and ```CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build``` ... yields this error message:

	      ```
        # runtime/cgo
        ld: unknown option: --build-id=none
        clang: error: linker command failed with exit code 1 (use -v to see invocation)
	      ```
	* see:
		* https://github.com/mattn/go-sqlite3/issues/106
		* https://github.com/mattn/go-sqlite3/issues/217
		* http://www.limitlessfx.com/cross-compile-golang-app-for-windows-from-linux.html
	* solutions:
		* don't cross compile, just install Go on the specific platform then build the app
		* don't use MySQL or SQLite for data storage
		* why? [KISS](https://en.wikipedia.org/wiki/KISS_principle)

### Godep

```
Go 1.5:
  - don't do: godep save -r
  - ~/.bash_profile contains:
    export GO15VENDOREXPERIMENT=1
    ... this causes godep to use the vendor folder for dependencies

godep save
  - creates Godeps folder with Godeps.json in there
  - creates vendor folder and copies dependencies into it

go build
  - which will use the vendor folder instead of ~/go/src/whatever
```

This app works as expected, so by vendoring the dependencies the app should continue to work
independent of any future changes to the dependenies on github or other repos.

***
***

> See the original project at [GoWebApp](https://github.com/josephspurrier/gowebapp).

***
***
