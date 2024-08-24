package passwords

import (
	"bytes"
	"crypto/rand"
	"errors"
	"runtime"

	"golang.org/x/crypto/argon2"
)

// The length of salts used
//
// Recommended in the docs: https://github.com/alexedwards/argon2id
const saltLen uint32 = 16

// Generates a Salt to use alongside a password to hash it
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLen)

	_, err := rand.Read(salt)

	return salt, err
}

// A struct that represents a hashed password. Includes the hash and the salt used to hash it
type HashedPassword struct {
	hash []byte
	salt []byte
}

func (h HashedPassword) Hash() []byte {
	return h.hash
}

func (h HashedPassword) Salt() []byte {
	return h.salt
}

// A struct that represents an Argon2id Hash
type Hash struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

// Generates the default hash configuration
func DefaultHash() *Hash {
	return NewHash(4, 128*1024, uint8(runtime.NumCPU()), 32)
}

// Generates a new Hash struct
func NewHash(time, memory uint32, threads uint8, keyLen uint32) *Hash {
	return &Hash{time, memory, threads, keyLen}
}

// Generates a new hash from a password. A salt is automatically generated
func (h Hash) GenerateNewHash(password []byte) (*HashedPassword, error) {
	if len(password) == 0 {
		return nil, errors.New("Password length is 0!")
	}

	salt, err := generateSalt()

	if err != nil {
		return nil, err
	}

	return h.GenerateHash(password, salt)
}

// Generates a hash from a password and a salt
func (h Hash) GenerateHash(password, salt []byte) (*HashedPassword, error) {
	if len(password) == 0 {
		return nil, errors.New("Password length is 0!")
	}

	if len(salt) == 0 {
		return nil, errors.New("Salt length is 0!")
	}

	hash := argon2.IDKey(password, salt, h.time, h.memory, h.threads, h.keyLen)

	return &HashedPassword{hash, salt}, nil
}

func (h Hash) Compare(hash, salt, password []byte) (bool, error) {
	hashedPassword, err := h.GenerateHash(password, salt)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hashedPassword.Hash(), hash), nil
}
