package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/caarlos0/env/v6"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	IV         string `env:"CIPHER_IV"`
	KeyVersion string `env:"CIPHER_KEY_VERSION"`
	HashKey    string `env:"HASH_KEY,required"`
	DEK        string `env:"DEK"`
	TEST       string `env:"TEST"`
}

var (
	km      *keyManager
	hashKey []byte
)

type keyManager struct {
	deks          map[string][]byte
	version       string
	key           []byte
	staticVersion string
	staticIV      []byte
}

func (k keyManager) getKey(version string) []byte {
	return k.deks[version]
}

func SetUp(wdeksJson string) error {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return err
	}

	staticIV, err := base64.StdEncoding.DecodeString(cfg.IV)
	if err != nil {
		return fmt.Errorf("static iv invalid %s", err)
	}

	wdeks := map[string]string{}
	err = json.Unmarshal([]byte(wdeksJson), &wdeks)
	if err != nil {
		return fmt.Errorf("wdeks err: %s", err)
	}

	deks := map[string][]byte{}
	for key, _ := range wdeks {
		dek := []byte(cfg.DEK)
		deks[key] = dek
	}
	// set up key manager
	versions := make([]string, 0, len(deks))
	hasStaticVersion := false
	for version := range deks {
		if version == cfg.KeyVersion {
			hasStaticVersion = true
		}
		versions = append(versions, version)
	}
	if !hasStaticVersion {
		return fmt.Errorf("static key version not found")
	}
	sort.Strings(versions)
	latestVersion := versions[len(versions)-1]

	km = &keyManager{
		deks:          deks,
		version:       latestVersion,
		key:           deks[latestVersion],
		staticVersion: cfg.KeyVersion,
		staticIV:      staticIV,
	}
	hashKey = []byte(cfg.HashKey)
	return nil
}

// Encrypt an AES CBC ciphertext deprecated
func Encrypt(p string) (string, error) {
	plaintext := []byte(p)
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	if len(plaintext)%aes.BlockSize != 0 {
		return "", ErrPlaintextIsNotMultipleOfBlockSize
	}
	block, err := aes.NewCipher(km.key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%s:%s:%s", km.version, base64.StdEncoding.EncodeToString(iv), base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:])), nil
}

// Decrypt an AES CBC ciphertext deprecated
func Decrypt(encrypted string) (decryptedString string, err error) {
	s := strings.Split(encrypted, ":")
	if len(s) != 3 {
		return "", ErrInvalidEncryptedValueFormat
	}
	version := s[0]
	iv, err := base64.StdEncoding.DecodeString(s[1])
	ciphertext, err := base64.StdEncoding.DecodeString(s[2])
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(km.getKey(version))
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", ErrInvalidEncryptedValueBlockSize
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", ErrInvalidEncryptedValueBlockSize
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	length := len(ciphertext)
	unppading := int(ciphertext[length-1])
	ciphertext = ciphertext[:(length - unppading)]
	return string(ciphertext), nil
}

func EncryptAes256StaticIV(plaintext string) (string, error) {
	ciphertext, err := EncryptGCM([]byte(plaintext), km.getKey(km.staticVersion), km.staticIV)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", km.staticVersion, base64.StdEncoding.EncodeToString(ciphertext)), nil
}

func EncryptAes256(plaintext string) (string, error) {
	//hVmYq3t6w9z$B&E)H@McQfTjWnZr4u7x
	ciphertext, err := EncryptGCM([]byte(plaintext), km.key, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", km.version, base64.StdEncoding.EncodeToString(ciphertext)), nil
}

func DecryptAes256StaticIV(encrypted string) (string, error) {
	s := strings.Split(encrypted, ":")
	if len(s) != 2 {
		return Decrypt(encrypted)
	}
	version := s[0]
	ciphertext, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	plainText, err := DecryptGCM(ciphertext, km.getKey(version), km.staticIV)
	return plainText, err
}

func DecryptAes256(encrypted string) (string, error) {
	s := strings.Split(encrypted, ":")
	if len(s) != 2 {
		return Decrypt(encrypted)
	}
	version := s[0]
	ciphertext, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return "", err
	}
	plainText, err := DecryptGCM(ciphertext, km.getKey(version), nil)
	return plainText, err
}

// EncryptGCM an AES GCM ciphertext
func EncryptGCM(value []byte, key []byte, nonce []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}
	var dst []byte
	if nonce == nil {
		nonce = make([]byte, gcm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			return []byte{}, err
		}
		dst = nonce
	}

	// Using nonce as Seal's dst argument results in it being the first
	// chunk of bytes in the ciphertext. Decrypt retrieves the nonce/IV from this.
	ciphertext := gcm.Seal(dst, nonce, value, nil)
	return ciphertext, nil
}

// DecryptGCM an AES GCM ciphertext
func DecryptGCM(ciphertext []byte, key []byte, nonce []byte) (string, error) {
	c, err := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	if nonce == nil {
		nonceSize := gcm.NonceSize()
		nonce, ciphertext = ciphertext[:nonceSize], ciphertext[nonceSize:]
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func Hash(data string) string {
	hash := hmac.New(sha256.New, hashKey)
	hash.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

/*
password := "secret"

	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var (
	ErrInvalidEncryptedValueFormat       = fmt.Errorf("invalid encrypted value format")
	ErrInvalidEncryptedValueBlockSize    = fmt.Errorf("invalid encrypted value BlockSize")
	ErrPlaintextIsNotMultipleOfBlockSize = fmt.Errorf("plaintext is not a multiple of the block size")
	ErrCipherMessageAuthenticationFailed = fmt.Errorf("cipher: message authentication failed")
)
