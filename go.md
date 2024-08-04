To create a Go program that can be executed directly without using `go run`, you'll need to:

1. **Use `go build`**: Compile your code into a binary executable file.
2. **Add a main function**: Define a single entry point in your Go code, typically named `main()`.
3. **Specify the GOOS and GOARCH environment variables**: Tell Go which operating system and architecture you're
targeting.

Here's an example:

**`main.go`**
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
```
To build this code into a binary executable file:
```bash
$ GOOS=linux GOARCH=amd64 go build main.go
```
This will create an executable binary file called `main` in the current directory. You can then execute it directly
like any other command-line program:
```bash
$ ./main
Hello, World!
```
Note:

* The `GOOS` and `GOARCH` environment variables are not required if you're building for a specific platform (e.g.,
Linux or macOS) on that same platform.
* If you're cross-compiling for another platform, specify the `GOOS` and `GOARCH` environment variables accordingly.

After executing the program directly, Go will automatically run the necessary initialization code to set up the
runtime.
