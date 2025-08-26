package user

type Role string

const (
	RoleBank     Role = "bank"
	RoleCustomer Role = "customer"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
