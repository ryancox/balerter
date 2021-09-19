package elastic

import (
	"testing"

	"github.com/balerter/balerter/internal/config/common"
)

func TestDataSourceElastic_Validate(t *testing.T) {
	type fields struct {
		Name       string
		Host       string
		Port       int
		BasicAuth  *common.BasicAuth
		Timeout    int
		MaxRetries int
		Scheme     string
		Sniff      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		errText string
	}{
		{
			name:    "empty name",
			fields:  fields{Name: "", Host: "test.com"},
			wantErr: true,
			errText: "name must be not empty",
		},
		{
			name:    "ok",
			fields:  fields{Name: "test", Host: "test.com", Port: 9200, Scheme: "https", Sniff: "true"},
			wantErr: false,
			errText: "",
		},
		{
			name:    "missing host",
			fields:  fields{Name: "test", Port: 9200, Scheme: "https", Sniff: "true"},
			wantErr: true,
			errText: "host must be defined",
		},
		{
			name:    "missing port",
			fields:  fields{Name: "test", Host: "test.com", Scheme: "https", Sniff: "true"},
			wantErr: true,
			errText: "port must be defined",
		},
		{
			name:    "bad timeout",
			fields:  fields{Name: "test", Host: "test.com", Port: 9200, Scheme: "https", Sniff: "true", Timeout: -1},
			wantErr: true,
			errText: "timeout must be greater than 0",
		},
		{
			name:    "bad scheme",
			fields:  fields{Name: "test", Host: "test.com", Port: 9200, Scheme: "tcp", Sniff: "true"},
			wantErr: true,
			errText: "scheme must be either http or https",
		},
		{
			name:    "missing sniff",
			fields:  fields{Name: "test", Host: "test.com", Port: 9200, Scheme: "https"},
			wantErr: true,
			errText: "sniff must be either true or false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := Elastic{
				Name:      tt.fields.Name,
				Host:      tt.fields.Host,
				Port:      tt.fields.Port,
				BasicAuth: tt.fields.BasicAuth,
				Timeout:   tt.fields.Timeout,
				Scheme:    tt.fields.Scheme,
				Sniff:     tt.fields.Sniff,
			}
			err := cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errText {
				t.Errorf("Validate() error = '%s', wantErrText '%s'", err.Error(), tt.errText)
			}
		})
	}
}
