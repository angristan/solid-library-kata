package model

type UserRole int

const (
	Admin  UserRole = iota
	Member          = iota
	Guest           = iota
)

func GetRole(roleName string) UserRole {
	switch roleName {
	case "admin":
		return Admin
	case "member":
		return Member
	default:
		panic("Invalid role")
	}
}

func (role *UserRole) GetString() string {
	switch *role {
	case Admin:
		return "admin"
	case Member:
		return "member"
	default:
		panic("Invalid role")
	}
}
