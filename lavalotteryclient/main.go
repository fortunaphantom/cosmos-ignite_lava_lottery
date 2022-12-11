package main

import (
    "context"
    "fmt"
    "log"

    // Importing the general purpose Cosmos blockchain client
    "github.com/ignite/cli/ignite/pkg/cosmosclient"

    // Importing the types package of your lavalottery blockchain
    "lavalottery/x/lavalottery/types"
)

func main() {
	// Prefix to use for account addresses.
    // The address prefix was assigned to the lavalottery blockchain
    // using the `--address-prefix` flag during scaffolding.
    addressPrefix := "cosmos"

    // Create a Cosmos client instance
    cosmos, err := cosmosclient.New(
        context.Background(),
        cosmosclient.WithAddressPrefix(addressPrefix),
    )
    if err != nil {
        log.Fatal(err)
    }

	// Account `alice` was initialized during `ignite chain serve`
    accountName := "client1"

    // Get account from the keyring
    account, err := cosmos.Account(accountName)
    if err != nil {
        log.Fatal(err)
    }

    addr, err := account.Address(addressPrefix)
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println(addr)

    // Define a message to create a post
    msg := &types.MsgSendTicket{
        Creator: addr,
        Fee: "5token",
        Bet: "1token",
    }

    // Broadcast a transaction from account `alice` with the message
    // to create a post store response in txResp
    txResp, err := cosmos.BroadcastTx(context.Background(), account, msg)
    if err != nil {
        log.Fatal(err)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgSendTicket:\n\n")
    fmt.Println(txResp)
}
