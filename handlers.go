package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type RoutineRepo interface {
	FindRoutineById(routineId int) (*Routine, error)
	FindRoutinesByCreator(creatorId int) (Routines, error)
	FindRoutinesByLibrary(userId int) (Routines, error)
	CreateRoutine(r *Routine) (*Routine, error)
	AddRoutineToLibrary(userId int, routineId int) error
	RemoveRoutineFromLibrary(userId int, routineId int) error
	UpdateRoutine(r *Routine) (*Routine, error)
	DeleteRoutine(routineId int) error
	GetAllRoutines() (Routines, error)
}

type UserRepo interface {
	FindUserById(userId int) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(userId int) error
	GetAllUsers() ([]User, error)
}

const JSON = "application/json; charset=UTF-8"

func Login(w http.ResponseWriter, r *http.Request) {

	email, password, _ := r.BasicAuth()

	// TODO: if email not found, return 401.
	hash, _ := HashPassword("secret")
	user := RepoFindUser(1)

	if !CheckPasswordHash(password, hash) {
		w.Write("Bad login", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", JSON)

	token, err := GetJWT(&user)
	if err != nil {
		panic(err)
	}

	t := struct {
		Token string `json:"token"`
	}{token}

	json.NewEncoder(w).Encode(t)
}

func GetLibrary(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	fmt.Fprintf(w, "this is an authenticated request")
	fmt.Fprintf(w, "Claim content:\n")
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s : \t%#v\n", k, v)
	}

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		panic(err)
	}

	library := RepoFindLibrary(userId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(library); err != nil {
		panic(err)
	}
}

func GetRoutine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	routineId, err := strconv.Atoi(vars["routineId"])
	if err != nil {
		panic(err)
	}

	routine := RepoFindRoutine(routineId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(routine); err != nil {
		panic(err)
	}
}

func CreateRoutine(w http.ResponseWriter, r *http.Request) {
}

func ForkRoutine(w http.ResponseWriter, r *http.Request) {
}

func EditRoutine(w http.ResponseWriter, r *http.Request) {
}

func DeleteRoutine(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		panic(err)
	}

	user := RepoFindUser(userId)

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
}

func EditUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
