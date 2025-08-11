package domain

type UserRole string

const (
	RoleAdmin   UserRole = "ADMIN"
	RoleManager UserRole = "MANAGER"
	RoleSales   UserRole = "SALES"
)

type AllowedRoles []UserRole

var (
	RolesAdminOnly       = AllowedRoles{RoleAdmin}
	RolesManagerAndAbove = AllowedRoles{RoleAdmin, RoleManager}
	RolesAll             = AllowedRoles{RoleAdmin, RoleManager, RoleSales}
)
