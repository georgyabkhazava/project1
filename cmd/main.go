package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/getByName", handlerGetByName)
	http.HandleFunc("/getByEmail", handlerGetUserByEmail)
	http.HandleFunc("/gika/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello gika!")
	})
	http.HandleFunc("/registre/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		email := r.URL.Query().Get("email")
		password := r.URL.Query().Get("password")

		if validateName(name) == false {
			fmt.Fprintf(w, "Error name")
			return
		}
		if validateEmail(email) == false {
			fmt.Fprintf(w, "Error")
			return
		}
		if validatePassword(password) == false {
			fmt.Fprintf(w, "Error")
			return
		}
		addUser(email, name, password)
		fmt.Fprintf(w, fmt.Sprintf("success addUser name: %s, email: %s, password: %s", name, email, password))
	})
	http.ListenAndServe(":80", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// проверить слайс пользователей, найти заданный емэйл, проверить совпадают ли значени пароля если есть такой написать ОК
func handlerLogin(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email") // получаем Квери параметр емэйл
	if validateEmail(email) == false {
		fmt.Fprintf(w, "Error")
		return
	}
	user := getEmail(email)
	password := r.URL.Query().Get("password")
	if user.Password == password {
		fmt.Fprintf(w, "OK")
	}

}

func handlerGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if validateEmail(email) == false {
		fmt.Fprintf(w, "Error")
		return
	}
	user := getEmail(email)
	fmt.Fprintf(w, fmt.Sprintf("name: %s, email: %s, password: %s", user.Name, user.Email, user.Password))
}

func handlerGetByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if validateName(name) == false {
		fmt.Fprintf(w, "Error")
		return
	}
	user := getUser(name)
	fmt.Fprintf(w, fmt.Sprintf("name: %s, email: %s, password: %s", user.Name, user.Email, user.Password))
}

func validateEmail(email string) bool {
	if len(email) < 4 || len(email) > 15 {
		return false
	}
	if !strings.Contains(email, "@") {
		return false
	}
	return true
}

func validateName(name string) bool {
	if name == "" {
		return false
	}
	return true
}

func validatePassword(password string) bool {
	if len(password) > 9 {
		return true
	}
	return false
}

type User struct {
	Email    string
	Name     string
	Password string
}

var users []User

func addUser(email string, name string, password string) {
	newUser := User{Email: email, Name: name, Password: password}
	users = append(users, newUser)
}

func getUser(name string) User {
	for _, user := range users {
		if user.Name == name {
			return user
		}
	}
	return User{}
}

func getEmail(email string) User {
	for _, user := range users {
		if user.Email == email {
			return user
		}
	}
	return User{}
}

func getByEmailAndPassword(email string, password string) User {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return user
		}
	}
	return User{}
}
