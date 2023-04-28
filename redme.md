mahar@DESKTOP-APF469J MINGW64 /c/go
$ cd golang-test/

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go mod init
go: cannot determine module path for source directory C:\go\golang-test (outside GOPATH, module path must be specified)

Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go init
go init: unknown command
Run 'go help' for usage.

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go mod init github.com/mahardhika21/golang-test.git
go: creating new go.mod: module github.com/mahardhika21/golang-test.git

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:5:2: package helper/hello_world is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\helper\hello_world)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:5:2: package hello_world is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\hello_world)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:5:2: package golang-test/hello_world is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\golang-test\hello_world)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:6:2: package golang-test/hello_world is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\golang-test\hello_world)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:6:2: package golang-test/helper is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\golang-test\helper)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ go run main.go 
main.go:6:2: no required module provides package github.com/golang-test/helper; to add it:
        go get github.com/golang-test/helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test
$ git init
Initialized empty Git repository in C:/go/golang-test/.git/

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git add remote add origin https://github.com/mahardhika21/golang-test.git
fatal: pathspec 'remote' did not match any files

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git remote add origin https://github.com/mahardhika21/golang-test.git

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git add -A
warning: in the working copy of 'go.mod', LF will be replaced by CRLF the next time Git touches it

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git add go.mod 

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git add helper/

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git add -A

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git commit -m "init"
[master (root-commit) 7f62eb0] init
 3 files changed, 19 insertions(+)
 create mode 100644 go.mod
 create mode 100644 helper/hello_world.go
 create mode 100644 main.go

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ git push origin master
Enumerating objects: 6, done.
Counting objects: 100% (6/6), done.
Delta compression using up to 8 threads
Compressing objects: 100% (5/5), done.
Writing objects: 100% (6/6), 573 bytes | 191.00 KiB/s, done.
Total 6 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/mahardhika21/golang-test.git
 * [new branch]      master -> master

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get
github.com/mahardhika21/golang-test.git imports
        https://github.com/mahardhika21/golang-test/helper: malformed import path "https://github.com/mahardhika21/golang-test/helper": double slash

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get
go: downloading github.com/mahardhika21/golang-test v0.0.0-20230418202934-7f62eb0ad492
github.com/mahardhika21/golang-test.git imports
        github.com/mahardhika21/golang-test/helper: cannot find module providing package github.com/mahardhika21/golang-test/helper        

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get
github.com/mahardhika21/golang-test.git imports
        github.com/mahardhika21/golang-test/helper: cannot find module providing package github.com/mahardhika21/golang-test/helper        

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ ^C

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper; to add it:
        go get github.com/mahardhika21/golang-test/helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper/hello_world; to add it:
        go get github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get
github.com/mahardhika21/golang-test.git imports
        github.com/mahardhika21/golang-test/helper/hello_world: cannot find module providing package github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper/hello_world; to add it:
        go get github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: malformed import path "/helper/hello_world": empty path element

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: "./helper/hello_world" is relative, but relative import paths are not supported in module mode

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: package helper/hello_world is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\helper\hello_world)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: "../helper/hello_world" is relative, but relative import paths are not supported in module mode

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper/hello_world; to add it:
        go get github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get github.com/mahardhika21/golang-test/helper/hello_world
go: module github.com/mahardhika21/golang-test@upgrade found (v0.0.0-20230418202934-7f62eb0ad492), but does not contain package github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper/hello_world; to add it:
        go get github.com/mahardhika21/golang-test/helper/hello_world

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper; to add it:
        go get github.com/mahardhika21/golang-test/helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get github.com/mahardhika21/golang-test/helper
go: github.com/mahardhika21/golang-test@v0.0.0-20230418202934-7f62eb0ad492: parsing go.mod:
        module declares its path as: github.com/mahardhika21/golang-test.git
                but was required as: github.com/mahardhika21/golang-test

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go
main.go:5:2: no required module provides package github.com/mahardhika21/golang-test/helper; to add it:
        go get github.com/mahardhika21/golang-test/helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:6:2: no required module provides package github.com/mahardhika21/golang-test/helper; to add it:
        go get github.com/mahardhika21/golang-test/helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get helper/
go: malformed module path "helper": missing dot in first path element

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get golang-test/helper
go: malformed module path "golang-test/helper": missing dot in first path element

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get helper/
go: malformed module path "helper": missing dot in first path element

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get helper/.
go: malformed module path "helper": missing dot in first path element

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go get ./helper

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:6:2: package helper is not in GOROOT (C:\laragon\bin\go\go1.18.5.windows-amd64\src\helper)

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
main.go:6:2: "./helper" is relative, but relative import paths are not supported in module mode

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go test
# github.com/mahardhika21/golang-test.git
.\main.go:4:2: imported and not used: "fmt"

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ cd helper/

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.320s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ cd ..

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go mod init github.com/mahardhika21/golang-test.git/helper
go: C:\go\golang-test\go.mod already exists

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
# command-line-arguments
.\main.go:5:2: imported and not used: "github.com/mahardhika21/golang-test.git/helper"

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go run main.go 
Hello zul

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$  cd helper/

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test -v
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.296s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test -v
=== RUN   TestHelloWorld
--- FAIL: TestHelloWorld (0.00s)
panic: result is not pass [recovered]
        panic: result is not pass

goroutine 6 [running]:
testing.tRunner.func1.2({0x60d8a0, 0x65f6e0})
        C:/laragon/bin/go/go1.18.5.windows-amd64/src/testing/testing.go:1389 +0x24e
testing.tRunner.func1()
        C:/laragon/bin/go/go1.18.5.windows-amd64/src/testing/testing.go:1392 +0x39f
panic({0x60d8a0, 0x65f6e0})
        C:/laragon/bin/go/go1.18.5.windows-amd64/src/runtime/panic.go:838 +0x207
github.com/mahardhika21/golang-test.git/helper.TestHelloWorld(0x0?)
        C:/go/golang-test/helper/hello_world_test.go:9 +0x6e
testing.tRunner(0xc00004dd40, 0x63ba08)
        C:/laragon/bin/go/go1.18.5.windows-amd64/src/testing/testing.go:1439 +0x102
created by testing.(*T).Run
        C:/laragon/bin/go/go1.18.5.windows-amd64/src/testing/testing.go:1486 +0x35f
exit status 2
FAIL    github.com/mahardhika21/golang-test.git/helper  0.348s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test -v
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestHelloWorldMahar
--- PASS: TestHelloWorldMahar (0.00s)
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.318s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test -v -run=TestHelloWorld
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestHelloWorldMahar
--- PASS: TestHelloWorldMahar (0.00s)
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.329s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ go test -v -run=TestHelloWorldMahar
=== RUN   TestHelloWorldMahar
--- PASS: TestHelloWorldMahar (0.00s)
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.260s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test/helper (master)
$ cd ..

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go test ./helper/
ok      github.com/mahardhika21/golang-test.git/helper  0.241s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ go test -v ./...
?       github.com/mahardhika21/golang-test.git [no test files]
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestHelloWorldMahar
--- PASS: TestHelloWorldMahar (0.00s)
PASS
ok      github.com/mahardhika21/golang-test.git/helper  0.271s

mahar@DESKTOP-APF469J MINGW64 /c/go/golang-test (master)
$ TestHelloWorldMahar