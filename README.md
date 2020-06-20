# How To Start A GO Project


First make a directory for your project in `~/projects`

```
cd ~/projects
mkdir dupe
cd dupe
```

Next init a go project with `go mod <namespace>`

```
go mod init github.com/parsnips/dupe
```

Then init a git repository:

```
git init
```

And add the files as a first commit:

```
git add .
git commit -m "first commit"
```

Now you can start adding files. In this case let's add a hello world program:

main.go

```golang
package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
}
```
And now you can build it and run:

```
go build ./...
./dupe
```

or use the `go run` command

```
go run main.go 
```