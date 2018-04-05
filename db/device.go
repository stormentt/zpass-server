package db

import (
	"encoding/hex"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stormentt/zpass-lib/zcrypto"
)

type Device struct {
	ID       int
	Selector string `gorm:"unique_index;size:20"`

	AuthKey string `gorm:"size:64"`

	User   User
	UserID int

	LastNonce     string `gorm:"size:32"`
	LastNonceTime *time.Time
}

func FindDevice(sel string) (*Device, error) {
	device := Device{}
	err := Con.Where("selector = ?", sel).Find(&device).Error
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (d *Device) AuthPair() *zcrypto.AuthPair {
	keyBytes := make([]byte, 32)
	n, err := hex.Decode(keyBytes, []byte(d.AuthKey))
	if err != nil || n != 32 {
		log.WithFields(log.Fields{
			"key":       keyBytes,
			"decodeLen": n,
		}).Error("malformed device auth key")
		return nil
	}

	pair, _ := zcrypto.AuthPairFromBytes(keyBytes)
	return pair
}

func (d *Device) CheckSig(msg, sig []byte) bool {
	pair := d.AuthPair()
	if pair == nil {
		return false
	}

	return pair.Verify(msg, sig)
}
