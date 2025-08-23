package tokening

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionToken(val string) (string, error) {
	token := val

	// Convert the UUID to a byte slice
	id := uuid.New()
	idBytes := id[:]

	// Hash the UUID using bcrypt
	hashedKey, err := bcrypt.GenerateFromPassword(idBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	exp := "exp:" + fmt.Sprint(time.Now().Add(time.Hour*24).Unix())
	body := base64.StdEncoding.EncodeToString([]byte(val))
	header := base64.StdEncoding.EncodeToString([]byte(exp))
	signature := base64.StdEncoding.EncodeToString([]byte(hashedKey))
	token = "Y3VzdG9t" + "." + body + "." + header + "." + signature
	return token, nil
}
