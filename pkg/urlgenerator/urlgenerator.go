package urlgenerator

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

const shortURLLength = 10 // max 43

func GenerateShortURL(originalURL string) string {
	hasher := sha256.New()
	hasher.Write([]byte(originalURL))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return strings.TrimRight(hash, "=")[:shortURLLength]
}
