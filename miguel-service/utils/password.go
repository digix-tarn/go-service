package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword รับ password ธรรมดา แล้ว return รหัสผ่านที่ถูก hash หรือ error
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash เปรียบเทียบ password ธรรมดากับ hashed password
// ถ้ารหัสผ่านตรงกันจะ return nil (ไม่มี error)
func CheckPasswordHash(password, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}