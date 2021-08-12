package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pkg/expect"
	"github.com/pkg/math"
)

func main() {
	fmt.Println(math.Max(1, 2))
	fmt.Println("hello aashish!")
	expect.True(1 == 1)
	fmt.Println(errors.New(("testing new errors")))
}
