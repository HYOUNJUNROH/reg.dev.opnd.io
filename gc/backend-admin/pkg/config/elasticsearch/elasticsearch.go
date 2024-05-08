package elasticsearch

import (

	// elasticsearch7 "github.com/elastic/go-elasticsearch/v7"

	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"git.dev.opnd.io/gc/backend-admin/pkg/config"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"

	//elasticsearch7 "github.com/elastic/go-elasticsearch/v7"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	EsClient *elasticsearch8.Client
	EsConfig elasticsearch8.Config
)

type LogRequest string

type LoggerTransport struct {
	transport http.RoundTripper
}

func (t *LoggerTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	lr := req.Context().Value(LogRequest("log"))
	if lr != nil && lr.(bool) == true {
		requestDump, err := httputil.DumpRequest(req, true)
		if err == nil {
			log.Printf("Request: %s", string(requestDump))
		}
	}

	return t.transport.RoundTrip(req)
}
func Init() {
	var err error
	EsConfig = elasticsearch8.Config{
		Addresses: strings.Split(config.Config.ElasticSearch.Addresses, ","),
		Username:  config.Config.ElasticSearch.User,
		Password:  config.Config.ElasticSearch.Password,
		// Transport: &LoggerTransport{transport: http.DefaultTransport},
	}
	EsClient, err = elasticsearch8.NewClient(EsConfig)
	if err != nil {
		log.Fatalf("Could not connect to ElasticSearch: %s", err)
	}
}
