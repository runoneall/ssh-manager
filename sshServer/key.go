package sshServer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"ssh-manager/helper"
	"ssh-manager/vars"

	"github.com/google/uuid"
)

func GenerateRSAKey() (*rsa.PrivateKey, error) {

	// 生成RSA私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, vars.VAR_RSA_KEY_LENGTH)
	if err != nil {
		return nil, err
	}

	// 验证私钥是否有效
	if err := privateKey.Validate(); err != nil {
		return nil, err
	}

	return privateKey, nil
}

func SavePrivateKey(path string, privateKey *rsa.PrivateKey) error {

	// 已存在则覆盖
	if !helper.DeleteFileIfExist(path) {
		return fmt.Errorf("不能删除文件: %s", path)
	}

	// 创建私钥文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 作为PEM编码写入文件
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(file, privateKeyBlock); err != nil {
		return err
	}

	return nil
}

func SavePublicKey(path string, publicKey *rsa.PublicKey) error {

	// 已存在则覆盖
	if !helper.DeleteFileIfExist(path) {
		return fmt.Errorf("不能删除文件: %s", path)
	}

	// 创建公钥文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 作为PEM编码写入文件
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(file, publicKeyBlock); err != nil {
		return err
	}

	return nil

}

func ParsePKCS1PrivateKeyFromBytes(keyBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("无效的PEM格式")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("不是RSA私钥PEM块")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ParsePKIXPublicKeyFromBytes(keyBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("无效的PEM格式")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("不是PKIX公钥PEM块")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析PKIX公钥失败: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("不是RSA公钥")
	}

	return rsaPub, nil
}

func VerifyKeyPair(privateKeyBytes []byte, publicKeyBytes []byte) bool {

	// 解析密钥对
	privateKey, err := ParsePKCS1PrivateKeyFromBytes(privateKeyBytes)
	if err != nil {
		return false
	}

	publicKey, err := ParsePKIXPublicKeyFromBytes(publicKeyBytes)
	if err != nil {
		return false
	}

	// 生成一个uuid
	testUUID := uuid.New().String()
	testMsg := []byte(testUUID)

	// 使用公钥加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, testMsg)
	if err != nil {
		fmt.Println("无法加密:", testUUID)
		return false
	}

	// 使用私钥解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		fmt.Println("无法解密:", ciphertext)
		return false
	}

	// 解密结果
	return string(plaintext) == testUUID

}

func LoadKeyBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}
