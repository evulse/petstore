package models

type Category struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

func (c Category) Validate() error {
	return nil
}