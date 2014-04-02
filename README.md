# ubuntu

```bash
mkdir -p src/github.com/javouhey
ln -s ../../.. src/github.com/javouhey/epictetus

$ export GOPATH=${PWD}

$ go build

$ go install github.com/javouhey/epictetus
```

# windows

Create the directory src/github.com/javouhey
change directory to `javouhey`

```bash
$ mklink /D epictetus ..\..\..
```

Same as above 
