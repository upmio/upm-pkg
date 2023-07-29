package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//对明文进行填充
func Padding(plainText []byte, blockSize int) []byte {
	//计算要填充的长度
	n := blockSize - len(plainText)%blockSize
	//对原来的明文填充n个n
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

//对密文删除填充
func UnPadding(cipherText []byte) []byte {
	//取出密文最后一个字节end
	end := cipherText[len(cipherText)-1]
	//删除填充
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

//AEC加密（CBC模式）
func AES_CBC_Encrypt(plainText []byte, key []byte) (string, error) {
	//指定加密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//进行填充
	plainText = Padding(plainText, block.BlockSize())
	//指定初始向量vi,长度和block的块尺寸一致
	iv := []byte("f8/NeLsJ*s*vygV@")
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密连续数据库
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	//返回密文字符串
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

//AEC解密（CBC模式）
func AES_CBC_Decrypt(cipherText string, key []byte) ([]byte, error) {
	encrypt_str, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//指定初始化向量IV,和加密的一致
	iv := []byte("f8/NeLsJ*s*vygV@")
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(encrypt_str))
	blockMode.CryptBlocks(plainText, encrypt_str)
	//删除填充
	plainText = UnPadding(plainText)
	return plainText, nil
}
