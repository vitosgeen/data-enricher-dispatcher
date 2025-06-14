package model

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) IsValid() bool {
	if u.Name == "" || u.Email == "" {
		return false
	}
	return true
}

func (u *User) IsEqual(other *User) bool {
	if u.Name != other.Name || u.Email != other.Email {
		return false
	}
	return true
}

func UserEmailHasSpecialPostfix(user *User, postfix []string) bool {
	if user == nil || user.Email == "" {
		return false
	}
	for _, p := range postfix {
		if len(user.Email) < len(p) {
			continue
		}
		isSpecialPostfix := user.Email[len(user.Email)-len(p):] == p
		if isSpecialPostfix {
			return true
		}
	}
	return false
}
