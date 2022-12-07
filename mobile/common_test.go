// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the common package.

package geth

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestHashTypedData(t *testing.T) {
	var inputWithIntChainId = "{\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Person\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"Mail\":[{\"name\":\"from\",\"type\":\"Person\"},{\"name\":\"to\",\"type\":\"Person\"},{\"name\":\"contents\",\"type\":\"string\"}]},\"primaryType\":\"Mail\",\"domain\":{\"name\":\"Ether Mail\",\"version\":\"1\",\"chainId\":1,\"verifyingContract\":\"0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC\"},\"message\":{\"from\":{\"name\":\"Cow\",\"wallet\":\"0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826\"},\"to\":{\"name\":\"Bob\",\"wallet\":\"0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB\"},\"contents\":\"Hello, Bob!\"}}"
	expectedSignature := "0xbe609aee343fb3c4b28e1df9e632fca64fcfaede20f02e86244efddf30957bd2"
	var inputWithStringChainId = "{\"types\":{\"EIP712Domain\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"Person\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"wallet\",\"type\":\"address\"}],\"Mail\":[{\"name\":\"from\",\"type\":\"Person\"},{\"name\":\"to\",\"type\":\"Person\"},{\"name\":\"contents\",\"type\":\"string\"}]},\"primaryType\":\"Mail\",\"domain\":{\"name\":\"Ether Mail\",\"version\":\"1\",\"chainId\":\"1\",\"verifyingContract\":\"0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC\"},\"message\":{\"from\":{\"name\":\"Cow\",\"wallet\":\"0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826\"},\"to\":{\"name\":\"Bob\",\"wallet\":\"0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB\"},\"contents\":\"Hello, Bob!\"}}"

	signature1, err1 := HashTypedData(inputWithIntChainId)
	assert.Equal(t, expectedSignature, EncodeToHex(signature1))
	assert.Equal(t, nil, err1)

	signature2, err2 := HashTypedData(inputWithStringChainId)
	assert.Equal(t, expectedSignature, EncodeToHex(signature2))
	assert.Equal(t, nil, err2)
}

func TestHashTypedData2(t *testing.T) {
	var inputWithIntChainId = "{\"types\":{\"EIP712Domain\":[{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"version\"},{\"type\":\"uint256\",\"name\":\"chainId\"},{\"type\":\"address\",\"name\":\"verifyingContract\"}],\"Part\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint96\"}],\"Mint721\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"},{\"name\":\"creators\",\"type\":\"Part[]\"},{\"name\":\"royalties\",\"type\":\"Part[]\"}]},\"domain\":{\"name\":\"Mint721\",\"version\":\"1\",\"chainId\":4,\"verifyingContract\":\"0x6ede7f3c26975aad32a475e1021d8f6f39c89d82\"},\"primaryType\":\"Mint721\",\"message\":{\"royalties\":[{\"account\":\"0xa3d87ab2b2f86f00453876249148d8c7db47342e\",\"value\":1000}],\"creators\":[{\"account\":\"0xa3d87ab2b2f86f00453876249148d8c7db47342e\",\"value\":10000}],\"tokenId\":\"74109480115837155093015070630862379235416487407035131988895773941165604208644\",\"tokenURI\":\"/ipfs/bafkreidbh5q5vzosqqki72ejugqiqcvhymefjvua4fvzuiop4ogbyekuam\"}}"
	expectedSignature := "0x9d95d5f675b3562cc0187bab13bc20b73b53132e4515fc7ba443c4fe1f467854"

	signature, err := HashTypedData(inputWithIntChainId)
	assert.Equal(t, expectedSignature, EncodeToHex(signature))
	assert.Equal(t, nil, err)
}

func TestJsonConversion(t *testing.T) {
	accessListReceipt, err := NewReceiptFromJSON("{\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"blockNumber\":\"0xa044e4\",\"contractAddress\":null,\"cumulativeGasUsed\":\"0xadf850\",\"effectiveGasPrice\":\"0x3b9af54c\",\"from\":\"0xb008e10e6633befd086028691ecc70648b288689\",\"gasUsed\":\"0x31f7c\",\"logs\":[{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0xcae9d16f553e92058883de29cb3135dbc0c1e31fd7eace79fef1d80577fe482e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b008e10e6633befd086028691ecc70648b2886890000000000000000000000000320de3378dcde180758ad2d41c0e1c6dcbb441db45a3ba100000000000000000000000000000000000000000000000000000000b3c5c69700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000005af3107a4000aaaebeba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x71\",\"removed\":false},{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0xcae9d16f553e92058883de29cb3135dbc0c1e31fd7eace79fef1d80577fe482e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000076c5855e93bd498b6331652854c4549d34bc3a30b45a3ba100000000000000000000000000000000000000000000000000000000dfdfdaf4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000016bcc41e9000aaaebeba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x72\",\"removed\":false},{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0xcae9d16f553e92058883de29cb3135dbc0c1e31fd7eace79fef1d80577fe482e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000076c5855e93bd498b6331652854c4549d34bc3a30b45a3ba100000000000000000000000000000000000000000000000000000000dfdfdaf4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000016bcc41e9000aaaebeba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x73\",\"removed\":false},{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0xcae9d16f553e92058883de29cb3135dbc0c1e31fd7eace79fef1d80577fe482e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000b008e10e6633befd086028691ecc70648b2886890000000000000000000000000320de3378dcde180758ad2d41c0e1c6dcbb441db45a3ba100000000000000000000000000000000000000000000000000000000a10bb5b200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000031bced02db000aaaebeba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x74\",\"removed\":false},{\"address\":\"0x21a932c8e5eac252be0a0860b18c4edb8ee66034\",\"topics\":[\"0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62\",\"0x0000000000000000000000007d47126a2600e22eab9ed6cf0e515678727779a6\",\"0x0000000000000000000000000320de3378dcde180758ad2d41c0e1c6dcbb441d\",\"0x000000000000000000000000b008e10e6633befd086028691ecc70648b288689\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x75\",\"removed\":false},{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0xcae9d16f553e92058883de29cb3135dbc0c1e31fd7eace79fef1d80577fe482e\"],\"data\":\"0x00000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000320de3378dcde180758ad2d41c0e1c6dcbb441d000000000000000000000000b008e10e6633befd086028691ecc70648b2886891a0388dd00000000000000000000000000000000000000000000000000000000a10bb5b20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000001973bb640000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000004000000000000000000000000021a932c8e5eac252be0a0860b18c4edb8ee660340000000000000000000000000000000000000000000000000000000000000001\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x76\",\"removed\":false},{\"address\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"topics\":[\"0x268820db288a211986b26a8fda86b1e0046281b21206936bb0e61c67b5c79ef4\"],\"data\":\"0x1581fa663524d578fdd5e8fd0eca7df2929fc3f35768498b3770a94c184bf63b9e10f6018e8c3edaeb7dd00fab291257f03e3bdb1c2d18402e684810ce6231210000000000000000000000000320de3378dcde180758ad2d41c0e1c6dcbb441d000000000000000000000000b008e10e6633befd086028691ecc70648b28868900000000000000000000000000000000000000000000000000038d7ea4c680000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001a0973bb640000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000004000000000000000000000000021a932c8e5eac252be0a0860b18c4edb8ee660340000000000000000000000000000000000000000000000000000000000000001aaaebeba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000\",\"blockNumber\":\"0xa044e4\",\"transactionHash\":\"0xad37193d2ec1d29e39e65a3a8c8a840d16e9c195622613ea2ddf8230efc78e3b\",\"transactionIndex\":\"0x40\",\"blockHash\":\"0xe2601fc63e9b0ac3a7ca618bd826730a47c03966532478becacd00d190534f3d\",\"logIndex\":\"0x77\",\"removed\":false}],\"logsBloom\":\"0x00000000000000000000000000020000000000000000000000000004000000000200000000000000000000000000080000000000000000000000000000002000000000000000000000000000200000000000000000800000000000000000000000000000000000800000000000000000000000000000000000040000004000008000000000008000008000000000000000008000000000002000000000000000000000000040000000000000001000008000000000000002080400000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000001000000000000000000000000000000080000000000\",\"status\":\"0x1\",\"to\":\"0xd4a57a3bd3657d0d46b4c5bac12b3f156b9b886b\",\"transactionIndex\":\"0x40\",\"type\":\"0x0\"}")
	assert.Equal(t, nil, err)
	println(accessListReceipt.EncodeJSON())
}

func TestPersonalEcRecoverSignedByMetamask(t *testing.T) {
	testAddr := common.HexToAddress("0x232BfEBD605a8b16509261a5EfE26e58E9537829")

	msg := []byte("Example `personal_sign` message")
	sig, _ := DecodeFromHex("0x606f9ef6e664241893293a3bbb766df1e13cea010181a90398e31a71805d8a675d765f8d085a228e2e7450a574cb8d4fb1273308bacc66a2cfd8592ae0698c461b")

	addressBytes, recoverError := PersonalEcRecover(msg, sig)
	assert.Equal(t, nil, recoverError)

	recoveredAddr := common.HexToAddress(common.Bytes2Hex(addressBytes))
	assert.Equal(t, testAddr, recoveredAddr)
}

func TestPersonalEcRecoverSignedByMetamaskTestSite(t *testing.T) {
	testAddr := common.HexToAddress("0x63b4512c705638bbba1ebc41af6b2fc3da1d8c03")

	msg := []byte("Example `personal_sign` message")
	sig, _ := DecodeFromHex("0x8a1a0dd717418a27a0763154d8a9db587b228408e5bd5a9445ee2ba38fd48c68692bd7af9b955ae0b4af5f19f992a0d699f95490a18df4ef6b9d344ca635ade91c")

	addressBytes, recoverError := PersonalEcRecover(msg, sig)
	assert.Equal(t, nil, recoverError)

	recoveredAddr := common.HexToAddress(common.Bytes2Hex(addressBytes))
	assert.Equal(t, testAddr, recoveredAddr)
}

func TestPersonalEcRecoverSignedByGo(t *testing.T) {
	var testAddrHex = "970e8128ab834e8eac17ab8e3812f010678cf791"
	var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

	key, _ := crypto.HexToECDSA(testPrivHex)
	testAddr := common.HexToAddress(testAddrHex)

	msg := []byte("foo")
	msgToSign, _ := accounts.TextAndHash(msg)
	sig, signError := crypto.Sign(msgToSign, key)
	sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	assert.Equal(t, nil, signError)

	addressBytes, recoverError := PersonalEcRecover(msg, sig)
	assert.Equal(t, nil, recoverError)

	recoveredAddr := common.HexToAddress(common.Bytes2Hex(addressBytes))
	assert.Equal(t, testAddr, recoveredAddr)
}
