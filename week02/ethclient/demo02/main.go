package main

import (
	"context"
	"crypto/ecdsa"
	"demo02/contracts/store"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// 合约的字节码
	contractByteCode = "608060405234801561000f575f5ffd5b5060405161085038038061085083398181016040528101906100319190610193565b805f908161003f91906103ea565b50506104b9565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f5f905090565b61032f610320565b61033a8184846102fb565b505050565b5b8181101561035d576103525f82610327565b600181019050610340565b5050565b601f8211156103a25761037381610241565b61037c84610253565b8101602085101561038b578190505b61039f61039785610253565b83018261033f565b50505b505050565b5f82821c905092915050565b5f6103c25f19846008026103a7565b1980831691505092915050565b5f6103da83836103b3565b9150826002028217905092915050565b6103f3826101da565b67ffffffffffffffff81111561040c5761040b61006f565b5b6104168254610211565b610421828285610361565b5f60209050601f831160018114610452575f8415610440578287015190505b61044a85826103cf565b8655506104b1565b601f19841661046086610241565b5f5b8281101561048757848901518255600182019150602085019450602081019050610462565b868310156104a457848901516104a0601f8916826103b3565b8355505b6001600288020188555050505b505050505050565b61038a806104c65f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806348f343f31461004357806354fd4d5014610073578063f56256c714610091575b5f5ffd5b61005d600480360381019061005891906101d6565b6100ad565b60405161006a9190610210565b60405180910390f35b61007b6100c2565b6040516100889190610299565b60405180910390f35b6100ab60048036038101906100a691906102b9565b61014d565b005b6001602052805f5260405f205f915090505481565b5f80546100ce90610324565b80601f01602080910402602001604051908101604052809291908181526020018280546100fa90610324565b80156101455780601f1061011c57610100808354040283529160200191610145565b820191905f5260205f20905b81548152906001019060200180831161012857829003601f168201915b505050505081565b8060015f8481526020019081526020015f2081905550817fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4826040516101939190610210565b60405180910390a25050565b5f5ffd5b5f819050919050565b6101b5816101a3565b81146101bf575f5ffd5b50565b5f813590506101d0816101ac565b92915050565b5f602082840312156101eb576101ea61019f565b5b5f6101f8848285016101c2565b91505092915050565b61020a816101a3565b82525050565b5f6020820190506102235f830184610201565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61026b82610229565b6102758185610233565b9350610285818560208601610243565b61028e81610251565b840191505092915050565b5f6020820190508181035f8301526102b18184610261565b905092915050565b5f5f604083850312156102cf576102ce61019f565b5b5f6102dc858286016101c2565b92505060206102ed858286016101c2565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061033b57607f821691505b60208210810361034e5761034d6102f7565b5b5091905056fea26469706673582212206e83b0751b19633d7e1e6be01f7bad3b157f871d16a7a909432653a6dfbd10c264736f6c634300081d0033"
	// 部署的合约地址
	contractAddrss = "0xC55A3204C436623F042b36846B9177921b784E38"

	testAccount1PrivateKey = "8a34079f38c2135d988dd18700a77e77bca8383d0ad3780e805b64496443cf89"

	clientAddr = "https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA"

	contractAbiStr = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func main() {
	// deployByGo()
	// deployByBin()
	// loadContracte()
	invokeContract()
	// queryLogs()
	// subscribeEvents()
}

// 通过abigen，将sol编译成go文件然后部署
func deployByGo() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyECDSA, err := crypto.HexToECDSA("8a34079f38c2135d988dd18700a77e77bca8383d0ad3780e805b64496443cf89")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)

	version := "1.0"
	storeAddress, tx, storeInstance, err := store.DeployStore(auth, client, version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storeAddress.Hex()) // 0x6C2E1709553d3Fe0Fd4499821817C88e573f8c0E
	fmt.Println(tx.Hash().Hex())    // 0x2037cfec5530cb1549bb2d359b46728b8891e13134b5441ad59f98c9c0007f02
	fmt.Println(storeInstance)      // &{{0xc0005e2008} {0xc0005e2008} {0xc0005e2008}}
}

// 通过二进制方式部署
func deployByBin() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyECDSA, err := crypto.HexToECDSA("8a34079f38c2135d988dd18700a77e77bca8383d0ad3780e805b64496443cf89")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gaPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 解码合约的字节码
	contract, err := hex.DecodeString(contractByteCode)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 创建交易
	gaPrice = big.NewInt(1).Mul(gaPrice, big.NewInt(20))
	tx := types.NewContractCreation(nonce, big.NewInt(0), uint64(3000000), gaPrice, contract)
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex()) // 0x0b43f91cc46cf5be2dd47e1143f413d38a4dc8e7ed2b89dbe118d53483bf7d25

	// 阻塞等待矿工挖矿
	receipt, err := waitForReicept(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed at: %s\n", receipt.ContractAddress.Hex()) // 0xD21ddC7c64068F58912c770C4C306d50205D9561
}

func waitForReicept(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		fmt.Println(err)

		time.Sleep(time.Second)
	}
}

// 加载合约
func loadContracte() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	storeAddress := common.HexToAddress(contractAddrss)
	instance, err := store.NewStore(storeAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instance)
}

func invokeContract() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	storeAddress := common.HexToAddress(contractAddrss)
	instance, err := store.NewStore(storeAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(testAccount1PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value"))

	// 11155111就是sepolia的chainId
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetItem(opts, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())

	// 获取合约返回的数据
	callOpts := bind.CallOpts{Context: context.Background()}
	valueInContract, err := instance.Items(&callOpts, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valueInContract)
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}

func invokeContractAbi() {
	client, err := ethclient.Dial(clientAddr)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(testAccount1PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 获取Abi文件信息
	contractAbi, err := abi.JSON(strings.NewReader(contractAbiStr))
	if err != nil {
		log.Fatal(err)
	}
	// 调用的合约方法
	methodName := "ItemSet"
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_key_abi"))
	copy(value[:], []byte("demo_value_abi"))
	// 封装调用的方法
	data, err := contractAbi.Pack(methodName, key, value)
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddrss), big.NewInt(0), uint64(300000), gasPrice, data)

	chainId := big.NewInt(11155111)
	// 签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReicept(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 查询刚刚设置的值
	callInput, err := contractAbi.Pack("items", key)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(contractAddrss)
	callMsg := ethereum.CallMsg{
		To:   &toAddress,
		Data: callInput,
	}

	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	err = contractAbi.UnpackIntoInterface(&unpacked, "items", result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unpacked)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func queryLogs() {
	client, err := ethclient.Dial(clientAddr)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xC55A3204C436623F042b36846B9177921b784E38")
	// 生成查询信息
	query := ethereum.FilterQuery{
		// FromBlock: big.NewInt(1),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
		// 封装data
		data := struct {
			key   [32]byte
			value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&data, "ItemSet", vLog.Data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(common.Bytes2Hex(data.key[:]))
		fmt.Println(common.Bytes2Hex(data.value[:]))

		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}
		fmt.Println("topics[0] = ", topics[0])
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex())
}

func subscribeEvents() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/Ng0L0W_L8-FPX4BWHR5FDvgzyAaRnubA")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xC55A3204C436623F042b36846B9177921b784E38")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
