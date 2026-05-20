package main

import (
	"fmt"
	"learn-go/13_structs_packages_design/user"
)

func main() {
	runWithInMemoryRepo()
	runWithFailingRepo()
}

func runWithInMemoryRepo() {
	repo := &user.InMemoryRepo{}
	svc := user.NewUserService(repo)
	u, err := svc.CreateUser("John", 30)
	if err != nil {
		fmt.Println("[InMemoryRepo] unexpected error:", err)
		return
	}
	u2, err := svc.CreateUser("jes", 30)
	if err != nil {
		fmt.Println("[InMemoryRepo] unexpected error:", err)
		return
	}

	fmt.Println("[InMemoryRepo] users:", u, u2)
	fmt.Println("[InMemoryRepo] count:", repo.Count())
}

func runWithFailingRepo() {
	repo := &user.FailingRepo{}
	svc := user.NewUserService(repo)
	_, err := svc.CreateUser("Alice", 25)
	if err != nil {
		fmt.Println("[FailingRepo] create user failed:", err)
	}
}
