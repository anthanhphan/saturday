package jwt

type Payload struct {
	UserId int64 `json:"user_id" validate:"required"`
}
