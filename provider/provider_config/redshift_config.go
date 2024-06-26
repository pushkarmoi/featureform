package provider_config

import (
	"encoding/json"
	"github.com/featureform/fferr"

	ss "github.com/featureform/helpers/string_set"
)

type RedshiftConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	SSLMode  string
}

func (rs *RedshiftConfig) Deserialize(config SerializedConfig) error {
	err := json.Unmarshal(config, rs)
	if err != nil {
		return fferr.NewInternalError(err)
	}
	return nil
}

func (rs *RedshiftConfig) Serialize() []byte {
	conf, err := json.Marshal(rs)
	if err != nil {
		panic(err)
	}
	return conf
}

func (rs RedshiftConfig) MutableFields() ss.StringSet {
	return ss.StringSet{
		"Username": true,
		"Password": true,
		"Port":     true,
		"SSLMode":  true,
	}
}

func (a RedshiftConfig) DifferingFields(b RedshiftConfig) (ss.StringSet, error) {
	return differingFields(a, b)
}
