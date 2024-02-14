/*
Package config implements an configuration object. that inherits as follows:
- Defaults
- Environment variables
- Command line arguments

The names of the environment variables and command line flags, as well as the default values, are set using struct tags.
*/
package config

import (
	"flag"
	"fmt"
	"log/slog"
	"reflect"
	"strconv"
)

type Config struct {
	HttpHost                string `env:"HTTP_HOST" default:"localhost"`
	HttpPort                int    `env:"HTTP_PORT" default:"8080"`
	HttpShutdownGracePeriod int    `env:"HTTP_SHUTDOWN_GRACE_PERIOD" default:"30"`
	MonitorInterval         int    `env:"MONITOR_INTERVAL" default:"5"`
	Password                string `env:"PASSWORD" default:"" redact:"true"`
}

func (c *Config) LogValue() slog.Value {
	var values []slog.Attr
	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i++ {
		val := fmt.Sprintf("%v", v.Field(i))
		tag := v.Type().Field(i).Tag
		if tag.Get("redact") == "true" {
			val = "REDACTED"
		}
		values = append(values, slog.String(
			tag.Get("env"),
			val,
		))
	}
	return slog.GroupValue(values...)
}

func New(lookupenv func(string) (string, bool), args []string) (*Config, error) {
	c := Config{}
	flagset := flag.NewFlagSet(args[0], flag.ContinueOnError)
	if len(args) > 1 {
		if err := flagset.Parse(args[1:]); err != nil {
			return nil, err
		}
	}

	v := reflect.ValueOf(&c).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag
		if env := tag.Get("env"); env != "" {
			val := tag.Get("default")
			if lookupenv != nil {
				if v, ok := lookupenv(env); ok {
					val = v
				}
			}
			if f := flagset.Lookup(env); f != nil {
				val = f.Value.String()
			}
			switch field.Kind() {
			case reflect.String:
				field.SetString(val)
			case reflect.Int:
				v, err := strconv.Atoi(val)
				if err != nil {
					return nil, err
				}
				field.SetInt(int64(v))
			default:
				return nil, fmt.Errorf("unsupported type %s", field.Kind())
			}
		}
	}

	return &c, nil
}
