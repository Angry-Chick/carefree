package db

import (
	"fmt"

	"github.com/carefree/project/common/resource"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type sql string

type Config struct {
	SQL          string `yaml:"db"`
	Host         string `yaml:"host"`
	Username     string `yaml:"user"`
	Password     string `yaml:"pwd"`
	DBName       string `yaml:"dbname"`
	DefaultTable string `yaml:"table"`
	Port         string `yaml:"port"`
}

// DB is a wrapper over gorm.DB with additional methods.
type DB struct {
	*gorm.DB
}

// New returns a db client.
func New(cfg *Config) (*DB, error) {
	var (
		err error
		db  *gorm.DB
	)
	switch cfg.SQL {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.DBName)
		db, err = gorm.Open("mysql", dsn)
	default:
		return nil, fmt.Errorf("don't support database %s", cfg.SQL)
	}
	if err != nil {
		return nil, err
	}

	return &DB{db.Table(cfg.DefaultTable)}, nil
}

func (db *DB) Close() {
	db.Close()
}

func (db *DB) Begin() *DB {
	return &DB{db.DB.Begin()}
}

func (db *DB) Transaction(fc func(db *DB) error) (err error) {
	panicked := true
	tx := &DB{db.DB.Begin()}

	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	err = fc(tx)
	if err == nil {
		err = tx.DB.Commit().Error
	}
	panicked = false
	return
}

func (db *DB) Commit() error {
	if err := db.DB.Commit().Error; err != nil {
		return err
	}
	return nil
}

// Row indicates data stored in a resource table row.
type Row struct {
	gorm.Model
	Name     string `gorm:"index;unique;not null"`
	Type     string
	Resource []byte
}

func (db *DB) Table(table string) {
	db.DB = db.DB.Table(table)
}

func (db *DB) Get(name string) (*Row, error) {
	nd := db.DB.Where("name= ?", name).First(&Row{})
	if err := nd.Error; err != nil {
		return nil, err
	}
	return getRow(nd)
}

func (db *DB) Create(r resource.Resource) (*Row, error) {
	b, err := Marshal(r)
	if err != nil {
		return nil, err
	}
	row := &Row{
		Name:     r.GetName(),
		Type:     proto.MessageName(r),
		Resource: b,
	}
	nd := db.DB.Create(row)
	if err := nd.Error; err != nil {
		return nil, err
	}
	return getRow(nd)
}

func (db *DB) Update(r resource.Resource) (*Row, error) {
	b, err := Marshal(r)
	if err != nil {
		return nil, err
	}
	nd := db.DB.Model(&Row{Name: r.GetName()}).Update("resource", b)
	if nd.Error != nil {
		return nil, err
	}
	return getRow(nd)
}

func (db *DB) Delete(name string) error {
	return db.DB.Delete(&Row{Name: name}).Error
}

func getRow(db *gorm.DB) (*Row, error) {
	var row Row
	if err := db.Scan(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
