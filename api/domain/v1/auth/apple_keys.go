package auth

type ApplePubKey struct {
	KTY string `json:"kty"`
	KID string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type AppleKeys struct {
	Keys []ApplePubKey `json:"keys"`
}

type AppleKeysRepository interface {
	GetAppleKeys() (*AppleKeys, error)
}
