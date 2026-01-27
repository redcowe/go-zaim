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
		Amount:        1000,
		Date:          time.Now(),
		FromAccountId: 1,
		ToAccountId:   1,
		Comment:       "test",
	}

	res, err := z.Money.CreateTransfer(transferRequest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("transfer created")

	time.Sleep(10 * time.Second)

	updateTransferRequest := &zaim.TransferRequest{
		Amount:        100000,
		Date:          time.Now(),
		FromAccountId: 1,
		ToAccountId:   3,
		Comment:       "test",
	}

	updateRes, err := z.Money.UpdateTransfer(res.PaymentMoney.ID, updateTransferRequest)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prettyJSON, _ := json.MarshalIndent(updateRes, "", "  ")
	fmt.Println(string(prettyJSON))
}
