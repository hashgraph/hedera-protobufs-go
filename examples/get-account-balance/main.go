package main

import (
	"context"
	"fmt"

	"github.com/hashgraph/hedera-protobufs-go/services"
	"google.golang.org/grpc"
)

func main() {
	// https://docs.hedera.com/guides/testnet/testnet-nodes#testnet-nodes
	conn, err := grpc.Dial("0.testnet.hedera.com:50211", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	cryptoClient := services.NewCryptoServiceClient(conn)

	// https://github.com/hashgraph/hedera-protobufs/blob/main/services/CryptoGetAccountBalance.proto#L35
	getBalanceQuery := services.Query{
		Query: &services.Query_CryptogetAccountBalance{
			CryptogetAccountBalance: &services.CryptoGetAccountBalanceQuery{
				BalanceSource: &services.CryptoGetAccountBalanceQuery_AccountID{
					AccountID: &services.AccountID{
						Account: &services.AccountID_AccountNum{
							AccountNum: 1001,},
					},
				},
			},
		},
	}

	// https://github.com/hashgraph/hedera-protobufs/blob/main/services/CryptoService.proto#L52
	response, err := cryptoClient.CryptoGetBalance(context.TODO(), &getBalanceQuery)
	if err != nil {
		panic(err)
	}

	// https://github.com/hashgraph/hedera-protobufs/blob/main/services/CryptoGetAccountBalance.proto#L47
	getBalanceResponse := response.GetCryptogetAccountBalance()

	fmt.Printf("balance = %v tℏ\n", getBalanceResponse.Balance)
}
