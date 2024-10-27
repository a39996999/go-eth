package handlers

import (
	"go-eth/consts"
	"go-eth/service"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func ReceiveNativeCoin(c *gin.Context) {
	type RequestBody struct {
		WalletAddress string `json:"walletAddress"`
	}

	var reqBody RequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil || !common.IsHexAddress(reqBody.WalletAddress) {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	privateKey, err := crypto.HexToECDSA(consts.WALLET_PRIVATE_KEY)
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid private key"})
		return
	}

	nonce, err := service.GetPendingNonceAt(address.Hex())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get nonce"})
		return
	}

	gasPrice, err := service.SuggestGasPrice()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get gas price"})
		return
	}

	value := big.NewInt(1000000000000000000) // 1 ETH
	tx := types.NewTransaction(nonce, common.HexToAddress(reqBody.WalletAddress), value, 21000, gasPrice, nil)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(consts.CHAIN_ID)), privateKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign transaction"})
		return
	}

	err = service.SendTransaction(signedTx)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send transaction"})
		return
	}

	c.JSON(200, gin.H{"status": "success", "transactionHash": signedTx.Hash().Hex()})
}
