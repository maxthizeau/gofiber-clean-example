package role

type RoleEnum string

const (
	Administrator RoleEnum = "ADMINISTRATOR"
	Moderator     RoleEnum = "MODERATOR"
)

type RolesList []RoleEnum

// type Roles interface {
// 	Has(roleEnum RoleEnum) bool
// 	NewRolesFromInterface(rolesInterface []interface{}) Roles
// }

type RoleManager struct{}

func (roles RolesList) Has(roleEnum RoleEnum) bool {
	for _, r := range roles {
		if r == roleEnum {
			return true
		}
	}
	return false
}

func (manager *RoleManager) NewRolesFromInterface(rolesInterface []interface{}) RolesList {
	roleList := make(RolesList, 0)
	for _, r := range rolesInterface {
		roleList = append(roleList, r.(RoleEnum))
	}
	return roleList
}

func (manager *RoleManager) ParseRoles(rolesInterface []interface{}) RolesList {
	r := make(RolesList, 0)
	for _, role := range rolesInterface {
		r = append(r, role.(RoleEnum))
	}
	return r
}
