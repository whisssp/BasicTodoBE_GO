package storage

import (
	"ToDoAppBackEnd/common"
	"ToDoAppBackEnd/modules/item/model"
	"context"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	deleteStatus := model.ItemStatusDeleted

	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": deleteStatus.String(),
	}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
