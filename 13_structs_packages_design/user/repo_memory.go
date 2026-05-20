package user

type InMemoryRepo struct {
	users []Person
}

func (r *InMemoryRepo) Save(p Person) error {
	r.users = append(r.users, p)
	return nil
}

func (r *InMemoryRepo) Count() int {
	return len(r.users)
}