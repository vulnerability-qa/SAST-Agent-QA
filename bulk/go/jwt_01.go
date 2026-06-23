// CWE-347: JWT verification skipped
package main
import (
	"encoding/base64"
	"encoding/json"
	"strings"
)
func parseJWT(token string) map[string]interface{} {
	parts := strings.Split(token, ".")
	payload, _ := base64.RawURLEncoding.DecodeString(parts[1])
	var claims map[string]interface{}
	json.Unmarshal(payload, &claims)
	return claims // signature never verified
}
