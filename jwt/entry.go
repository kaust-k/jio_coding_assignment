package jwt

import (
	"sync"
	"time"
)

type entry struct {
	UserID  string `gorm:"primary_key"`
	AuthID  string `gorm:"primary_key"` // To support login / logout per device
	Expires time.Time
}

func newEntry(userID, authID string, expiry time.Time) *entry {
	return &entry{
		UserID:  userID,
		AuthID:  authID,
		Expires: expiry,
	}
}

func (e *entry) saveToCache(provider cacheProvider) error {
	return provider.save(e)
}

func (e *entry) saveToDb(provider dbProvider) error {
	return provider.save(e)
}

func (e *entry) deleteFromCache(provider cacheProvider) error {
	return provider.delete(e)
}

func (e *entry) deleteFromDb(provider dbProvider) error {
	return provider.delete(e)
}

// save stores the JWT entry in database and cache
func (e *entry) save(cProvider cacheProvider, provider dbProvider) error {
	err := e.saveToDb(provider)
	if err == nil {
		err = e.saveToCache(cProvider)
	}
	return err
}

// delete removes JWT entry from database and cache
func (e *entry) delete(cProvider cacheProvider, dProvider dbProvider) error {
	var wg sync.WaitGroup
	var err1, err2 error
	wg.Add(2)

	go func() {
		err1 = e.deleteFromCache(cProvider)
		wg.Done()
	}()

	go func() {
		err2 = e.deleteFromDb(dProvider)
		wg.Done()
	}()

	wg.Wait()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

// isValid checks validity of JWT entry stored in database and cache
// If the entry is considered as invalid, if it is not present in cache as well as database
func (e *entry) isValid(cProvider cacheProvider, dProvider dbProvider) (bool, error) {
	valid, _ := cProvider.isValid(e)
	if valid {
		return valid, nil
	}

	valid, err := dProvider.isValid(e)

	if valid {
		go cProvider.save(e)
	}

	return valid, err
}
