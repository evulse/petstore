package models

type Tags []Tag

func (t Tags) Validate() error {
	for idx := range t {
		err := t[idx].Validate()
		if err != nil {
			return err
		}
	}
	return nil
}