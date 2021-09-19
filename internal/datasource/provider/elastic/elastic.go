package elastic

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/balerter/balerter/internal/config/datasources/elastic"
	"github.com/balerter/balerter/internal/script/script"
	ec "github.com/olivere/elastic/v7"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
)

/*

TODO:

	- implement provider to load config and connect + pin
	- implement manager / query

FUTURE
	- enhance yaml config:
			- aws region
			- client certs: ca certs / client cert / client key
			- doHealthCheck / healthcheckinterval

*/

var (
	defaultTimeout = time.Second * 5
)

// ModuleName returns the module name
func ModuleName(name string) string {
	return "elastic." + name
}

// Methods returns module methods
func Methods() []string {
	return []string{
		"query",
	}
}

type httpClient interface {
	CloseIdleConnections()
	Do(r *http.Request) (*http.Response, error)
}

// Prometheus represents the datasource of the type Prometheus
type Elastic struct {
	logger            *zap.Logger
	name              string
	host              string
	port              int
	useBasicAuth      bool
	basicAuthUsername string
	basicAuthPassword string
	timeout           time.Duration
	scheme            string
	sniff             bool
	client            *ec.Client
}

type elasticClient interface {
	Start()
	Stop()
}

type retrier struct {
	backoff ec.Backoff
}

func newRetrier(timeout time.Duration) *retrier {
	return &retrier{
		backoff: ec.NewExponentialBackoff(1*time.Millisecond, timeout),
	}
}

func (r *retrier) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	wait, stop := r.backoff.Next(retry)
	return wait, stop, nil
}

// New creates new Prometheus datasource
func New(cfg elastic.Elastic, logger *zap.Logger) (*Elastic, error) {
	m := &Elastic{
		logger:       logger,
		name:         ModuleName(cfg.Name),
		host:         cfg.Host,
		port:         cfg.Port,
		timeout:      time.Millisecond * time.Duration(cfg.Timeout),
		scheme:       cfg.Scheme,
		sniff:        cfg.Sniff == "true",
		useBasicAuth: false,
	}

	if cfg.BasicAuth != nil {
		m.useBasicAuth = true
		m.basicAuthUsername = cfg.BasicAuth.Username
		m.basicAuthPassword = cfg.BasicAuth.Password
	}

	if m.timeout == 0 {
		m.timeout = defaultTimeout
	}

	var err error

	m.client, err = ec.NewClient(ec.SetURL(m.clientURL()), ec.SetSniff(m.sniff), ec.SetRetrier(newRetrier(m.timeout)))
	if err != nil {
		return nil, err
	}
	_, _, err = m.client.Ping(m.clientURL()).Do(context.TODO())
	if err != nil {
		return nil, err
	}
	m.client.Start()

	return m, nil
}

// Generate client URL
func (m *Elastic) clientURL() string {
	return fmt.Sprintf("%v://%v:%v", m.scheme, m.host, m.port)
}

// Stop the datasource
func (m *Elastic) Stop() error {
	m.client.Stop()
	return nil
}

// Name returns the datasource name
func (m *Elastic) Name() string {
	return m.name
}

// GetLoader returns the datasource lua loader
func (m *Elastic) GetLoader(_ *script.Script) lua.LGFunction {
	return m.loader
}

func (m *Elastic) loader(luaState *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		//	"query": m.doQuery,
	}

	mod := luaState.SetFuncs(luaState.NewTable(), exports)

	luaState.Push(mod)
	return 1
}
