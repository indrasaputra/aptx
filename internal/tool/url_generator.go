package tool

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/indrasaputra/aptx/entity"
)

const (
	letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// ShortURLGenerator is responsible to generate a unique short URL.
// It uses a simple random string. Being said, uniqueness is not 100% guaranteed.
type ShortURLGenerator struct {
	domain        string
	lettersLength int
	randMaxInt    *big.Int
}

// NewShortURLGenerator creates an instance of ShortURLGenerator.
func NewShortURLGenerator(domain string) *ShortURLGenerator {
	domain = strings.TrimRightFunc(domain, func(r rune) bool {
		return r == '/'
	})

	return &ShortURLGenerator{
		domain:        domain,
		lettersLength: len(letters),
		randMaxInt:    big.NewInt(int64(len(letters))),
	}
}

// Generate generates a code and a short URL with the given length.
func (sg *ShortURLGenerator) Generate(length uint) (string, string, error) {
	code, err := sg.generateRandomString(length)
	if err != nil {
		return "", "", entity.ErrInternal(err.Error())
	}
	return code, fmt.Sprintf("%s/%s", sg.domain, code), nil
}

func (sg *ShortURLGenerator) generateRandomString(length uint) (string, error) {
	b := make([]byte, length)
	for i := 0; i < int(length); i++ {
		num, err := rand.Int(rand.Reader, sg.randMaxInt)
		if err != nil {
			return "", err
		}
		b[i] = letters[num.Int64()]
	}
	return string(b), nil
}
