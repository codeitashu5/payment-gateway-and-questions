package ccAvenueEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

import (
	"bytes"
	"crypto/md5"
)

func Pad(data []byte) []byte {
	length := 16 - (len(data) % 16)
	padding := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, padding...)
}
func Encrypt(arrayOfString []string, workingKey string) string {

	encEncryptionString := ""
	for i, val := range arrayOfString {
		if i > 0 {
			encEncryptionString = encEncryptionString + "&" + val
			continue
		}

		encEncryptionString = encEncryptionString + val

	}

	plainText := encEncryptionString

	iv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	plainTextByte := []byte(plainText)
	workingKeyByte := []byte(workingKey)
	plainTextByte = Pad(plainTextByte)
	encDigest := md5.New()
	encDigest.Write(workingKeyByte)
	encCipher, _ := aes.NewCipher(encDigest.Sum(nil))
	encMode := cipher.NewCBCEncrypter(encCipher, iv)
	encryptedText := make([]byte, len(plainTextByte))
	encMode.CryptBlocks(encryptedText, plainTextByte)
	return hex.EncodeToString(encryptedText)
}
func Decrypt(cipherText string, workingKey string) []byte {
	workingKeyByte := []byte(workingKey)
	iv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	decDigest := md5.New()
	decDigest.Write(workingKeyByte)
	encryptedText, _ := hex.DecodeString(cipherText)
	decCipher, _ := aes.NewCipher(decDigest.Sum(nil))
	decMode := cipher.NewCBCDecrypter(decCipher, iv)
	decryptedText := make([]byte, len(encryptedText))
	decMode.CryptBlocks(decryptedText, encryptedText)
	return Unpad(decryptedText)
}
func Unpad(data []byte) []byte {
	length := int(data[len(data)-1])
	return data[:len(data)-length]
}

func DecryptEncVal(encryptedText string, workingKey string) (string, error) {
	iv := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	encryptedTextBytes, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	decDigest := md5.Sum([]byte(workingKey))
	decCipher, err := aes.NewCipher(decDigest[:])
	if err != nil {
		return "", err
	}

	decMode := cipher.NewCBCDecrypter(decCipher, iv)
	decryptedTextBytes := make([]byte, len(encryptedTextBytes))
	decMode.CryptBlocks(decryptedTextBytes, encryptedTextBytes)

	return strings.TrimRight(string(decryptedTextBytes), "\x00"), nil
}

func DoRequest() {
	// Set your encrypted data and access code
	encryptedData := "YOUR_ENCRYPTED_DATA"
	accessCode := "AVUM05KH10AQ49MUQA"

	// URL for the CCAvenue initiation endpoint
	reqUrl := "https://test.ccavenue.com/transaction/transaction.do?command=initiateTransaction"

	// Create the POST request body
	formData := url.Values{}
	formData.Set("encRequest", encryptedData)
	formData.Set("access_code", accessCode)
	reqBody := formData.Encode()

	// Create the HTTP request
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBufferString(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(respBody))
}
