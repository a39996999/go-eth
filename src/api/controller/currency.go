package controller

import (
	"context"
	"go-eth/bootstrap"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	EthClient *ethclient.Client
	Env       *bootstrap.Env
}

func (cc *CurrencyController) ReceiveNativeCoin(c *gin.Context) {
	type RequestBody struct {
		WalletAddress string `json:"walletAddress"`
	}

	var reqBody RequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil || !common.IsHexAddress(reqBody.WalletAddress) {
		log.Println("Invalid request body:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	privateKey, err := crypto.HexToECDSA(cc.Env.PrivateKey)
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	if err != nil {
		log.Println("Invalid private key:", err)
		c.JSON(400, gin.H{"error": "Invalid private key"})
		return
	}

	nonce, err := cc.EthClient.PendingNonceAt(context.Background(), common.HexToAddress(address.Hex()))
	if err != nil {
		log.Println("Failed to get nonce:", err)
		c.JSON(500, gin.H{"error": "Failed to get nonce"})
		return
	}

	gasPrice, err := cc.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("Failed to get gas price:", err)
		c.JSON(500, gin.H{"error": "Failed to get gas price"})
		return
	}

	value := big.NewInt(1000000000000000000) // 1 ETH
	tx := types.NewTransaction(nonce, common.HexToAddress(reqBody.WalletAddress), value, 21000, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(cc.Env.ChainID))), privateKey)
	if err != nil {
		log.Println("Failed to sign transaction:", err)
		c.JSON(500, gin.H{"error": "Failed to sign transaction"})
		return
	}

	err = cc.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("Failed to send transaction:", err)
		c.JSON(500, gin.H{"error": "Failed to send transaction"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "transactionHash": signedTx.Hash().Hex()})
}
