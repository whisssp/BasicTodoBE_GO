package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{"DOING", "DONE", "DELETED"}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to Scan data from Sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to Scan data from Sql: %s", value))
	}

	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseStr2ItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
