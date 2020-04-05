package jwt

type cacheProvider interface {
	// save stores JWT token entry in cache
	save(*entry) error

	// delete removes JWT token entry from cache
	delete(*entry) error

	// isValid checks validity of JWT token entry stored in cache
	isValid(*entry) (bool, error)
}

type dbProvider interface {
	init()

	// save stores JWT token entry in database
	save(*entry) error

	// delete removes JWT token entry from database
	delete(*entry) error

	// isValid checks validity of JWT token entry stored in database
	isValid(*entry) (bool, error)
}
