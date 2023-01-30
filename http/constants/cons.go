package constants

type AuthnPayload struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type AuthnRole string

const (
	User  AuthnRole = "user"
	Admin AuthnRole = "admin"
)
