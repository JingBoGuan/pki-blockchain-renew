package main

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func populate() {
	// 创建一个缓冲区来存储请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件
	file, err := ioutil.ReadFile("D:\\goProject\\src\\pki-blockchain-master\\catest\\subca.crt")
	if err != nil {
		panic(err)
	}

	// 解码 PEM 编码的证书
	block, _ := pem.Decode(file)
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the certificate")
		return
	}

	fileWriter, err := writer.CreateFormFile("UplFiles", "cacert.crt")
	if err != nil {
		panic(err)
	}

	// 将文件内容复制到请求体中
	_, err = io.Copy(fileWriter, bytes.NewReader(block.Bytes))
	if err != nil {
		panic(err)
	}

	// 添加文本字段
	writer.WriteField("ContrAddr", "0xe58A4f402A2c3136445Cfd7F2cCf31Eeb0E2b5c6")
	writer.WriteField("ParentAddr", "0x1CbE49F033a416149212bccC9681F353d5d5F410")
	writer.WriteField("NewUserAddr", "0x36d231aD83971B2fF1918db43AD2AE04101fae5E")
	writer.WriteField("CurrentUserAddr", "0xD24196422b6163D7DdB49F740DDD18945496CF33")

	//结束多部分写入
	writer.Close()

	// 创建 POST 请求
	request, err := http.NewRequest("POST", "http://localhost:8081/populate_contract", &requestBody)
	if err != nil {
		panic(err)
	}

	// 设置正确的 Content-Type
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 处理响应
	if response.StatusCode == http.StatusOK {
		// 请求成功
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response:", string(body))
	} else {
		// 请求失败，处理错误信息
		fmt.Println("Error:", response.Status)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Error Response:", string(body))
	}
}

func create() {
	// 创建一个缓冲区来存储请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文本字段
	writer.WriteField("CurrentUserAddr", "0xD24196422b6163D7DdB49F740DDD18945496CF33")
	writer.WriteField("NewUserAddr", "0x36d231aD83971B2fF1918db43AD2AE04101fae5E")
	writer.WriteField("ParentAddr", "0x1CbE49F033a416149212bccC9681F353d5d5F410")

	//结束多部分写入
	writer.Close()

	// 创建 POST 请求
	request, err := http.NewRequest("POST", "http://localhost:8081/create_contract", &requestBody)
	if err != nil {
		panic(err)
	}

	// 设置正确的 Content-Type
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 处理响应
	if response.StatusCode == http.StatusOK {
		// 请求成功
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response:", string(body))
	} else {
		// 请求失败，处理错误信息
		fmt.Println("Error:", response.Status)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Error Response:", string(body))
	}
}

func enroll_user() {
	// 创建一个缓冲区来存储请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件
	file, err := ioutil.ReadFile("D:\\goProject\\src\\pki-blockchain-master\\catest\\user3.crt")
	if err != nil {
		panic(err)
	}

	// 解码 PEM 编码的证书
	block, _ := pem.Decode(file)
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the certificate")
		return
	}

	fileWriter, err := writer.CreateFormFile("UplFiles", "cacert.crt")
	if err != nil {
		panic(err)
	}

	// 将文件内容复制到请求体中
	_, err = io.Copy(fileWriter, bytes.NewReader(block.Bytes))
	if err != nil {
		panic(err)
	}

	// 添加文本字段
	writer.WriteField("ParentAddr", "0x1CbE49F033a416149212bccC9681F353d5d5F410")
	writer.WriteField("CurrentUserAddr", "0xD24196422b6163D7DdB49F740DDD18945496CF33")
	writer.WriteField("NewUserAddr", "0xD9Ca9162582B77333E82007c57AfD85a7460d05C")

	//结束多部分写入
	writer.Close()

	// 创建 POST 请求
	request, err := http.NewRequest("POST", "http://localhost:8081/enroll_user", &requestBody)
	if err != nil {
		panic(err)
	}

	// 设置正确的 Content-Type
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 处理响应
	if response.StatusCode == http.StatusOK {
		// 请求成功
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Response:", string(body))
	} else {
		// 请求失败，处理错误信息
		fmt.Println("Error:", response.Status)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Error Response:", string(body))
	}
}

func download() {
	// 创建一个缓冲区来存储请求体
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文本字段
	writer.WriteField("ContrAddr", "0x1CbE49F033a416149212bccC9681F353d5d5F410")

	//结束多部分写入
	writer.Close()

	// 创建 POST 请求
	request, err := http.NewRequest("POST", "http://localhost:8081/download_cacert", &requestBody)
	if err != nil {
		panic(err)
	}

	// 设置正确的 Content-Type
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 处理响应
	if response.StatusCode == http.StatusOK {
		// 请求成功
		if err != nil {
			panic(err)
		}
		file, err := os.Create("ca.crt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		// 将响应内容写入本地文件
		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("CA certifate successfully!")
	} else {
		// 请求失败，处理错误信息
		fmt.Println("Error:", response.Status)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Error Response:", string(body))
	}
}

func hashtest() {
	// 添加文件
	file, err := ioutil.ReadFile("D:\\goProject\\src\\pki-blockchain-master\\catest\\user3.crt")
	if err != nil {
		panic(err)
	}

	// 解码 PEM 编码的证书
	block, _ := pem.Decode(file)
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the certificate")
		return
	}

	hasher := sha3.NewLegacyKeccak256()
	dst4hash := block.Bytes
	if err != nil {
		fmt.Sprintf("Error in open file for hash: ", err.Error())
	}

	hasher.Write(dst4hash)
	hashSum := hasher.Sum(nil)
	fmt.Printf("dst4hash: %x\n", dst4hash)
	fmt.Printf("Hash: %x\n", hashSum)
}

func cal(address string) {
	hash := common.HexToHash(address)
	for _, value := range hash.Bytes() {
		hex := fmt.Sprintf("%x:", value)
		fmt.Print(hex)
	}
}

func main() {
	hashtest()
}
