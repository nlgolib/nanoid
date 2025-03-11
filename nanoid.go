package nanoid

import (
	"strings"

	go_nanoid "github.com/matoous/go-nanoid"
)

const (
	ALPHABET_NUMBERS   = "0123456789"
	ALPHABET_LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	ALPHABET_UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// New generates a new nanoid with the given size and alphabet.
func New(size int, alphabet string) (string, error) {
	nanoid, err := go_nanoid.Generate(alphabet, size)
	if err != nil {
		return "", err
	}
	return nanoid, nil
}

// NewWithoutAlike generates a new nanoid without alike characters (1lI0Oouv5Ss).
// It removes the alike characters from the alphabet and generates a new nanoid.
func NewWithoutAlike(size int, alphabet string) (string, error) {
	alikes := "1lI0Oouv5Ss"
	for _, char := range alikes {
		alphabet = strings.ReplaceAll(alphabet, string(char), "")
	}
	return New(size, alphabet)
}

// NewSafe generates a new nanoid with all alphabet characters without alike characters.
// The length is 21, which is proofed to be safe from collisions.
func NewSafe() string {
	output, _ := NewWithoutAlike(21, ALPHABET_NUMBERS+ALPHABET_LOWERCASE+ALPHABET_UPPERCASE)
	return output
}
