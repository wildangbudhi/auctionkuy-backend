package auth

type AppleTokenHeader struct {
	KID string `json:"kid"`
	ALG string `json:"alg"`
}
