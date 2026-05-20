package user

import "errors"

type FailingRepo struct{}

func (r *FailingRepo) Save(p Person) error {
	return errors.New("repo save failed")
}
