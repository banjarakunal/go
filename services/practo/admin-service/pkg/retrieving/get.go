package retrieving

import (
	"log"

	"github.com/banjarakunal/go/services/go/practo/admin-services/pkg/adding"
)

type Repository interface {
	Get() []adding.Doctor
}

type Service interface {
	GetDocotsDetails() []adding.Doctor
}

type service struct {
	r Repository
}

func NewService(r Repository) service {
	return service{r}
}

func (c service) GetDocotsDetails() []adding.Doctor {
	log.Println("In create adding doctor")
	return c.r.Get()

}
