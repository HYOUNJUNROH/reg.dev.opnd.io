package testing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"runtime"
	"testing"
	"time"

	"git.dev.opnd.io/gc/backend-admin/pkg/config"
	"git.dev.opnd.io/gc/backend-admin/pkg/config/db"
	"git.dev.opnd.io/gc/backend-admin/pkg/config/elasticsearch"
	"git.dev.opnd.io/gc/backend-admin/pkg/config/kvstore"
	"git.dev.opnd.io/gc/backend-admin/pkg/config/s3"
	"git.dev.opnd.io/gc/backend-admin/pkg/logger"
	"golang.org/x/crypto/ssh"

	// elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	// elasticsearch8 "github.com/elastic/go-elasticsearch/v8"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type SetupHandler func() error

type TeardownHandler func() error

type TestDocker struct {
	UsePostgres      bool
	UseRedis         bool
	UseElasticSearch bool
	UseMinio         bool
	UseSftp          bool
	UseGoTrue        bool
	Setup            SetupHandler
	Teardown         TeardownHandler
	Expire           time.Duration
}

func NewTestDocker() *TestDocker {
	return &TestDocker{
		UsePostgres:      false,
		UseRedis:         false,
		UseElasticSearch: false,
		UseMinio:         false,
		UseSftp:          false,
		UseGoTrue:        false,
		Setup:            nil,
		Teardown:         nil,
		Expire:           240 * time.Second,
	}
}

func (t *TestDocker) testMain(m *testing.M) {
	config.Init()

	timeOut := uint(t.Expire / time.Second)

	code := func() int {
		// uses a sensible default on windows (tcp/http) and linux/osx (socket)
		if os.Getenv("DOCKER_HOST") == "" {
			switch runtime.GOOS {
			case "windows":
				os.Setenv("DOCKER_HOST", "npipe:////./pipe/docker_engine")
			case "linux":
				os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
			case "darwin":
				os.Setenv("DOCKER_HOST", fmt.Sprintf("unix://%v/.docker/run/docker.sock", os.Getenv("HOME")))
			default:
				os.Setenv("DOCKER_HOST", "tcp://localhost:2375")
			}
		}

		pool, err := dockertest.NewPool(os.Getenv("DOCKER_HOST"))
		if err != nil {
			log.Fatalf("Could not construct pool: %s", err)
		}

		err = pool.Client.Ping()
		if err != nil {
			log.Fatalf("Could not connect to Docker: %s", err)
		}

		config.Config.Cache.MasterName = "mymaster"
		config.Config.Cache.Adapter = "redis"
		config.Config.Cache.Host = "127.0.0.1"
		config.Config.Cache.User = "root"
		config.Config.Cache.Password = "changeme"
		// config.Config.Cache.LogMode = true
		config.Config.DB.Name = "postgres"
		config.Config.DB.Adapter = "postgres"
		config.Config.DB.Host = "127.0.0.1"
		config.Config.DB.User = "openerd"
		config.Config.DB.Password = "changeme"
		config.Config.DB.SSLMode = "disable"
		config.Config.DB.ReadReplicaHosts = "" // []string{}
		// config.Config.DB.LogMode = false
		config.Config.S3.Endpoint = "http://localhost.dev.opnd.io"
		config.Config.S3.AccessKeyID = "openerd"
		config.Config.S3.SecretAccessKey = "changeme"
		config.Config.S3.Region = "region"
		config.Config.S3.Bucket = "bucket"
		config.Config.S3.ForcePathStyle = false
		config.Config.ElasticSearch.User = "elastic"
		config.Config.ElasticSearch.Password = "changeme"
		config.Config.Sftp.Host = "127.0.0.1"
		config.Config.Sftp.User = "openerd"
		config.Config.Sftp.Password = "changeme"

		// postgres

		// pulls an image, creates a container based on it and runs it
		var dockerPostgres *dockertest.Resource
		if t.UsePostgres {
			initDir := path.Join(path.Dir(config.Config.RootDir), "schema", "database", "init")
			_ = os.MkdirAll(initDir, 0755)
			dockerPostgres, err = pool.RunWithOptions(&dockertest.RunOptions{
				Repository: "postgres",
				Tag:        "14",
				Cmd:        []string{"postgres", "-c", "max_connections=10000"},
				Env: []string{
					fmt.Sprintf("POSTGRES_PASSWORD=%v", config.Config.DB.Password),
					fmt.Sprintf("POSTGRES_USER=%v", config.Config.DB.User),
					fmt.Sprintf("POSTGRES_DB=%v", config.Config.DB.Name),
					"POSTGRES_INITDB_ARGS=--auth-host=scram-sha-256",
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{
					"5432/tcp": {
						{
							HostIP:   "",
							HostPort: "",
						},
					},
				},
				// Mounts: []string{
				// 	fmt.Sprintf("%v:/docker-entrypoint-initdb.d/", initDir),
				// },
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerPostgres: %s", err)
			}
			defer func() {
				if err := dockerPostgres.Close(); err != nil {
					log.Fatalf("Could not close dockerPostgres: %s", err)
				}
			}()
		}

		var dockerRedis *dockertest.Resource

		if t.UseRedis {
			// pulls an image, creates a container based on it and runs it
			dockerRedis, err = pool.RunWithOptions(&dockertest.RunOptions{
				Repository: "docker.io/bitnami/redis",
				Tag:        "6.2",
				Env: []string{
					fmt.Sprintf("REDIS_PASSWORD=%v", config.Config.Cache.Password),
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{
					"6379/tcp": {
						{
							HostIP:   "",
							HostPort: "",
						},
					},
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerRedis: %s", err)
			}
			defer func() {
				if err := dockerRedis.Close(); err != nil {
					log.Fatalf("Could not stop dockerRedis: %s", err)
				}
			}()
		}

		var dockerElasticsearch *dockertest.Resource
		if t.UseElasticSearch {
			dictionaryDir := path.Join(config.Config.RootDir, "tmp", "dictionary")
			err = os.MkdirAll(dictionaryDir, 0755)
			if err != nil {
				log.Fatalf("Could not create dictionaryDir: %s", err)
			}
			// pulls an image, creates a container based on it and runs it
			dockerElasticsearch, err = pool.RunWithOptions(&dockertest.RunOptions{
				Repository: "docker.io/bitnami/elasticsearch",
				Tag:        "8.6.2",
				Env: []string{
					"TZ=Asia/Seoul",
					"ELASTICSEARCH_ENABLE_SECURITY=true",
					"ELASTICSEARCH_NODE_NAME=es",
					"ELASTICSEARCH_SKIP_TRANSPORT_TLS=true",
					// "ELASTICSEARCH_PLUGINS=analysis-icu,analysis-nori",
					"ELASTICSEARCH_PLUGINS=analysis-nori",
					fmt.Sprintf("ELASTICSEARCH_USERNAME=%v", config.Config.ElasticSearch.User),
					fmt.Sprintf("ELASTICSEARCH_PASSWORD=%v", config.Config.ElasticSearch.Password),
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{
					"9200/tcp": {
						{
							HostIP:   "",
							HostPort: "",
						},
					},
				},
				Mounts: []string{
					fmt.Sprintf("%v:/opt/bitnami/elasticsearch/config/dictionary", dictionaryDir),
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerElasticsearch: %s", err)
			}
			defer func() {
				if err := dockerElasticsearch.Close(); err != nil {
					log.Fatalf("Could not stop dockerElasticsearch: %s", err)
				}
			}()
		}

		var dockerMinio *dockertest.Resource
		if t.UseMinio {
			dataDir := path.Join(config.Config.RootDir, "tmp", "s3")
			dataBucketDir := path.Join(dataDir, config.Config.S3.Bucket)
			err = os.MkdirAll(dataBucketDir, 0755)
			if err != nil {
				log.Fatalf("Could not create dataDir: %s", err)
			}
			u, err := url.Parse(config.Config.S3.Endpoint)
			if err != nil {
				log.Fatalf("Could not parse S3 endpoint: %s", err)
			}
			// pulls an image, creates a container based on it and runs it
			dockerMinio, err = pool.RunWithOptions(&dockertest.RunOptions{
				// Apache 2.0 until 2021-04-24
				// Tag: "RELEASE.2021-04-22T15-44-28Z.hotfix.56647434e",
				Repository: "docker.io/minio/minio",
				Tag:        "RELEASE.2021-04-18T19-26-29Z",
				Env: []string{
					"TZ=Asia/Seoul",
					fmt.Sprintf("MINIO_DOMAIN=%v", u.Hostname()),
					fmt.Sprintf("MINIO_ROOT_USER=%v", config.Config.S3.AccessKeyID),
					fmt.Sprintf("MINIO_ROOT_PASSWORD=%v", config.Config.S3.SecretAccessKey),
				},
				Cmd: []string{
					"server",
					"/data",
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{
					"9000/tcp": {
						{
							HostIP:   "",
							HostPort: "",
						},
					},
				},
				Mounts: []string{
					fmt.Sprintf("%v:/data", dataDir),
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerMinio: %s", err)
			}
			defer func() {
				if err := dockerMinio.Close(); err != nil {
					log.Fatalf("Could not stop dockerMinio: %s", err)
				}
			}()
		}

		var dockerSftp *dockertest.Resource
		if t.UseSftp {
			dataDir := path.Join(config.Config.RootDir, "tmp", "sftp")
			err = os.MkdirAll(dataDir, 0755)
			if err != nil {
				log.Fatalf("Could not create dataDir: %s", err)
			}
			// pulls an image, creates a container based on it and runs it
			dockerSftp, err = pool.RunWithOptions(&dockertest.RunOptions{
				// atmoz/sftp
				// reg.dev.opnd.io/library/sftp:latest
				Repository: "reg.dev.opnd.io/library/sftp",
				Tag:        "latest",
				Env: []string{
					"TZ=Asia/Seoul",
				},
				Cmd: []string{
					fmt.Sprintf("%v:%v:1001:1001:work", config.Config.Sftp.User, config.Config.Sftp.Password),
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{
					"22/tcp": {
						{
							HostIP:   "",
							HostPort: "",
						},
					},
				},
				Mounts: []string{
					// fmt.Sprintf("%v:/home/%v/work", dataDir, config.Config.Sftp.User),
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerSftp: %s", err)
			}
			defer func() {
				if err := dockerSftp.Close(); err != nil {
					log.Fatalf("Could not stop dockerSftp: %s", err)
				}
			}()
		}

		var dockerGoTrue *dockertest.Resource
		if t.UseGoTrue {

			// pulls an image, creates a container based on it and runs it
			dockerGoTrue, err = pool.RunWithOptions(&dockertest.RunOptions{
				Repository: "docker.io/supabase/gotrue",
				Tag:        "v2.31.0",
				Env: []string{
					"TZ=Asia/Seoul",
					"GOTRUE_SITE_URL=empty",
					"GOTRUE_JWT_SECRET=empty",
					"GOTRUE_DB_DRIVER=postgres",
					fmt.Sprintf("GOTRUE_DB_DATABASE_URL=postgres://%v:%v@%v:%v/%v?search_path=auth",
						config.Config.DB.User,
						config.Config.DB.Password,
						dockerPostgres.Container.NetworkSettings.IPAddress,
						"5432",
						// config.Config.DB.Host,
						// dbPort,
						// config.Config.DB.Port,
						config.Config.DB.Name,
					),
				},
				Cmd: []string{
					"gotrue",
					"migrate",
				},
				Labels: map[string]string{
					"temp": "true",
				},
				PortBindings: map[docker.Port][]docker.PortBinding{},
				Mounts:       []string{
					// fmt.Sprintf("%v:/home/%v/work", dataDir, config.Config.Sftp.User),
				},
			}, func(config *docker.HostConfig) {
				// set AutoRemove to true so that stopped container goes away by itself
				config.AutoRemove = false
				config.RestartPolicy = docker.RestartPolicy{Name: "on-failure"}
			})
			if err != nil {
				log.Fatalf("Could not start dockerSftp: %s", err)
			}
			defer func() {
				if err := dockerGoTrue.Close(); err != nil {
					log.Fatalf("Could not stop dockerGoTrue: %s", err)
				}
			}()
		}

		////////////////////////////////////////////////////////////////////////////
		if t.UsePostgres {
			dockerPostgres.Expire(timeOut) // Tell docker to hard kill the container in timeOut seconds

			var db_ *sql.DB
			// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
			pool.MaxWait = t.Expire
			if err = pool.Retry(func() error {
				r, err := pool.Client.InspectContainer(dockerPostgres.Container.ID)
				if err != nil {
					return err
				}
				dockerPostgres.Container = r

				config.Config.DB.Port = dockerPostgres.GetPort("5432/tcp")
				databaseUrl := fmt.Sprintf("postgres://%v:%v@%v:%s/%v?sslmode=disable", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Name)

				fmt.Println("Trying to connect to database", databaseUrl)

				if config.Config.DB.Port == "" || config.Config.DB.Port == "0" {
					return errors.New("postgresql port is not ready")
				}

				db_, err = sql.Open("postgres", databaseUrl)
				if err != nil {
					fmt.Println(err)
					return err
				}
				err = db_.Ping()
				return err
			}); err != nil {
				log.Fatalf("Could not connect to docker: %s", err)
			}

			if _, err := db_.Exec("CREATE SCHEMA IF NOT EXISTS auth;"); err != nil {
				log.Fatalf("Could not create schema: %s", err)
			}

			fmt.Println(dockerPostgres.Container.NetworkSettings)
			// docker run --rm -ti docker.io/supabase/gotrue:v2.31.0 gotrue migrate
			cmd := exec.Command(
				"docker",
				"run",
				"--network",
				"bridge",
				"--rm",
				"-t",
				"--label",
				"temp=true",
				"-e",
				"TZ=Asia/Seoul",
				"-e",
				"GOTRUE_DB_DRIVER=postgres",
				"-e",
				"GOTRUE_SITE_URL=empty",
				"-e",
				"GOTRUE_JWT_SECRET=empty",
				"-e",
				fmt.Sprintf("GOTRUE_DB_DATABASE_URL=postgres://%v:%v@%v:%v/%v?search_path=auth",
					config.Config.DB.User,
					config.Config.DB.Password,
					// strings.Trim(dockerPostgres.Container.Name, "/"),
					dockerPostgres.Container.NetworkSettings.IPAddress,
					"5432",
					// config.Config.DB.Host,
					// config.Config.DB.Port,
					config.Config.DB.Name,
				),
				"docker.io/supabase/gotrue:v2.31.0",
				"gotrue",
				"migrate",
			)
			fmt.Println(cmd.String())
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()

			if err != nil {
				logger.Logger.Errorf("Error running gotrue migrate: %v", err)
			}

			fmt.Printf("%v", path.Join(path.Dir(config.Config.RootDir), "schema", "database", "migrations"))
			driver, err := postgres.WithInstance(db_, &postgres.Config{})
			if err != nil {
				log.Fatalln(err)
			}
			m, err := migrate.NewWithDatabaseInstance(
				fmt.Sprintf("file://%v", path.Join(path.Dir(config.Config.RootDir), "schema", "database", "migrations")),
				"postgres", driver)
			if err != nil {
				log.Fatalln(err)
			}
			if err := m.Up(); err != nil {
				log.Fatalln(err)
			}

			db_.Close()
			db.Init()
		}

		// kvstore

		if t.UseRedis {
			dockerRedis.Expire(timeOut) // Tell docker to hard kill the container in timeOut seconds

			// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
			pool.MaxWait = t.Expire
			if err = pool.Retry(func() error {
				r, err := pool.Client.InspectContainer(dockerRedis.Container.ID)
				if err != nil {
					return err
				}
				dockerRedis.Container = r

				redisPort := dockerRedis.GetPort("6379/tcp")
				config.Config.Cache.Port = dockerRedis.GetPort("6379/tcp")

				log.Println("Connecting to redis on url: ", redisPort)

				if config.Config.Cache.Port == "" || config.Config.Cache.Port == "0" {
					return errors.New("redis port is empty")
				}

				kvstore.Init()
				return nil
			}); err != nil {
				log.Fatalf("Could not connect to docker: %s", err)
			}
		}

		// es

		if t.UseElasticSearch {
			dockerElasticsearch.Expire(timeOut) // Tell docker to hard kill the container in timeOut seconds

			// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
			pool.MaxWait = t.Expire
			if err = pool.Retry(func() error {
				r, err := pool.Client.InspectContainer(dockerElasticsearch.Container.ID)
				if err != nil {
					return err
				}
				dockerElasticsearch.Container = r

				esPort := dockerElasticsearch.GetPort("9200/tcp")
				config.Config.ElasticSearch.Addresses = fmt.Sprintf("http://127.0.0.1:%v", dockerElasticsearch.GetPort("9200/tcp"))

				log.Println("Connecting to elasticsearch on url: ", esPort)

				if esPort == "" || esPort == "0" {
					return errors.New("elasticsearch port is empty")
				}

				elasticsearch.Init()
				res, err := elasticsearch.EsClient.Ping(elasticsearch.EsClient.Ping.WithContext(context.WithValue(context.Background(), elasticsearch.LogRequest("log"), true)))
				if err != nil {
					return err
				}
				b, err := io.ReadAll(res.Body)
				// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(res.Status(), string(b))
				// es.Ping.WithContext(context.Background())(pingRequest)
				return nil
			}); err != nil {
				log.Fatalf("Could not connect to docker: %s", err)
			}

		}

		// s3

		if t.UseMinio {
			dockerMinio.Expire(timeOut) // Tell docker to hard kill the container in timeOut seconds

			// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
			pool.MaxWait = t.Expire
			if err = pool.Retry(func() error {
				r, err := pool.Client.InspectContainer(dockerMinio.Container.ID)
				if err != nil {
					return err
				}
				dockerMinio.Container = r

				s3Port := dockerMinio.GetPort("9000/tcp")
				config.Config.S3.Endpoint = fmt.Sprintf("http://localhost.dev.opnd.io:%v", s3Port)

				if s3Port == "" || s3Port == "0" {
					return errors.New("s3 port is empty")
				}

				url := fmt.Sprintf("%s/minio/health/live", config.Config.S3.Endpoint)
				fmt.Println("Trying to connect to s3", url)
				resp, err := http.Get(url)
				if err != nil {
					return err
				}
				if resp.StatusCode != http.StatusOK {
					return fmt.Errorf("status code not OK")
				}
				return nil
			}); err != nil {
				log.Fatalf("Could not connect to docker: %s", err)
			}
			s3.Init()
		}

		// sftp

		if t.UseSftp {
			dockerSftp.Expire(timeOut) // Tell docker to hard kill the container in timeOut seconds

			// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
			pool.MaxWait = t.Expire
			if err = pool.Retry(func() error {
				r, err := pool.Client.InspectContainer(dockerSftp.Container.ID)
				if err != nil {
					return err
				}
				dockerSftp.Container = r

				hostAndPort := dockerSftp.GetPort("22/tcp")
				config.Config.Sftp.Port = dockerSftp.GetPort("22/tcp")

				log.Println("Connecting to ssh on: ", hostAndPort)

				if config.Config.Sftp.Port == "" || config.Config.Sftp.Port == "0" {
					return errors.New("sftp port is empty")
				}

				addr := fmt.Sprintf("%v:%v", config.Config.Sftp.Host, config.Config.Sftp.Port)
				config := ssh.ClientConfig{
					User:            config.Config.Sftp.User,
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Auth: []ssh.AuthMethod{
						ssh.Password(config.Config.Sftp.Password),
					},
				}
				conn, err := ssh.Dial("tcp", addr, &config)
				if err != nil {
					fmt.Printf("Failed to connect to [%s]: %v", addr, err)
					return err
				}
				defer conn.Close()
				return nil
			}); err != nil {
				log.Fatalf("Could not connect to docker: %s", err)
			}
		}

		////////////////////////////////////////////////////////////////////////////
		//Run tests
		if t.Setup != nil {
			err = t.Setup()
			if err != nil {
				log.Fatalln(err)
			}
		}
		code := m.Run()
		if t.Teardown != nil {
			err = t.Teardown()
			if err != nil {
				log.Fatalln(err)
			}
		}
		return code
	}()
	os.Exit(code)
}

func (t *TestDocker) Start(m *testing.M) {
	t.testMain(m)
}
