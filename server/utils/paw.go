package utils

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/paw-digital/nano/address"
	"github.com/paw-digital/nano/types"
)

const pawRegexStr = "(?:paw)(?:_)(?:1|3)(?:[13456789abcdefghijkmnopqrstuwxyz]{59})"

var pawRegex = regexp.MustCompile(pawRegexStr)

func GenerateAddress() string {
	pub, _ := address.GenerateKey()
	return strings.Replace(string(address.PubKeyToAddress(pub)), "nano_", "paw_", -1)
}

// ValidateAddress - Returns true if a nano address is valid
func ValidateAddress(account string) bool {
	if !pawRegex.MatchString(account) {
		return false
	}
	return address.ValidateAddress(types.Account(account))
}

// Convert address to pubkey
func AddressToPub(account string) string {
	pubkey, _ := address.AddressToPub(types.Account(account))
	return hex.EncodeToString(pubkey)
}
