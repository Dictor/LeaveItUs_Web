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

// CreateRow create serveral rows, row can be every type of pointer
func CreateRow(row interface{}) error {
	return currentDatabase.Create(row).Error
}

// ListTable return all rows on table, table is decided by result parameter's reflect. And returing data is contained to it and it musb be every type of pointer.
func ListTable(result interface{}) error {
	return currentDatabase.Find(result).Error
}

func DeleteRow(row interface{}) error {
	return currentDatabase.Delete(row).Error
}

func DeleteRowByKeys(model interface{}, keys interface{}) error {
	return currentDatabase.Delete(model, keys).Error
}

// Migrate is migrate database with all models
func Migrate() {
	currentDatabase.AutoMigrate(&Tag{})
}