// scrypt.go
package main
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    // "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "golang.org/x/crypto/scrypt"
)
func Encrypt(key, data []byte) ([]byte, error) {
    key, salt, err := DeriveKey(key, nil)
    if err != nil {
        return nil, err
    }
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err = rand.Read(nonce); err != nil {
        return nil, err
    }
    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    ciphertext = append(ciphertext, salt...)
    return ciphertext, nil
}
func Decrypt(key, data []byte) ([]byte, error) {
    salt, data := data[len(data)-32:], data[:len(data)-32]
    key, _, err := DeriveKey(key, salt)
    if err != nil {
        return nil, err
    }
    blockCipher, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    gcm, err := cipher.NewGCM(blockCipher)
    if err != nil {
        return nil, err
    }
    nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, err
    }
    return plaintext, nil
}
func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
    if salt == nil {
        salt = make([]byte, 32)
        if _, err := rand.Read(salt); err != nil {
            return nil, nil, err
        }
    }
    key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
    if err != nil {
        return nil, nil, err
    }
    return key, salt, nil
}
func main() {
    var (
        password = []byte("mysecretpassword")
        data     = []byte("our super secret text")
    )
    ciphertext, err := Encrypt(password, data)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ciphertext: %s\n", hex.EncodeToString(ciphertext))
    ciphertext = []byte("7a953b876e1dcb0207530fd03a699abedd0989f081d0f7925bca2bff5c9886fe095d2a97a4d02a0a6b638cdd0934f19e56b2bad9b247becb5543f295c02ebf1acbc8ea45909cf2c322d24925a966d3a20b")
    plaintext, err := Decrypt(password, ciphertext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("plaintext: %s\n", plaintext)
}