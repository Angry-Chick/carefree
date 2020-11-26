package db

import (
	"fmt"

	"github.com/carefree/project/common/resource"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/protobuf/proto"
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
	b, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	// 不指定 id 默认自增长
	row := &Row{
		Name:     r.GetName(),
		Type:     string(r.ProtoReflect().Descriptor().FullName()),
		Resource: b,
	}
	nd := db.DB.Create(row)
	if err := nd.Error; err != nil {
		return nil, err
	}
	return getRow(nd)
}

func (db *DB) Update(r resource.Resource) (*Row, error) {
	b, err := proto.Marshal(r)
	if err != nil {
		return nil, err
	}
	row, err := db.Get(r.GetName())
	if err != nil {
		return nil, err
	}
	nd := db.DB.Model(&Row{Model: gorm.Model{ID: row.ID}}).Updates(Row{Resource: b})
	if nd.Error != nil {
		return nil, err
	}
	return getRow(nd)
}

func (db *DB) Delete(name string) error {
	row, err := db.Get(name)
	if err != nil {
		return err
	}
	return db.DB.Delete(&Row{Model: gorm.Model{ID: row.ID}}).Error
}

func (db *DB) Purge(name string) error {
	row, err := db.Get(name)
	if err != nil {
		return err
	}
	return db.DB.Unscoped().Delete(&Row{Model: gorm.Model{ID: row.ID}}).Error
}

func getRow(db *gorm.DB) (*Row, error) {
	var row Row
	if err := db.Scan(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}
