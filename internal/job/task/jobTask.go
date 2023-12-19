package task

import (
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
		From(table.TbaAccessLog).
		OrderBy("id")

	// fetch all rows into a struct array
	var users []struct {
		ID  int64  `db:"id"`
		URL string `db:"url"`
	}
	err := q.All(&users)

	if err != nil {
		u.Logger.Error("", zap.Error(err))
	}
	fmt.Println(users[1].URL)
}
