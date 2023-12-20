package task

import (
	"encoding/json"
	"fiber-wire-template/internal/model"
	"fiber-wire-template/internal/repository"
	"fiber-wire-template/pkg/util/table"
	"fmt"
	"go.uber.org/zap"
)

type jobTask struct {
	*repository.Repository
}
type JobTask interface {
	Run()
}

func NewJobTask(r *repository.Repository) JobTask {
	return &jobTask{r}
}

func (u *jobTask) Run() {
	q := u.Db.Select().
		From(table.TbaUser).
		OrderBy("id")

	// fetch all rows into a struct array
	var users []model.User
	err := q.All(&users)

	if err != nil {
		u.Logger.Error("", zap.Error(err))
	}
	for _, user := range users {
		fmt.Println(user)
	}
	v, _ := json.Marshal(users)
	fmt.Printf("%s\\n", v)
}
