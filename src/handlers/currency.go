package handlers

import (
	"context"
	"go-eth/consts"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func SendNativeCoin(c *gin.Context) {
	// TODO: Implement the logic to send a transaction
	c.JSON(200, gin.H{"status": "transaction_sent"})
}

func ReceiveNativeCoin(c *gin.Context) {
	type RequestBody struct {
		WalletAddress string `json:"walletAddress"`
	}

	var reqBody RequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to connect to the Ethereum node"})
		return
	}

	privateKey, err := crypto.HexToECDSA(consts.WALLET_PRIVATE_KEY)
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid private key"})
		return
	}

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get nonce"})
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get gas price"})
		return
	}

	value := big.NewInt(1000000000000000000) // 1 ETH
	tx := types.NewTransaction(nonce, common.HexToAddress(reqBody.WalletAddress), value, 21000, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get chain ID"})
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign transaction"})
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send transaction"})
		return
	}

	c.JSON(200, gin.H{"status": "transaction_sent", "tx_hash": signedTx.Hash().Hex()})
}
