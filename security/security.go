package security

import "golang.org/x/crypto/bcrypt"

// Compara senha com Hash
func CompareHashPassword(passWithHash string, passWithoutHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passWithHash), []byte(passWithoutHash))

}

//Cria Hash para uma senha

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
