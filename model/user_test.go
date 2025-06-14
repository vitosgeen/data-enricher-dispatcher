package model

import "testing"

func TestUserEmailHasSpecialPostfix(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		postfix  []string
		expected bool
	}{
		{
			name:     "nil user",
			user:     nil,
			postfix:  []string{"@gmail.com"},
			expected: false,
		},
		{
			name:     "empty email",
			user:     &User{Name: "Alice", Email: ""},
			postfix:  []string{"@gmail.com"},
			expected: false,
		},
		{
			name:     "no postfix match",
			user:     &User{Name: "Bob", Email: "bob@example.com"},
			postfix:  []string{"@gmail.com", "@yahoo.com"},
			expected: false,
		},
		{
			name:     "single postfix match",
			user:     &User{Name: "Carol", Email: "carol@gmail.com"},
			postfix:  []string{"@gmail.com"},
			expected: true,
		},
		{
			name:     "multiple postfixes, one matches",
			user:     &User{Name: "Dave", Email: "dave@yahoo.com"},
			postfix:  []string{"@gmail.com", "@yahoo.com"},
			expected: true,
		},
		{
			name:     "multiple postfixes, none match",
			user:     &User{Name: "Eve", Email: "eve@outlook.com"},
			postfix:  []string{"@gmail.com", "@yahoo.com"},
			expected: false,
		},
		{
			name:     "empty postfix list",
			user:     &User{Name: "Frank", Email: "frank@gmail.com"},
			postfix:  []string{},
			expected: false,
		},
		{
			name:     "postfix longer than email",
			user:     &User{Name: "Grace", Email: "g@a.co"},
			postfix:  []string{"@gmail.com"},
			expected: false,
		},
		{
			name:     "postfix is empty string",
			user:     &User{Name: "Heidi", Email: "heidi@gmail.com"},
			postfix:  []string{""},
			expected: true,
		},
		{
			name:     "postfix is .biz",
			user:     &User{Name: "Ivan", Email: "Ivan@gmail.biz"},
			postfix:  []string{".biz"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UserEmailHasSpecialPostfix(tt.user, tt.postfix)
			if result != tt.expected {
				t.Errorf("UserEmailHasSpecialPostfix() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUser_IsEqual(t *testing.T) {
	tests := []struct {
		name     string
		user1    *User
		user2    *User
		expected bool
	}{
		{
			name:     "both users equal",
			user1:    &User{Name: "Alice", Email: "alice@example.com"},
			user2:    &User{Name: "Alice", Email: "alice@example.com"},
			expected: true,
		},
		{
			name:     "different names",
			user1:    &User{Name: "Alice", Email: "alice@example.com"},
			user2:    &User{Name: "Bob", Email: "alice@example.com"},
			expected: false,
		},
		{
			name:     "different emails",
			user1:    &User{Name: "Alice", Email: "alice@example.com"},
			user2:    &User{Name: "Alice", Email: "alice@gmail.com"},
			expected: false,
		},
		{
			name:     "both fields different",
			user1:    &User{Name: "Alice", Email: "alice@example.com"},
			user2:    &User{Name: "Bob", Email: "bob@gmail.com"},
			expected: false,
		},
		{
			name:     "empty users",
			user1:    &User{},
			user2:    &User{},
			expected: true,
		},
		{
			name:     "one empty, one filled",
			user1:    &User{},
			user2:    &User{Name: "Alice", Email: "alice@example.com"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.user1.IsEqual(tt.user2)
			if result != tt.expected {
				t.Errorf("IsEqual() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestUser_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		expected bool
	}{
		{
			name:     "valid user",
			user:     &User{Name: "Alice", Email: "alice@example.com"},
			expected: true,
		},
		{
			name:     "empty name",
			user:     &User{Name: "", Email: "alice@example.com"},
			expected: false,
		},
		{
			name:     "empty email",
			user:     &User{Name: "Alice", Email: ""},
			expected: false,
		},
		{
			name:     "both empty",
			user:     &User{Name: "", Email: ""},
			expected: false,
		},
		{
			name:     "nil user",
			user:     nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool
			if tt.user == nil {
				result = false
			} else {
				result = tt.user.IsValid()
			}
			if result != tt.expected {
				t.Errorf("IsValid() = %v, want %v", result, tt.expected)
			}
		})
	}
}
