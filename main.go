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

	res, v := z.VerifyUserAuthentication()
	if res {
		fmt.Println(v.ProfileImageUrl)
	}

}
