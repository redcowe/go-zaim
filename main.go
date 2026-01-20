package main

import (
	"fmt"
	"os"

	"github.com/redcowe/go-zaim/zaim"
)

func main() {
	z, err := zaim.NewZaim()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := z.VerifyAuthentication()
	fmt.Println(res)
}
