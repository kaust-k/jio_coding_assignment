package jwt

import (
	"jwt_server/db"
)

type postgresDbProvider struct {
}

func (cp *postgresDbProvider) init() {
	db := db.GetClient()
	db.Debug().AutoMigrate(&entry{})
}

func (cp *postgresDbProvider) save(e *entry) error {
	db := db.GetClient()
	return db.Debug().Create(e).Error
}

func (cp *postgresDbProvider) delete(e *entry) error {
	db := db.GetClient()
	return db.Debug().Delete(e).Error
}

func (cp *postgresDbProvider) isValid(e *entry) (bool, error) {
	db := db.GetClient()
	record := &entry{}
	db = db.Debug().Where(&entry{UserID: e.UserID, AuthID: e.AuthID}).First(record)
	return record.AuthID == e.AuthID, db.Error
}
