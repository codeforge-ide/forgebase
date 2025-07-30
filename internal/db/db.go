package db

// Placeholder for database abstraction

type Database interface {
	Init() error
	Close() error
	// Add CRUD methods here
}

func NewDatabase() Database {
	// TODO: return a real implementation
	return nil
}
