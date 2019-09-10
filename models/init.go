package models

import "github.com/evulse/petstore/db"

var database *db.MemoryDB //database

func init() {
	database = db.NewMemoryDB()
	database.CreateTable("pet")
	fido := Pet{
		Name: "Fido",
		PhotoUrls: []string{"https://www.purina.com.au/-/media/Project/Purina/Main/Breeds/Dog/Mobile/Dog_Dalmatian_Mobile.jpg?h=300&la=en&w=375&hash=F7C8808D6B26754DAFFD0A0CF56FFA5A"},
		Tags: Tags{{Id: 1, Name: "male"}},
		Category: Category{Id: 1, Name: "Dalmatian"},
	}
	fido.Save()
	lassie := Pet{
		Name: "Lassie",
		PhotoUrls: []string{"https://m.media-amazon.com/images/M/MV5BZTQxMzk4YzYtZTU0NS00YjI4LWJjMjQtM2FlOGNhOTNkYzA1XkEyXkFqcGdeQXVyMTExNDQ2MTI@._V1_UY317_CR13,0,214,317_AL_.jpg"},
		Tags: Tags{{Id: 2, Name: "female"}},
		Category: Category{Id: 15, Name: "Collie"},
	}
	lassie.Save()
}

func DB() *db.MemoryDB {
	return database
}