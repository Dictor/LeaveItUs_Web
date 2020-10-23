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

func LocalSqliteHandler(path string) *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(path), &gorm.Config{})
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

func ListPreloadTable(result interface{}, assoc []string) error {
	tx := currentDatabase
	for _, a := range assoc {
		tx = tx.Preload(a)
	}
	return tx.Find(result).Error
}

// SelectTable return row on table which has given primary key. result is pointer of every type. pk is primary key value.
func SelectTable(result interface{}, pk interface{}) error {
	return currentDatabase.First(result, pk).Error
}

// DeleteRow deletes rows, row can be every type of pointer
func DeleteRow(row interface{}) error {
	return currentDatabase.Delete(row).Error
}

// DeleteRowByKeys deletes row by primary key. model can be every type of pointer.
func DeleteRowByKeys(model interface{}, keys interface{}) error {
	return currentDatabase.Delete(model, keys).Error
}

// UpdateRow updates row, row can be every type of pointer and updated row is decided by pk.
func UpdateRow(row interface{}) error {
	return currentDatabase.Save(row).Error
}

// Migrate is migrate database with all models
func Migrate() {
	currentDatabase.AutoMigrate(&Tag{}, &Person{}, &Locker{})
}
