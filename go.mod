module github.com/balerter/balerter

go 1.17

require (
	github.com/ClickHouse/clickhouse-go v1.3.13
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/aws/aws-sdk-go v1.40.32
	github.com/deckarep/gosx-notifier v0.0.0-20180201035817-e127226297fb // indirect
	github.com/diamondburned/arikawa v0.9.2
	github.com/fatih/color v1.7.0
	github.com/go-chi/chi v1.5.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/grafana/loki v1.3.0
	github.com/hashicorp/hcl/v2 v2.9.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.0.0
	github.com/martinlindhe/notify v0.0.0-20181008203735-20632c9a275a
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/mavolin/dismock v1.0.0
	github.com/nlopes/slack v0.6.0
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/prometheus/alertmanager v0.19.0
	github.com/prometheus/client_golang v1.4.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.9.1
	github.com/robfig/cron/v3 v3.0.0
	github.com/stretchr/testify v1.6.1
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	github.com/xhit/go-simple-mail/v2 v2.3.1
	github.com/yuin/gluamapper v0.0.0-20150323120927-d836955830e7
	github.com/yuin/gopher-lua v0.0.0-20191220021717-ab39c6098bdb
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	gonum.org/v1/plot v0.0.0-20200212202559-4d97eda4de95
	gopkg.in/toast.v1 v1.0.0-20180812000517-0a84660828b2 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

// Override reference causing proxy error.  Otherwise it attempts to download https://proxy.golang.org/golang.org/x/net/@v/v0.0.0-20190813000000-74dc4d7220e7.info
// See repo github.com/grafana/loki
replace golang.org/x/net => golang.org/x/net v0.0.0-20190923162816-aa69164e4478
