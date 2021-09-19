package elastic

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	elastic "github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	// Create an Elasticsearch client
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), elastic.SetSniff(false), elastic.SetHealthcheckInterval(1*time.Second), elastic.SetTraceLog(log.New(os.Stdout, "", log.LstdFlags)), elastic.SetHealthcheck(true))
	require.NoError(t, err)
	_, _, err = client.Ping("http://127.0.0.1:9200").Do(context.TODO())
	require.NoError(t, err)
}
