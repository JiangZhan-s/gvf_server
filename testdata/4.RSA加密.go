package main

import (
	"fmt"
	"gvf_server/utils/valid"
	"log"
	"os"
)

func main() {
	// 生成 RSA 密钥对
	privateKey, publicKey, err := valid.GenerateRSAKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	// 将私钥保存到文件
	privateKeyPEM, err := valid.EncodePrivateKeyToPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("private.pem", privateKeyPEM, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("私钥已保存到 private.pem 文件")

	// 将公钥保存到文件
	publicKeyPEM, err := valid.EncodePublicKeyToPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("public.pem", publicKeyPEM, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("公钥已保存到 public.pem 文件")
}
