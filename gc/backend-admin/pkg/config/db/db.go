package db

import (
  "errors"
  "fmt"
  "os"
  "runtime/debug"
  "strings"

  "git.dev.opnd.io/gc/backend-admin/pkg/config"
  "git.dev.opnd.io/gc/backend-admin/pkg/logger"

  "gorm.io/driver/mysql"
  _ "gorm.io/driver/mysql"
  "gorm.io/driver/postgres"
  _ "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/plugin/dbresolver"
)

// DB Global DB connection
var DB *gorm.DB

// Unscoped DB
var UnscopedDB *gorm.DB

func deleteEmpty(s []string) []string {
  var r []string
  for _, str := range s {
    if str != "" {
      r = append(r, str)
    }
  }
  return r
}

func OpenDB1() (*gorm.DB, error) {
  var err error
  var DB *gorm.DB
  {
    dbConfig := config.Config.DB
    fmt.Println("dbConfig: ", dbConfig)
    if dbConfig.Adapter == "mysql" {
      DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)), &gorm.Config{})
      // DB = DB.Set("gorm:table_options", "CHARSET=utf8")
    } else if dbConfig.Adapter == "postgres" {
      DB, err = gorm.Open(postgres.Open(fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%s TimeZone=Asia/Seoul", dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Host, dbConfig.Port, dbConfig.SSLMode)), &gorm.Config{})
    } else {
      logger.Logger.Fatal(errors.New("not supported database adapter"))
    }
  }
  if err != nil {
    return nil, err
  }

  replicaHosts := strings.Split(config.Config.DB.ReadReplicaHosts, ",")
  replicaHosts = deleteEmpty(replicaHosts)

  dialectors := []gorm.Dialector{}
  for _, v := range replicaHosts {
    dbConfig := config.Config.DB
    if dbConfig.Adapter == "mysql" {
      dialectors = append(dialectors, mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, v, dbConfig.Port, dbConfig.Name)))
      // DB = DB.Set("gorm:table_options", "CHARSET=utf8")
    } else if dbConfig.Adapter == "postgres" {
      dialectors = append(dialectors, postgres.Open(fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%s TimeZone=Asia/Seoul", dbConfig.User, dbConfig.Password, dbConfig.Name, v, dbConfig.Port, dbConfig.SSLMode)))
    } else {
      logger.Logger.Fatal(errors.New("not supported database adapter"))
    }
  }

  err = DB.Use(dbresolver.Register(dbresolver.Config{
    Replicas: dialectors,
    Policy:   dbresolver.RandomPolicy{},
  }))
  if err != nil {
    return nil, err
  }
  return DB, nil
}

func Init() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in DB f", r)
      debug.PrintStack()
      os.Exit(2)
    }
  }()

  var err error

  if DB, err = OpenDB1(); err != nil {
    logger.Logger.Fatal(err)
  }

  UnscopedDB = DB.Unscoped()

  if config.Config.DB.LogMode {
    DB = DB.Debug()
  }
  if config.Config.DB.LogMode {
    UnscopedDB = UnscopedDB.Debug()
  }
}
