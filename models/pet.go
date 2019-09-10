package models

import (
	"errors"
)

type Pet struct {
	Id        uint64 `json:"id"`
	Category  Category `json:"category,omitempty"`
	Name      string `json:"name,omitempty"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      Tags `json:"tags,omitempty"`
	Status    PetStatus `json:"status,omitempty"`
}

func (p *Pet) GetId() uint64 {
	return p.Id
}

func (p *Pet) SetId(id uint64)  {
	p.Id = id
}

func (p Pet) Validate() error {
	if p.PhotoUrls == nil || len(p.PhotoUrls) < 1 {
		return errors.New("at least one photoUrls must be defined")
	}
	if p.Name == "" {
		return errors.New("name can not be blank")
	}
	err := p.Status.Validate()
	if err != nil {
		return err
	}
	err = p.Category.Validate()
	if err != nil {
		return err
	}
	err = p.Tags.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (p *Pet) Save() error {
	err := p.Validate()
	if err != nil {
		return err
	}
	if p.Id == 0 {
		result := DB().Insert("pet", p)
		if result == nil {
			return errors.New("unable to save record")
		}
	} else {
		result := DB().Update("pet", p)
		if result == nil {
			return errors.New("unable to find record")
		}
	}
	return nil
}

func LoadPet(id uint64) (*Pet, error) {
	record := DB().Find("pet", id)
	if record == nil {
		return nil, errors.New("record does not exist")
	}
	pet, petOK := record.(*Pet)
	if !petOK {
		return nil, errors.New("record does not exist")
	}
	return pet, nil
}

func LoadPets() ([]*Pet, error) {
	record := DB().FindAll("pet")
	if len(record) == 0 {
		return nil, errors.New("no records exist")
	}
	pets := make([]*Pet, len(record))
	for idx := range record {
		pet, petOK := record[idx].(*Pet)
		if petOK {
			pets[idx] = pet
		}
	}

	return pets, nil
}


