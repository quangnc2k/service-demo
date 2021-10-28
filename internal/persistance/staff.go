package persistance

import (
	"context"
	"git.cyradar.com/phinc/my-awesome-project/internal/model"
)

type StaffRepository interface {
	FindById(ctx context.Context, id string) (model.Staff, error)
	Insert(ctx context.Context, data *model.Staff) error
	Update(ctx context.Context, id string, data *model.Staff) error
	Remove(ctx context.Context, id string) error
}
