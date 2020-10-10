package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	currentDatabase *gorm.DB
)

// TestDBHander create inmemory sqlite instance for testing
func TestDBHander() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	return db
}

// SetDBHander set internal current dabatase singleton instance
func SetDBHander(h *gorm.DB) {
	currentDatabase = h
}

// CreateRow create serveral rows, row can be every type of pointer or object
func CreateRow(row interface{}) error {
	return currentDatabase.Create(row).Error
}

// ListTable return all rows on table, table is decided by result parameter's reflect. And returing data is contained to it
func ListTable(result interface{}) error {
	return currentDatabase.Find(result).Error
}

// Migrate is migrate database with all models
func Migrate() {
	currentDatabase.AutoMigrate(&Tag{})
}
