package task

import (
	"encoding/json"
	"fiber-wire-template/internal/model"
	"fiber-wire-template/internal/repository"
	"fiber-wire-template/pkg/util/table"
	"fmt"
	"go.uber.org/zap"
)

type demoTask struct {
	*repository.Repository
}
type DemoTask interface {
	Run()
}

func NewDemoTask(r *repository.Repository) DemoTask {
	return &demoTask{r}
}

func (u *demoTask) Run() {
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
