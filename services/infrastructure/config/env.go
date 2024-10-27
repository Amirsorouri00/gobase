package config

import (
	"os"
	"strconv"
	"strings"
)

// envKeyConverter normalizes environment variables by replacing all invalid
// characters with underscores.
var envKeyConverter = strings.NewReplacer(".", "_", "-", "_", "/", "_", " ", "_")

func envvar(prefix, key string) string {
	var sb strings.Builder
	if prefix != "" {
		sb.WriteString(prefix)
		sb.WriteString("_")
	}
	sb.WriteString(envKeyConverter.Replace(key))
	return os.Getenv(sb.String())
}

type envAdapter struct {
	prefix string
}

// Env returns a configuration adapter that reads values from the environment
// variables.
func Env(prefix string) Adapter {
	return &envAdapter{prefix: prefix}
}

func (ad *envAdapter) String(key string) (string, bool) {
	v := envvar(ad.prefix, key)
	return v, v != ""
}

func (ad *envAdapter) Int(key string) (int, bool) {
	v := envvar(ad.prefix, key)
	if v == "" {
		return 0, false
	}
	i, err := strconv.Atoi(v)
	return i, err == nil
}

func (ad *envAdapter) Bool(key string) (bool, bool) {
	v := envvar(ad.prefix, key)
	if v == "" {
		return false, false
	}
	b, err := strconv.ParseBool(v)
	return b, err == nil
}
