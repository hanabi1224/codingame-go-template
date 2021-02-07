# codingame-go-template
Codingame template for golang to allow multi-file local repo and non-std libraries

## How to use
1. Install bundle tool
  ```go get -u golang.org/x/tools/...```

1. Customize go generate directive
    ```go
    // Open main.go and change below line
    // go:generate bundle -o ./../out/bundle.go -prefix "" .
    ```

 1. Use codingame chrome extension to connect generated single file to browser IDE


## How to bundle other libraries

    Take go-linq as example


```
go install github.com/ahmetalpbalkan/go-linq
cd $GOPATH/github.com/ahmetalpbalkan/go-linq
bundle -o ./../pkg_go_linq.go -pkg main -prefix  .
```
