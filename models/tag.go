package models

type Tag struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

func (t Tag) Validate() error {
	return nil
}