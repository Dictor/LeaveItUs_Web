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

// Preload make transaction instance which preload given association data.
func Preload(assoc []string) *gorm.DB {
	tx := currentDatabase
	for _, a := range assoc {
		tx = tx.Preload(a)
	}
	return tx
}

// ListPreloadTable return all rows on table and fill association data with given assoc param. result is every type of pointer.
func ListPreloadTable(result interface{}, assoc []string) error {
	return Preload(assoc).Find(result).Error
}

// SelectTable return row on table which has given primary key. result is pointer of every type. pk is primary key value.
func SelectTable(result interface{}, param ...interface{}) error {
	return currentDatabase.Find(result, param...).Error
}

// SelectPreloadTable select rows on table from given condition param and fill association data with given assoc param.
// result is every type of pointer. param[0] is string sql condition and param[1:] is every type of value
func SelectPreloadTable(result interface{}, assoc []string, param ...interface{}) error {
	return Preload(assoc).Find(result, param...).Error
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
	currentDatabase.AutoMigrate(&Tag{}, &Person{}, &Locker{}, &LockerRecord{}, &LockerDoorEvent{})
}

func GetDatabase() *gorm.DB {
	return currentDatabase
}
