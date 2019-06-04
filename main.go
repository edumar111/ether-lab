package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	/*privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	*/

	privateKey, err := crypto.HexToECDSA("a29c5861938d1af5d7cc4cd903449b89f33ba0b158aa8bc8d397a9ec6ea74656")
	//A29C5861938D1AF5D7CC4CD903449B89F33BA0B158AA8BC8D397A9EC6EA74656
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E
	//0xf927a06a3c0c1DFe205d43434647F8890700d31E

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	//account, err := ks.NewAccount(password)
	//privateKeyECDSA, ok := privateKey.(*ecdsa.PrivateKey)

	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}
