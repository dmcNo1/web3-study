package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
)

func main() {
	// getMsg()
	// getAccount()
	// createPrivateKey()
	// transfer()
	transferToken()
	// subscribe()
}

func getMsg() {
	// 开启一个连接
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	// 获取指定的区块头，如果不指定区块高度，传参为nil的话，会返回最新的区块头
	blockNumber := big.NewInt(22195096)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.Uint64())     // 22195096
	fmt.Println(header.Time)                // 1743764159
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xb48fe2994ea399c13722e14ec26bbe4d3e201e926cc1bbedb2bf53c150ceab8e

	// 获取区块的交易数
	count, err := client.TransactionCount(context.Background(), header.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)

	// 获取chainId
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chain id =", chainId)

	// 获取区块
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 遍历区块中的交易信息
	var txHash common.Hash
	for _, tx := range block.Transactions() {
		fmt.Println("hash =", tx.Hash().Hex())
		fmt.Println("交易数量 =", tx.Value().String())         // 100000000000000000
		fmt.Println("gas =", tx.Gas())                     // 21000
		fmt.Println("gas price =", tx.GasPrice().Uint64()) // 100000000000
		fmt.Println("nonce =", tx.Nonce())                 // 245132
		fmt.Println(tx.Data())                             // []
		fmt.Println("to hash =", tx.To().Hex())            // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 获取到交易sender的信息
		if sender, err := types.Sender(types.NewEIP155Signer(chainId), tx); err == nil {
			fmt.Println("sender =", sender.Hex())
		}

		// 每个交易都有一个收据，其中包含执行交易的结果，例如所有的返回值和日志，以及“1”（成功）或“0”（失败）的交易结果状态。
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receipt status =", receipt.Status) // 1
		fmt.Println("receipt logs =", receipt.Logs)     // []

		txHash = tx.Hash()

		break // 只查看第一条
	}

	// 也可以这样遍历交易
	blockHash := common.HexToHash("0xb48fe2994ea399c13722e14ec26bbe4d3e201e926cc1bbedb2bf53c150ceab8e")
	for i := uint(0); i < count; i++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tx =", tx.Hash().Hex())
		break
	}

	// pendingFlag:交易是否在等待中
	tx, pendingFlag, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx =", tx)
	fmt.Println("pendingFlag =", pendingFlag)

	// 获取区块中的交易收据
	receipts, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	// 也可以这样获取
	// receiptsByNumber, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for _, receipt := range receipts {
		fmt.Println(receipt.Status)
		fmt.Println(receipt.Logs)
		fmt.Println(receipt.TxHash.Hex())
		fmt.Println(receipt.TransactionIndex)
		fmt.Println("contract address =", receipt.TransactionIndex)
		break
	}
}

func getToken() {

}

func getAccount() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	// 获取账户的最新余额
	account := common.HexToAddress("0x572573b8abE8328e1891c4DF79ECD887e6f42A15")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	// 获取指定区块上该账户的余额
	// blockNumber := big.NewInt(22195096)
	// balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balanceAt)

	// Wei -> eth
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	// 查看待处理的余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pendingBalance)
}

// 创建钱包（也就是生成私钥）
func createPrivateKey() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// 也可以这样，用已知的私钥生成*ecdsa.PrivateKey
	// privateKey, err := crypto.HexToECDSA("f8d90b48facaffe74069609e73b07e00f05c1c844ac6a474d3ee94e1732e")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(privateKeyBytes)
	// 去掉开头的0x
	fmt.Println(hexutil.Encode(privateKeyBytes[2:]))

	// 用私钥生成对应的公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 公钥以0x04打头，前四个字节无效
	fmt.Println("public key =", hexutil.Encode(publicKeyBytes[4:]))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

func transfer() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	// 获取转出方信息
	privateKey, err := crypto.HexToECDSA("8a34079f38c2135d988dd18700a77e77bca8383d0ad3780e805b64496443cf89")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 转账0.001
	// value := big.NewInt(0)
	// gasLimit := uint64(0)
	// gasPrice := big.NewInt(0)
	value := big.NewInt(1000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 转入方信息
	toAddress := common.HexToAddress("0xfb1465D90eA24429487Bdcb8e8B5310c2Fa31140")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}

func transferToken() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	// 获取转出方信息
	privateKey, err := crypto.HexToECDSA("8a34079f38c2135d988dd18700a77e77bca8383d0ad3780e805b64496443cf89")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	value := big.NewInt(0)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xfb1465D90eA24429487Bdcb8e8B5310c2Fa31140")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	// 交易的函数
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))
	// 左填充地址到32位
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10)
	// 左填充token数量到32位
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	// 交易的data
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 获取gas相关
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

// 订阅区块
func subscribe() {
	// 订阅区块的话，必须通过websocket
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个channel，用于接收订阅消息
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 通过channel接收消息
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xec3da9167a5f48f2f988b35d5bba292324cf1ab5259273549328588654e930cb
			fmt.Println(block.Number().Uint64())   // 8055043
			fmt.Println(block.Time())              // 1743842880
			fmt.Println(block.Nonce())             // 0
			fmt.Println(len(block.Transactions())) // 81
		}
	}
}
