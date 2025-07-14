package sshServer

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/vars"

	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

func generateAndSaveKeys(privateKeyFile string, publicKeyFile string) {

	// 生成密钥对
	privateKey, err := GenerateRSAKey()
	if err != nil {
		panic(fmt.Errorf("生成RSA密钥对失败:%v", err))
	}
	publicKey := &privateKey.PublicKey

	// 保存密钥对文件
	if err := SavePrivateKey(privateKeyFile, privateKey); err != nil {
		panic(fmt.Errorf("保存私钥失败:%v", err))
	}
	if err := SavePublicKey(publicKeyFile, publicKey); err != nil {
		panic(fmt.Errorf("保存公钥失败:%v", err))
	}

}

func getKeys(privateKeyFile string, publicKeyFile string) ([]byte, []byte) {

	// 读取私钥
	privateKeyBytes, err := LoadKeyBytes(privateKeyFile)
	if err != nil {
		panic(fmt.Errorf("读取私钥失败:%v", err))
	}

	// 读取公钥
	publicKeyBytes, err := LoadKeyBytes(publicKeyFile)
	if err != nil {
		panic(fmt.Errorf("读取公钥失败:%v", err))
	}

	// 校验是否是同一对
	if !VerifyKeyPair(privateKeyBytes, publicKeyBytes) {
		panic(fmt.Errorf("私钥和公钥不匹配, 请删除 key 文件夹并重新运行"))
	}

	return privateKeyBytes, publicKeyBytes

}

func Start(OnPasswordAuth func(ctx ssh.Context, password string) bool, OnConnect func(session ssh.Session)) {

	// 读取配置文件
	fmt.Println("读取配置文件...")
	config := GetConfig()

	// 获取密钥对文件
	privateKeyFile := vars.FILE_SSH_SERVER_PRIVATE_KEY
	publicKeyFile := vars.FILE_SSH_SERVER_PUBLIC_KEY

	// 不存在则新建
	if !helper.IsExist(privateKeyFile) || !helper.IsExist(publicKeyFile) {
		fmt.Println("密钥对不存在, 正在生成...")
		generateAndSaveKeys(
			privateKeyFile, publicKeyFile,
		)
	}

	// 读取密钥对文件
	fmt.Println("读取密钥对文件...")
	privateKeyBytes, _ := getKeys(
		privateKeyFile, publicKeyFile,
	)

	// 启动SSH服务器
	fmt.Println("启动SSH服务器在端口:", config.Port)
	server := &ssh.Server{
		Addr:            fmt.Sprintf(":%d", config.Port),
		Handler:         OnConnect,
		PasswordHandler: OnPasswordAuth,
	}

	// 加载私钥
	signer, err := gossh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		panic(fmt.Errorf("服务器不能启动, 因为解析私钥失败: %v", err))
	}
	server.AddHostKey(signer)

	// 启动服务器
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("服务器不能启动: %v", err))
	}

}
