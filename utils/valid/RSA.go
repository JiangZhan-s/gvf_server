package valid

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// GenerateRSAKeyPair RSA密钥对生成
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// RsaEncrypt RSA公钥加密
func RsaEncrypt(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// RsaDecrypt RSA私钥解密
func RsaDecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// EncodePrivateKeyToPEM RSA私钥PEM格式编码
func EncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) ([]byte, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	if privateKeyPEM == nil {
		return nil, errors.New("Failed to encode private key to PEM format")
	}
	return privateKeyPEM, nil
}

// EncodePublicKeyToPEM RSA公钥PEM格式编码
func EncodePublicKeyToPEM(publicKey *rsa.PublicKey) ([]byte, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	if publicKeyPEM == nil {
		return nil, errors.New("Failed to encode public key to PEM format")
	}
	return publicKeyPEM, nil
}

// DecodePrivateKeyFromPEM RSA私钥PEM格式解码
func DecodePrivateKeyFromPEM(privateKeyPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return nil, errors.New("Failed to decode private key from PEM format")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// DecodePublicKeyFromPEM RSA公钥PEM格式解码
func DecodePublicKeyFromPEM(publicKeyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return nil, errors.New("Failed to decode public key from PEM format")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Failed to parse RSA public key")
	}
	return publicKey, nil
}

func EncryptWithRSA(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	encryptedData := make([]byte, 0)

	// 计算每个块的大小
	blockSize := publicKey.Size() - 11

	// 对数据进行分块加密
	for i := 0; i < len(data); i += blockSize {
		endIndex := i + blockSize
		if endIndex > len(data) {
			endIndex = len(data)
		}

		// 获取当前块的数据
		block := data[i:endIndex]

		// 使用 RSA 公钥加密当前块
		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, block)
		if err != nil {
			return nil, fmt.Errorf("encryption error: %v", err)
		}

		// 将加密后的块追加到结果中
		encryptedData = append(encryptedData, encryptedBlock...)
	}

	return encryptedData, nil
}

func DecryptWithRSA(privateKey *rsa.PrivateKey, encryptedData []byte) ([]byte, error) {
	decryptedData := make([]byte, 0)

	// 计算每个块的大小
	blockSize := privateKey.Size()

	// 对数据进行分块解密
	for i := 0; i < len(encryptedData); i += blockSize {
		endIndex := i + blockSize
		if endIndex > len(encryptedData) {
			endIndex = len(encryptedData)
		}

		// 获取当前块的数据
		block := encryptedData[i:endIndex]

		// 使用 RSA 私钥解密当前块
		decryptedBlock, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, block)
		if err != nil {
			return nil, fmt.Errorf("decryption error: %v", err)
		}

		// 将解密后的块追加到结果中
		decryptedData = append(decryptedData, decryptedBlock...)
	}

	return decryptedData, nil
}
