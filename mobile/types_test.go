package geth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	from := "0xb008e10e6633befd086028691ecc70648b288689"
	to := "0x0c54FcCd2e384b4BB6f2E405Bf5Cbc15a017AaFb"
	nonce := int64(1)
	gas := int64(21000)
	maxFeePerGas := NewBigInt(10000000000)
	maxPriorityFeePerGas := NewBigInt(10000000000)
	value := NewBigInt(0)
	chainId := NewBigInt(1)
	data := common.FromHex("0xa9059cbb000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000000000000000000000000000000b1a2bc2ec50000")

	fromAddress := Address{
		address: common.HexToAddress(from),
	}
	toAddress := Address{
		address: common.HexToAddress(to),
	}

	tx := NewTransaction(&fromAddress, &toAddress, nonce, gas, maxFeePerGas, maxPriorityFeePerGas, value, chainId, data)

	str, err := tx.EncodeJSON()

	expectedJson := "{\"type\":\"0x2\",\"nonce\":\"0x1\",\"gasPrice\":null,\"maxPriorityFeePerGas\":\"0x2540be400\",\"maxFeePerGas\":\"0x2540be400\",\"gas\":\"0x5208\",\"value\":\"0x0\",\"input\":\"0xa9059cbb000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000000000000000000000000000000b1a2bc2ec50000\",\"v\":\"0x0\",\"r\":\"0x0\",\"s\":\"0x0\",\"to\":\"0x0c54fccd2e384b4bb6f2e405bf5cbc15a017aafb\",\"chainId\":\"0x1\",\"accessList\":[],\"hash\":\"0xaf85ec87dc6bc9c993bfcfa18692c52adea59cc5503a58443b0d054b8951fc70\"}"

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJson, str)
}

func TestNewTransaction_noReceiver(t *testing.T) {
	from := "0xb008e10e6633befd086028691ecc70648b288689"
	var toAddress *Address = nil
	nonce := int64(1)
	gas := int64(21000)
	maxFeePerGas := NewBigInt(10000000000)
	maxPriorityFeePerGas := NewBigInt(10000000000)
	value := NewBigInt(0)
	chainId := NewBigInt(1)
	data := common.FromHex("0xa9059cbb000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000000000000000000000000000000b1a2bc2ec50000")

	fromAddress := Address{
		address: common.HexToAddress(from),
	}

	tx := NewTransaction(&fromAddress, toAddress, nonce, gas, maxFeePerGas, maxPriorityFeePerGas, value, chainId, data)

	str, err := tx.EncodeJSON()

	expectedJson := "{\"type\":\"0x2\",\"nonce\":\"0x1\",\"gasPrice\":null,\"maxPriorityFeePerGas\":\"0x2540be400\",\"maxFeePerGas\":\"0x2540be400\",\"gas\":\"0x5208\",\"value\":\"0x0\",\"input\":\"0xa9059cbb000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000000000000000000000000000000b1a2bc2ec50000\",\"v\":\"0x0\",\"r\":\"0x0\",\"s\":\"0x0\",\"to\":null,\"chainId\":\"0x1\",\"accessList\":[],\"hash\":\"0x47150a651e312339694c605c8b8631be729523e265a4491a539143467afa2528\"}"

	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJson, str)
}
