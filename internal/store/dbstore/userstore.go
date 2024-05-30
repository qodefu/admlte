// Package dbstore is designated for data storage operations, specifically for user data in this context.
package dbstore

// Importing necessary packages:
// "errors" for handling error conditions,
// "goth/internal/store" presumably contains shared types or interfaces, including the User struct definition.
import (
	"errors" // Standard package for creating error objects.
	"fmt"
	"goth/internal/store" // Internal package where the User struct is defined.
)

// UserStore struct holds a slice of users, acting as an in-memory storage mechanism.
// This is a simple approach and not suitable for production environments due to its non-persistent nature.
type UserStore struct {
	users []store.User // Slice of User structs to store user data.
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewUserStore() *UserStore {
	return &UserStore{
		users: []store.User{
			// Initializing with a default user.
			{
				Name:     "one",
				Email:    "1@example.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Rajah Owen",
				Email:    "qiwatohud@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Breanna Sellers",
				Email:    "bsellers@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Aspen Armstrong",
				Email:    "aatrong@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Ezra Boyd",
				Email:    "eboyd@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Keely Calhoun",
				Email:    "kcalhoun@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Kay Ryan",
				Email:    "kryan@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Elijah Ayala",
				Email:    "eayala@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Timothy Barnes",
				Email:    "tbarnes@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Ciaran Leach",
				Email:    "cleach@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Oscar Goodman",
				Email:    "ogodman@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Erich Delacruz",
				Email:    "edelcaruz@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
			{
				Name:     "Rigel Barlow",
				Email:    "rbarlow@mailinator.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
		},
	}
}

func (s *UserStore) UpdateUser(name, email string, password string) error {
	for i, user := range s.users {
		if user.Email == email {
			fmt.Println("updated")
			s.users[i].Name = name
			s.users[i].Password = password
			return nil
		}
	}

	// Appending the new user to the users slice if no duplicate is found.
	return errors.New("user not exists")
}

func (s *UserStore) DeleteUser(email string) error {
	for i, user := range s.users {
		if user.Email == email {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}

	// Appending the new user to the users slice if no duplicate is found.
	return errors.New("user not exists")
}

// CreateUser attempts to add a new user to the store.
// It checks for existing users with the same email to avoid duplicates and returns an error if found.
func (s *UserStore) CreateUser(name, email string, password string) error {
	for _, user := range s.users {
		if user.Email == email {
			// Preventing duplicate user registration.
			return errors.New("user already exists")
		}
	}

	// Appending the new user to the users slice if no duplicate is found.
	s.users = append(s.users, store.User{Name: name, Email: email, Password: password})
	return nil
}

// GetUser searches for a user by email and returns the user object if found.
// If no user is found with the provided email, it returns an error.
func (s *UserStore) GetUser(email string) (*store.User, error) {
	for _, user := range s.users {
		if user.Email == email {
			// Found the user, return a pointer to the User struct.
			return &user, nil
		}
	}

	// If the loop completes without finding a user, return an error indicating this.
	return nil, errors.New("user not found")
}

func (s *UserStore) ListUsers() []store.User {
	return s.users[:]
}

type UserPagination struct {
	baseUrl      string
	curPage      int
	itemsPerPage int
	store        store.UserStore
}

func NewUserPagination(url string, store store.UserStore, pg int) UserPagination {
	return UserPagination{baseUrl: url,
		curPage:      pg,
		itemsPerPage: 5,
		store:        store,
	}
}

func (thing UserPagination) PageUrl(page int) string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, page)
}

func (thing UserPagination) Pages() []int {
	var ret []int
	for i := 0; i < thing.PageCount(); i++ {
		ret = append(ret, i+1)
	}
	return ret
}

func (thing UserPagination) Items() []store.User {
	var ret []store.User
	k := (thing.curPage - 1) * thing.itemsPerPage
	for i, u := range thing.store.ListUsers() {
		if i >= k && i < k+thing.itemsPerPage {
			ret = append(ret, u)
		}
	}
	return ret
}

func (thing UserPagination) PageCount() int {
	k := thing.Total() / thing.itemsPerPage
	if k%thing.itemsPerPage != 0 {
		k += 1
	}
	return k
}

func (thing UserPagination) CurrentPage() int {
	return thing.curPage
}

func (thing UserPagination) PerPage() int {
	return thing.itemsPerPage
}

func (thing UserPagination) PreviousPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, thing.curPage-1)
}

func (thing UserPagination) Total() int {
	return len(thing.store.ListUsers())
}

func (thing UserPagination) NextPageUrl() string {
	return fmt.Sprintf("%s?page=%d", thing.baseUrl, thing.curPage+1)
}