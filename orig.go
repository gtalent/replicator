package main

import (
	"fmt"
	"strings"
)

var pgm = `package main

import (
	"fmt"
	"strings"
)

var pgm = narf

func main() {
	out := strings.Replace(pgm, "narf", string(96) + pgm + string(96), 1)
	fmt.Println(out)
}`

func main() {
	out := strings.Replace(pgm, "narf", string(96) + pgm + string(96), 1)
	fmt.Println(out)
}
