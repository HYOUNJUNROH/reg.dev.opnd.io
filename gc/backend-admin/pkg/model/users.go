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


Table: users
[ 0] instance_id                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
[ 1] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
[ 2] aud                                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] role                                           VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 4] email                                          VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 5] encrypted_password                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] email_confirmed_at                             TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 7] invited_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 8] confirmation_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] confirmation_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[10] recovery_token                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] recovery_sent_at                               TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[12] email_change_token_new                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] email_change                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] email_change_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[15] last_sign_in_at                                TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[16] raw_app_meta_data                              JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
[17] raw_user_meta_data                             JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
[18] is_super_admin                                 BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[19] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[20] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[21] phone                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: [NULL]
[22] phone_confirmed_at                             TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[23] phone_change                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[24] phone_change_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[25] phone_change_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[26] confirmed_at                                   TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[27] email_change_token_current                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[28] email_change_confirm_status                    INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: [0]
[29] banned_until                                   TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[30] reauthentication_token                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[31] reauthentication_sent_at                       TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[32] is_sso_user                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
[33] deleted_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []




*/

// Users struct is a row record of the users table in the auth database
type Users struct {
	//[ 0] instance_id                                    UUID                 null: true   primary: false  isArray: false  auto: false  col: UUID            len: -1      default: []
	InstanceID uuid.NullUUID `gorm:"column:instance_id;type:UUID;" json:"instance_id"`
	//[ 1] id                                             UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: []
	ID uuid.UUID `gorm:"primary_key;column:id;type:UUID;" json:"id"`
	//[ 2] aud                                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	Aud null.String `gorm:"column:aud;type:VARCHAR;size:255;" json:"aud"`
	//[ 3] role                                           VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Role null.String `gorm:"column:role;type:VARCHAR;" json:"role"`
	//[ 4] email                                          VARCHAR              null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	Email null.String `gorm:"column:email;type:VARCHAR;" json:"email"`
	//[ 5] encrypted_password                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EncryptedPassword null.String `gorm:"column:encrypted_password;type:VARCHAR;size:255;" json:"encrypted_password"`
	//[ 6] email_confirmed_at                             TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	EmailConfirmedAt null.Time `gorm:"column:email_confirmed_at;type:TIMESTAMPTZ;" json:"email_confirmed_at"`
	//[ 7] invited_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	InvitedAt null.Time `gorm:"column:invited_at;type:TIMESTAMPTZ;" json:"invited_at"`
	//[ 8] confirmation_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ConfirmationToken null.String `gorm:"column:confirmation_token;type:VARCHAR;size:255;" json:"confirmation_token"`
	//[ 9] confirmation_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	ConfirmationSentAt null.Time `gorm:"column:confirmation_sent_at;type:TIMESTAMPTZ;" json:"confirmation_sent_at"`
	//[10] recovery_token                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	RecoveryToken null.String `gorm:"column:recovery_token;type:VARCHAR;size:255;" json:"recovery_token"`
	//[11] recovery_sent_at                               TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	RecoverySentAt null.Time `gorm:"column:recovery_sent_at;type:TIMESTAMPTZ;" json:"recovery_sent_at"`
	//[12] email_change_token_new                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EmailChangeTokenNew null.String `gorm:"column:email_change_token_new;type:VARCHAR;size:255;" json:"email_change_token_new"`
	//[13] email_change                                   VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EmailChange null.String `gorm:"column:email_change;type:VARCHAR;size:255;" json:"email_change"`
	//[14] email_change_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	EmailChangeSentAt null.Time `gorm:"column:email_change_sent_at;type:TIMESTAMPTZ;" json:"email_change_sent_at"`
	//[15] last_sign_in_at                                TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	LastSignInAt null.Time `gorm:"column:last_sign_in_at;type:TIMESTAMPTZ;" json:"last_sign_in_at"`
	//[16] raw_app_meta_data                              JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
	RawAppMetaData datatypes.JSON `gorm:"column:raw_app_meta_data;type:JSONB;" json:"raw_app_meta_data"`
	//[17] raw_user_meta_data                             JSONB                null: true   primary: false  isArray: false  auto: false  col: JSONB           len: -1      default: []
	RawUserMetaData datatypes.JSON `gorm:"column:raw_user_meta_data;type:JSONB;" json:"raw_user_meta_data"`
	//[18] is_super_admin                                 BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	IsSuperAdmin null.Bool `gorm:"column:is_super_admin;type:BOOL;" json:"is_super_admin"`
	//[19] created_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:TIMESTAMPTZ;" json:"created_at"`
	//[20] updated_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:TIMESTAMPTZ;" json:"updated_at"`
	//[21] phone                                          TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: [NULL]
	Phone null.String `gorm:"column:phone;type:TEXT;" json:"phone"`
	//[22] phone_confirmed_at                             TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	PhoneConfirmedAt null.Time `gorm:"column:phone_confirmed_at;type:TIMESTAMPTZ;" json:"phone_confirmed_at"`
	//[23] phone_change                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	PhoneChange null.String `gorm:"column:phone_change;type:TEXT;" json:"phone_change"`
	//[24] phone_change_token                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	PhoneChangeToken null.String `gorm:"column:phone_change_token;type:VARCHAR;size:255;" json:"phone_change_token"`
	//[25] phone_change_sent_at                           TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	PhoneChangeSentAt null.Time `gorm:"column:phone_change_sent_at;type:TIMESTAMPTZ;" json:"phone_change_sent_at"`
	//[26] confirmed_at                                   TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	ConfirmedAt null.Time `gorm:"column:confirmed_at;type:TIMESTAMPTZ;" json:"confirmed_at"`
	//[27] email_change_token_current                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	EmailChangeTokenCurrent null.String `gorm:"column:email_change_token_current;type:VARCHAR;size:255;" json:"email_change_token_current"`
	//[28] email_change_confirm_status                    INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: [0]
	EmailChangeConfirmStatus null.Int32 `gorm:"column:email_change_confirm_status;type:INT2;default:0;" json:"email_change_confirm_status"`
	//[29] banned_until                                   TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	BannedUntil null.Time `gorm:"column:banned_until;type:TIMESTAMPTZ;" json:"banned_until"`
	//[30] reauthentication_token                         VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	ReauthenticationToken null.String `gorm:"column:reauthentication_token;type:VARCHAR;size:255;" json:"reauthentication_token"`
	//[31] reauthentication_sent_at                       TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	ReauthenticationSentAt null.Time `gorm:"column:reauthentication_sent_at;type:TIMESTAMPTZ;" json:"reauthentication_sent_at"`
	//[32] is_sso_user                                    BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: [false]
	IsSsoUser bool `gorm:"column:is_sso_user;type:BOOL;default:false;" json:"is_sso_user"`
	//[33] deleted_at                                     TIMESTAMPTZ          null: true   primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	DeletedAt null.Time `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deleted_at"`
}

var usersTableInfo = &TableInfo{
	Name: "auth.users",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "instance_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "UUID",
			DatabaseTypePretty: "UUID",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "UUID",
			ColumnLength:       -1,
			GoFieldName:        "InstanceID",
			GoFieldType:        "uuid.NullUUID",
			JSONFieldName:      "instance_id",
			ProtobufFieldName:  "instance_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "UUID",
			DatabaseTypePretty: "UUID",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "UUID",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uuid.UUID",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "aud",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Aud",
			GoFieldType:        "null.String",
			JSONFieldName:      "aud",
			ProtobufFieldName:  "aud",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "role",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "Role",
			GoFieldType:        "null.String",
			JSONFieldName:      "role",
			ProtobufFieldName:  "role",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		{
			Index:              5,
			Name:               "encrypted_password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "EncryptedPassword",
			GoFieldType:        "null.String",
			JSONFieldName:      "encrypted_password",
			ProtobufFieldName:  "encrypted_password",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		{
			Index:              6,
			Name:               "email_confirmed_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "EmailConfirmedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "email_confirmed_at",
			ProtobufFieldName:  "email_confirmed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		{
			Index:              7,
			Name:               "invited_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "InvitedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "invited_at",
			ProtobufFieldName:  "invited_at",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		{
			Index:              8,
			Name:               "confirmation_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ConfirmationToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "confirmation_token",
			ProtobufFieldName:  "confirmation_token",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		{
			Index:              9,
			Name:               "confirmation_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "ConfirmationSentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "confirmation_sent_at",
			ProtobufFieldName:  "confirmation_sent_at",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		{
			Index:              10,
			Name:               "recovery_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RecoveryToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "recovery_token",
			ProtobufFieldName:  "recovery_token",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		{
			Index:              11,
			Name:               "recovery_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "RecoverySentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "recovery_sent_at",
			ProtobufFieldName:  "recovery_sent_at",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		{
			Index:              12,
			Name:               "email_change_token_new",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "EmailChangeTokenNew",
			GoFieldType:        "null.String",
			JSONFieldName:      "email_change_token_new",
			ProtobufFieldName:  "email_change_token_new",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		{
			Index:              13,
			Name:               "email_change",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "EmailChange",
			GoFieldType:        "null.String",
			JSONFieldName:      "email_change",
			ProtobufFieldName:  "email_change",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		{
			Index:              14,
			Name:               "email_change_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "EmailChangeSentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "email_change_sent_at",
			ProtobufFieldName:  "email_change_sent_at",
			ProtobufType:       "uint64",
			ProtobufPos:        15,
		},

		{
			Index:              15,
			Name:               "last_sign_in_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "LastSignInAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_sign_in_at",
			ProtobufFieldName:  "last_sign_in_at",
			ProtobufType:       "uint64",
			ProtobufPos:        16,
		},

		{
			Index:              16,
			Name:               "raw_app_meta_data",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSONB",
			DatabaseTypePretty: "JSONB",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSONB",
			ColumnLength:       -1,
			GoFieldName:        "RawAppMetaData",
			GoFieldType:        "datatypes.JSON",
			JSONFieldName:      "raw_app_meta_data",
			ProtobufFieldName:  "raw_app_meta_data",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		{
			Index:              17,
			Name:               "raw_user_meta_data",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSONB",
			DatabaseTypePretty: "JSONB",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSONB",
			ColumnLength:       -1,
			GoFieldName:        "RawUserMetaData",
			GoFieldType:        "datatypes.JSON",
			JSONFieldName:      "raw_user_meta_data",
			ProtobufFieldName:  "raw_user_meta_data",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		{
			Index:              18,
			Name:               "is_super_admin",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "IsSuperAdmin",
			GoFieldType:        "null.Bool",
			JSONFieldName:      "is_super_admin",
			ProtobufFieldName:  "is_super_admin",
			ProtobufType:       "bool",
			ProtobufPos:        19,
		},

		{
			Index:              19,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        20,
		},

		{
			Index:              20,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        21,
		},

		{
			Index:              21,
			Name:               "phone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Phone",
			GoFieldType:        "null.String",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		{
			Index:              22,
			Name:               "phone_confirmed_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "PhoneConfirmedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "phone_confirmed_at",
			ProtobufFieldName:  "phone_confirmed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        23,
		},

		{
			Index:              23,
			Name:               "phone_change",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "PhoneChange",
			GoFieldType:        "null.String",
			JSONFieldName:      "phone_change",
			ProtobufFieldName:  "phone_change",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		{
			Index:              24,
			Name:               "phone_change_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "PhoneChangeToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "phone_change_token",
			ProtobufFieldName:  "phone_change_token",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		{
			Index:              25,
			Name:               "phone_change_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "PhoneChangeSentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "phone_change_sent_at",
			ProtobufFieldName:  "phone_change_sent_at",
			ProtobufType:       "uint64",
			ProtobufPos:        26,
		},

		{
			Index:              26,
			Name:               "confirmed_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "ConfirmedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "confirmed_at",
			ProtobufFieldName:  "confirmed_at",
			ProtobufType:       "uint64",
			ProtobufPos:        27,
		},

		{
			Index:              27,
			Name:               "email_change_token_current",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "EmailChangeTokenCurrent",
			GoFieldType:        "null.String",
			JSONFieldName:      "email_change_token_current",
			ProtobufFieldName:  "email_change_token_current",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		{
			Index:              28,
			Name:               "email_change_confirm_status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "EmailChangeConfirmStatus",
			GoFieldType:        "null.Int32",
			JSONFieldName:      "email_change_confirm_status",
			ProtobufFieldName:  "email_change_confirm_status",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		{
			Index:              29,
			Name:               "banned_until",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "BannedUntil",
			GoFieldType:        "null.Time",
			JSONFieldName:      "banned_until",
			ProtobufFieldName:  "banned_until",
			ProtobufType:       "uint64",
			ProtobufPos:        30,
		},

		{
			Index:              30,
			Name:               "reauthentication_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ReauthenticationToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "reauthentication_token",
			ProtobufFieldName:  "reauthentication_token",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		{
			Index:              31,
			Name:               "reauthentication_sent_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "ReauthenticationSentAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "reauthentication_sent_at",
			ProtobufFieldName:  "reauthentication_sent_at",
			ProtobufType:       "uint64",
			ProtobufPos:        32,
		},

		{
			Index:              32,
			Name:               "is_sso_user",
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
			GoFieldName:        "IsSsoUser",
			GoFieldType:        "bool",
			JSONFieldName:      "is_sso_user",
			ProtobufFieldName:  "is_sso_user",
			ProtobufType:       "bool",
			ProtobufPos:        33,
		},

		{
			Index:              33,
			Name:               "deleted_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "DeletedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "deleted_at",
			ProtobufFieldName:  "deleted_at",
			ProtobufType:       "uint64",
			ProtobufPos:        34,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "auth.users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users) TableInfo() *TableInfo {
	return usersTableInfo
}
