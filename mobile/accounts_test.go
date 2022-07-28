package geth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestSignTx(t *testing.T) {
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

	tmpDir := t.TempDir()
	tmpDirPath, _ := filepath.Abs(tmpDir)
	ks := NewKeyStore(tmpDirPath, 2, 1)

	pass := "" // not used but required by API
	a1, err1 := ks.NewAccount(pass)
	if err1 != nil {
		t.Fatal(err1)
	}
	if err := ks.Unlock(a1, ""); err != nil {
		t.Fatal(err)
	}

	signed, err2 := ks.SignTx(a1, tx, NewBigInt(1))
	if err2 != nil {
		t.Fatal(err2)
	}

	assert.NotEqual(t, nil, signed)
}
