package adding

import "log"

type Repository interface {
	Add(d Doctor)
}

type Service interface {
	AddDoctor(d Doctor)
}

type service struct {
	r Repository
}

func NewService(r Repository) service {
	return service{r}
}

func (c service) AddDoctor(d Doctor) {
	log.Println("In create adding doctor")
	c.r.Add(d)

}
