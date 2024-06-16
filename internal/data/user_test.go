package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"
)

type TestUserProvider struct {
	users  map[int64]*User
	nextID int64
}

func NewTestUserProvider() *TestUserProvider {
	return &TestUserProvider{
		users:  make(map[int64]*User),
		nextID: 1,
	}
}

func (p *TestUserProvider) PutUser(user *User) (*User, error) {
	// Check for DB constraints
	for _, existingUser := range p.users {
		if existingUser.Email == user.Email {
			return nil, fmt.Errorf("user with email %s already exists", user.Email)
		}
		if existingUser.Username == user.Username {
			return nil, fmt.Errorf("user with username %s already exists", user.Username)
		}
	}

	user.Id = p.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	p.users[user.Id] = user
	p.nextID++
	return user, nil
}

func (p *TestUserProvider) GetUser(id int64) (*User, error) {
	user, ok := p.users[id]
	if !ok {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

func (p *TestUserProvider) GetUserByEmail(email string) (*User, error) {
	for _, user := range p.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with email %s not found", email)
}

func (p *TestUserProvider) GetUserByUsername(userName string) (*User, error) {
	for _, user := range p.users {
		if user.Username == userName {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with username %s not found", userName)
}

func (p *TestUserProvider) ListUsers() ([]*User, error) {
	var users []*User
	for _, user := range p.users {
		users = append(users, user)
	}

	return users, nil
}

func (p *TestUserProvider) UpdateUser(user *User) (*User, error) {
	if _, ok := p.users[user.Id]; !ok {
		return nil, fmt.Errorf("user with ID %d not found", user.Id)
	}
	if user.Email != "" {
		p.users[user.Id].Email = user.Email
	}
	if user.Username != "" {
		p.users[user.Id].Username = user.Username
	}

	p.users[user.Id].UpdatedAt = time.Now()
	return p.users[user.Id], nil
}

func (p *TestUserProvider) DeleteUser(id int64) error {
	if _, ok := p.users[id]; !ok {
		return fmt.Errorf("user with ID %d not found", id)
	}
	delete(p.users, id)
	return nil
}

func (p *TestUserProvider) TearDown() {
	p.nextID = 1
	p.users = make(map[int64]*User)
}

func TestUserService_CreateNewUser(t *testing.T) {
	type args struct {
		email    string
		username string
	}

	testCases := []struct {
		name             string
		args             args
		generateUsername bool
		wantErr          bool
		preRun           func(t *testing.T, provider *TestUserProvider)
		expect           *User
	}{
		{
			name: "create new user",
			args: args{
				email:    "test@example.com",
				username: "test",
			},
			wantErr: false,
			expect: &User{
				Id:       1,
				Email:    "test@example.com",
				Username: "test",
			},
		},
		{
			name: "create new user with existing username",
			args: args{
				email:    "different@example.com",
				username: "same",
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "same",
				})
			},
		},
		{
			name: "create new user with existing email",
			args: args{
				email:    "same@example.com",
				username: "different",
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "same@example.com",
					Username: "test",
				})
			},
		},
		{
			name: "create new user with too long username",
			args: args{
				email:    "test@example.com",
				username: strings.Repeat("toolong", 10),
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "create new user with invalid email",
			args: args{
				email:    "invalid",
				username: "test",
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "create new user generated username",
			args: args{
				email:    "test@example.com",
				username: "",
			},
			generateUsername: true,
			wantErr:          false,
			expect: &User{
				Id:       1,
				Email:    "test@example.com",
				Username: "",
			},
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			user, err := service.CreateNewUser(tc.args.email, tc.args.username)
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateNewUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if tc.generateUsername {
				if user.Username == "" {
					t.Errorf("CreateNewUser() = %v, expected generated username", user)
				}
				return
			}
			if !user.Equal(tc.expect) {
				t.Errorf("CreateNewUser() = %v, want %v", user, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestUserService_GetUser(t *testing.T) {
	type args struct {
		id int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
		expect  *User
	}{
		{
			name: "get existing user",
			args: args{
				id: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
			expect: &User{
				Id:       1,
				Email:    "test@example.com",
				Username: "test",
			},
		},
		{
			name: "get non-existing user",
			args: args{
				id: 2,
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			user, err := service.GetUser(tc.args.id)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !user.Equal(tc.expect) {
				t.Errorf("GetUser() = %v, want %v", user, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
		expect  *User
	}{
		{
			name: "get existing user",
			args: args{
				email: "test@example.com",
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
			expect: &User{
				Id:       1,
				Email:    "test@example.com",
				Username: "test",
			},
		},
		{
			name: "get non-existing user",
			args: args{
				email: "different@example.com",
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			user, err := service.GetUserByEmail(tc.args.email)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !user.Equal(tc.expect) {
				t.Errorf("GetUser() = %v, want %v", user, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestUserService_GetUserByUsername(t *testing.T) {
	type args struct {
		username string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
		expect  *User
	}{
		{
			name: "get existing user",
			args: args{
				username: "test",
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
			expect: &User{
				Id:       1,
				Email:    "test@example.com",
				Username: "test",
			},
		},
		{
			name: "get non-existing user",
			args: args{
				username: "different",
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			user, err := service.GetUserByUsername(tc.args.username)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !user.Equal(tc.expect) {
				t.Errorf("GetUser() = %v, want %v", user, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestUserService_ListUsers(t *testing.T) {
	testCases := []struct {
		name    string
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
		expect  []*User
	}{
		{
			name: "list existing users",
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
				provider.PutUser(&User{
					Email:    "test2@example.com",
					Username: "test2",
				})
			},
			expect: []*User{
				{
					Id:       1,
					Email:    "test@example.com",
					Username: "test",
				},
				{
					Id:       2,
					Email:    "test2@example.com",
					Username: "test2",
				},
			},
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			users, err := service.ListUsers()
			if (err != nil) != tc.wantErr {
				t.Errorf("ListUsers() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			for _, user := range users {
				found := false
				for _, expectUser := range tc.expect {
					if user.Equal(expectUser) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("ListUsers() = %v, want %v", user, tc.expect)
				}
			}
		})
		provider.TearDown()
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type args struct {
		user *User
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
		expect  *User
	}{
		{
			name: "update existing user",
			args: args{
				&User{
					Id:       1,
					Email:    "new@example.com",
					Username: "new",
				},
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
			expect: &User{
				Id:       1,
				Email:    "new@example.com",
				Username: "new",
			},
		},
		{
			name: "update non-existing user",
			args: args{
				&User{
					Id:       2,
					Email:    "new@example.com",
					Username: "new",
				},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "update user with 0 id",
			args: args{
				&User{
					Id:       0,
					Email:    "new@example.com",
					Username: "new",
				},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "update user with empty email",
			args: args{
				&User{
					Id:       1,
					Email:    "",
					Username: "new",
				},
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "update user with empty username",
			args: args{
				&User{
					Id:       1,
					Email:    "new@example.com",
					Username: "",
				},
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
		},
		{
			name: "update user with too long username",
			args: args{
				&User{
					Id:       1,
					Email:    "new@example.com",
					Username: strings.Repeat("toolong", 10),
				},
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
		},
		{
			name: "update user with invalid email",
			args: args{
				&User{
					Id:       1,
					Email:    "invalid",
					Username: "new",
				},
			},
			wantErr: true,
			expect:  nil,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			user, err := service.UpdateUser(tc.args.user)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !user.Equal(tc.expect) {
				t.Errorf("UpdateUser() = %v, want %v", user, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	type args struct {
		id int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestUserProvider)
	}{
		{
			name: "delete existing user",
			args: args{
				id: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestUserProvider) {
				provider.PutUser(&User{
					Email:    "test@example.com",
					Username: "test",
				})
			},
		},
		{
			name: "delete non-existing user",
			args: args{
				id: 2,
			},
			wantErr: true,
		},
		{
			name: "delete user with 0 id",
			args: args{
				id: 0,
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		provider := NewTestUserProvider()
		service := NewUserService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			err := service.DeleteUser(tc.args.id)
			if (err != nil) != tc.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
		provider.TearDown()
	}
}

func TestUserService_GenerateUserName(t *testing.T) {
	NameRegex := regexp.MustCompile("^[a-zA-Z0-9]+$")

	t.Run("should generate a random name", func(t *testing.T) {
		service := NewUserService(nil)
		name := service.GenerateUserName()
		if len(name) == 0 {
			t.Errorf("name should not be empty")
		}
		if !validator.Matches(name, NameRegex) {
			t.Errorf("name should match regex %s", NameRegex)
		}
	})
}
