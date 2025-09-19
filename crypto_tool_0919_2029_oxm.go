// 代码生成时间: 2025-09-19 20:29:28
package main

import (
    "bufio"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "errors"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

// AESGCM is a struct that holds the AES cipher and nonce
type AESGCM struct {
    "block cipher.aes"
    nonce []byte
}

// NewAESGCM creates a new AESGCM instance
func NewAESGCM(key []byte) *AESGCM {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    return &AESGCM{
        aes:  block,
        nonce: make([]byte, aes.BlockSize),
    }
}

// Seal encrypts the data
func (a *AESGCM) Seal(plaintext []byte) ([]byte, error) {
    a.nonce = a.nonce[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, a.nonce); err != nil {
        return nil, err
    }
    return a.seal(a.nonce, plaintext, nil), nil
}

// Open decrypts the data
func (a *AESGCM) Open(ciphertext []byte) ([]byte, error) {
    if len(ciphertext) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    nonceSize := aes.BlockSize
    return a.open(ciphertext[:nonceSize], ciphertext[nonceSize:]), nil
}

// seal encrypts the plaintext using the provided nonce
func (a *AESGCM) seal(nonce, plaintext, additionalData []byte) []byte {
    // Note: Since no additional data is used, we pass nil
    return a.aes.Seal(nonce, nonce, plaintext, nil)
}

// open decrypts the ciphertext using the provided nonce
func (a *AESGCM) open(nonce, ciphertext []byte) ([]byte, error) {
    if len(ciphertext) < aes.BlockSize {
        return nil, errors.New("ciphertext too short")
    }
    return a.aes.Open(nonce, ciphertext[:aes.BlockSize], ciphertext[aes.BlockSize:], nil)
}

func main() {
    key := []byte("your_secret_key_here") // Replace with your secret key
    aesgcm := NewAESGCM(key)

    // Read input from user
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text to encrypt/decrypt: ")
    input, _ := reader.ReadString('
')
    input = input[:len(input)-1] // Remove newline

    var encrypted, decrypted []byte
    var err error

    // Check if the input is already hex-encoded
    if _, err = hex.DecodeString(input); err == nil {
        // Assume it's encrypted and try to decrypt
        decrypted, err = aesgcm.Open([]byte(input))
        if err != nil {
            fmt.Printf("Decryption failed: %s
", err)
            return
        }
        fmt.Printf("Decrypted text: %s
", string(decrypted))
    } else {
        // Assume it's plain text and try to encrypt
        encrypted, err = aesgcm.Seal([]byte(input))
        if err != nil {
            fmt.Printf("Encryption failed: %s
", err)
            return
        }
        fmt.Printf("Encrypted text: %s
", hex.EncodeToString(encrypted))
    }
}
