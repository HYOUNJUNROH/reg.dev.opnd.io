package config

import (
  "crypto/tls"
  "crypto/x509"
  "encoding/pem"
  "flag"
  "log"
  "net/url"
  "os"
  "path"
  "runtime"
  "strings"
  "time"

  awss3 "github.com/aws/aws-sdk-go/service/s3"
  "github.com/jinzhu/configor"
  "github.com/joho/godotenv"
  "github.com/lestrrat-go/strftime"
  amazonpay "github.com/qor/amazon-pay-sdk-go"
  "github.com/qor/gomerchant"
  "github.com/qor/mailer"
  "github.com/qor/media/oss"
  "github.com/qor/oss/s3"
  "github.com/qor/redirect_back"
  "github.com/qor/session/manager"
  "github.com/unrolled/render"

  "git.dev.opnd.io/gc/backend-admin/pkg/logger"
)

type SMTPConfig struct {
  Host     string
  Port     string
  User     string
  Password string
}

var CertKeyPair *tls.Certificate
var CertPool *x509.CertPool

var ExternalHost string
var ExternalMobileHost string

var LocalTimeZone, _ = time.LoadLocation("Asia/Seoul")
var DataPrefixPattern, _ = strftime.New(`%y/%m/%d`)
var Config = struct {
  HTTPS             bool   `default:"false" env:"HTTPS"`
  CookieDomain      string `default:"" env:"COOKIE_DOMAIN"`
  Port              uint   `default:"1323" env:"PORT"`
  NodeEnv           string `default:"development" env:"NODE_ENV"`
  ExternalURL       string `default:"http://localhost:3000" env:"EXTERNAL_URL"`
  ExternalMobileURL string `default:"http://localhost:3001" env:"EXTERNAL_MOBILE_URL"`
  TrustedCIDR       string `default:"127.0.0.1/32" env:"TRUSTED_CIDR"`
  Cache             struct {
    MasterName string `default:"mymaster" env:"CACHE_MASTER_NAME"`
    Adapter    string `env:"CACHE_ADAPTER"`
    Host       string `env:"CACHE_HOST"`
    Port       string `env:"CACHE_PORT"`
    User       string `env:"CACHE_USER"`
    Password   string `env:"CACHE_PASSWORD"`
    LogMode    bool
  }
  DB struct {
    Name             string `env:"DB_NAME" default:"qor_example"`
    Adapter          string `env:"DB_ADAPTER" default:"mysql"`
    Host             string `env:"DB_HOST" default:"localhost"`
    Port             string `env:"DB_PORT" default:"3306"`
    User             string `env:"DB_USER"`
    Password         string `env:"DB_PASSWORD"`
    SSLMode          string `env:"DB_SSL_MODE"`
    ReadReplicaHosts string `env:"DB_READ_REPLICA_HOSTS"`
    LogMode          bool   `env:"DB_LOG_MODE" default:"false"`
    Schema           string `env:"DB_SCHEMA" default:""`
  }
  S3 struct {
    AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" yaml:"access_key_id"`
    SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY" yaml:"secret_access_key"`
    Region          string `env:"AWS_REGION" yaml:"region"`
    Bucket          string `env:"AWS_BUCKET" yaml:"bucket"`
    Endpoint        string `env:"AWS_ENDPOINT" yaml:"endpoint"`
    ForcePathStyle  bool   `env:"AWS_FORCE_PATH_STYLE" yaml:"force_path_style"`
  }
  ElasticSearch struct {
    Addresses string `env:"ELASTICSEARCH_ADDRESSES" yaml:"addresses"`
    User      string `env:"ELASTICSEARCH_USER" yaml:"user"`
    Password  string `env:"ELASTICSEARCH_PASSWORD" yaml:"password"`
  }
  Sftp struct {
    Host     string `env:"SFTP_HOST" yaml:"host"`
    Port     string `env:"SFTP_PORT" yaml:"port"`
    User     string `env:"SFTP_USER" yaml:"user"`
    Password string `env:"SFTP_PASSWORD" yaml:"password"`
  }
  Supabase struct {
    Url string `env:"SUPABASE_URL" yaml:"supabase_url"`
    Key string `env:"SUPABASE_KEY" yaml:"supabase_key"`
  }
  GoTrue struct {
    JwtSecret string `env:"GOTRUE_JWT_SECRET" yaml:"gotrue_jwt_secret"`
  }
  CDN struct {
    DataCDNUrl string `env:"DATA_CDN_URL" yaml:"data_cdn_url"`
  }
  TLS struct {
    CertFile string `env:"TLS_CERT_FILE" yaml:"tls_cert_file" default:"cert/localhost.pem"`
    KeyFile  string `env:"TLS_KEY_FILE" yaml:"tls_key_file" default:"cert/localhost-key.pem"`
  }
  Dictionary struct {
    Method    string `env:"DICTIONARY_METHOD" yaml:"dictionary_method" default:""`
    Directory string `env:"DICTIONARY_DIRECTORY" yaml:"dictionary_directory" default:""`
  }
  ElasticSearchRemoteDirectory string `env:"ELASTICSEARCH_REMOTE_DIRECTORY" yaml:"elasticsearch_remote_directory" default:""`
  RootDir                      string
}{}

func IsDevelopment() bool {
  nodeenv := strings.TrimSpace(strings.ToLower(Config.NodeEnv))
  if nodeenv == "development" || nodeenv == "devel" || nodeenv == "dev" {
    return true
  }
  if nodeenv == "production" || nodeenv == "product" || nodeenv == "pro" {
    return false
  }
  return true
}

var (
  // Root           = os.Getenv("GOPATH") + "/src/git.dev.opnd.io/{{.GIR_GROUP}}/admin"
  Root, _        = os.Getwd()
  Mailer         *mailer.Mailer
  Render         = render.New()
  AmazonPay      amazonpay.AmazonPayService
  PaymentGateway gomerchant.PaymentGateway
  RedirectBack   = redirect_back.New(&redirect_back.Config{
    SessionManager:  manager.SessionManager,
    IgnoredPrefixes: []string{"/auth"},
  })
)

func IsTestRun() bool {
  return flag.Lookup("test.v") != nil
}

func Init() {
  //testing.Init()
  flag.Parse()

  _, filename, _, _ := runtime.Caller(0)
  dir := path.Join(path.Dir(filename), "..", "..")
  Config.RootDir = path.Join(path.Dir(filename), "..", "..")

  envFiles_ := []string{
    ".env",
    path.Join(dir, ".env"),
  }
  if IsDevelopment() {
    envFiles_ = []string{
      ".env",
      path.Join(dir, ".env"),
      path.Join(path.Dir(dir), ".env"),
      path.Join(path.Dir(path.Dir(dir)), ".env"),
    }
  }
  envFiles := []string{}

  for _, v := range envFiles_ {
    if _, err := os.Stat(v); err == nil {
      envFiles = append(envFiles, v)
    }
  }
  log.Println(envFiles)
  log.Println(IsTestRun())

  err := godotenv.Load(envFiles...)
  if err != nil {
    log.Println("Error loading .env file", err)
  }

  if err := configor.Load(&Config); err != nil {
    log.Fatal(err)
  }

  logger.Initialize(IsDevelopment())
  logger.Logger.Info("Is Development : ", IsDevelopment())
  logger.Logger.Info("Is Testing : ", IsTestRun())

  {
    if certificate, err := tls.LoadX509KeyPair(Config.TLS.CertFile, Config.TLS.KeyFile); err != nil {
      logger.Logger.Warn("failed to load key pair: %s", err)
    } else {
      CertKeyPair = &certificate
    }

    {
      CertPool := x509.NewCertPool()
      if cert, err := os.ReadFile(Config.TLS.CertFile); err != nil {
        logger.Logger.Warn("failed to load key pair: %s", err)
      } else {
        block, _ := pem.Decode(cert)
        Certs, err := x509.ParseCertificates(block.Bytes)
        if err != nil {
          logger.Logger.Fatalf("Failed to parse certificate:", err)
        }
        for _, c := range Certs {
          CertPool.AddCert(c)
        }
      }
    }
  }

  logger.Logger.Info("Use TLS : ", CertKeyPair != nil)

  {
    u, err := url.Parse(Config.ExternalURL)
    if err != nil {
      logger.Logger.Panic(err)
    }
    ExternalHost = u.Host
  }

  {
    u, err := url.Parse(Config.ExternalMobileURL)
    if err != nil {
      logger.Logger.Panic(err)
    }
    ExternalMobileHost = u.Host
  }

  if Config.S3.AccessKeyID != "" {
    oss.Storage = s3.New(&s3.Config{
      AccessID:         Config.S3.AccessKeyID,
      AccessKey:        Config.S3.SecretAccessKey,
      Region:           Config.S3.Region,
      Bucket:           Config.S3.Bucket,
      Endpoint:         Config.S3.Endpoint,
      S3Endpoint:       Config.S3.Endpoint,
      ACL:              awss3.BucketCannedACLPublicRead,
      S3ForcePathStyle: Config.S3.ForcePathStyle,
    })
  }
}
