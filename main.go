package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redcowe/go-zaim/zaim"
)

func main() {
	z, err := zaim.NewZaim()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	transferRequest := &zaim.TransferRequest{
		Amount:        10000,
		Date:          time.Now(),
		FromAccountId: 1,
		ToAccountId:   2,
		Comment:       "test",
	}

	res, err := z.Money.CreateTransfer(transferRequest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prettyJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(prettyJSON))

}
