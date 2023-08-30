package helpers

import "golang.org/x/crypto/bcrypt"

func HassPass(pass string) (string, error) {
	p := []byte(pass)
	salt := 8
	Hass, err := bcrypt.GenerateFromPassword(p, salt)
	if err != nil {
		return "", err
	}
	return string(Hass), nil
}

func ComparePass(hass, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hass, pass)

	return err == nil
}
