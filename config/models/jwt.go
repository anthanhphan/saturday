package models

type JwtConfig struct {
	PrivateKeyPath     string `yaml:"private_key_path" json:"private_key_path"`
	PublicKeyPath      string `yaml:"public_key_path" json:"public_key_path"`
	AccessTokenExpiry  int64  `yaml:"access_token_expiry" json:"access_token_expiry"`
	RefreshTokenExpiry int64  `yaml:"refresh_token_expiry" json:"refresh_token_expiry"`
}
