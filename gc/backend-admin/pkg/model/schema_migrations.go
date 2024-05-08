package model

import (
	"database/sql"
	"time"

	pq "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v9"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	_ = decimal.Decimal{}
	_ = pq.Int64Array{}
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
	_ = datatypes.JSON{}
	_ = gorm.DeletedAt{}
)

/*
DB Table Details
-------------------------------------


Table: schema_migrations
[ 0] version                                        INT8(14)             null: false  primary: true   isArray: false  auto: false  col: INT8            len: 14      default: []
[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []




*/

// SchemaMigrations_ struct is a row record of the schema_migrations table in the public database
type SchemaMigrations_ struct {
	//[ 0] version                                        INT8(14)             null: false  primary: true   isArray: false  auto: false  col: INT8            len: 14      default: []
	Version int64 `gorm:"primary_key;column:version;type:INT8;size:14;" json:"version"`
	//[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Dirty bool `gorm:"column:dirty;type:BOOL;" json:"dirty"`
}

var schema_migrationsTableInfo = &TableInfo{
	Name: "public.schema_migrations",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8(14)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       14,
			GoFieldName:        "Version",
			GoFieldType:        "int64",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "dirty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Dirty",
			GoFieldType:        "bool",
			JSONFieldName:      "dirty",
			ProtobufFieldName:  "dirty",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SchemaMigrations_) TableName() string {
	return "public.schema_migrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SchemaMigrations_) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SchemaMigrations_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SchemaMigrations_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SchemaMigrations_) TableInfo() *TableInfo {
	return schema_migrationsTableInfo
}
