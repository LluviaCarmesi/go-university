package utilities;

import "golang.org/x/crypto/bcrypt";

func EncryptPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func DoesPasswordMatch(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}