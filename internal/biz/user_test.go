package biz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	hashPassword("123456")
}

func TestVerifyPassword(t *testing.T) {
	assert := assert.New(t)
	hash := "$2a$10$8vqJDPnIO4N3iCocsG4e3OdzaWBI0WEwimJDoWckEKsu1FFSzROai"
	pwd := "123456"
	assert.True(verifyPassword(hash, pwd))
}
