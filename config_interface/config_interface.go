package config_interface

import "time"

type SetConfigInterface interface {
	Set(key string, value any)
}

type GetConfigInterface interface {
	Get(key string, defaultValue any) any
	GetBool(key string, defaultValue bool) bool
	GetInt(key string, defaultValue int) int
	GetInt32(key string, defaultValue int32) int32
	GetInt64(key string, defaultValue int64) int64
	GetFloat64(key string, defaultValue float64) float64
	GetStringSlice(key string, defaultValue []string) []string
	GetString(key string, defaultValue string) string
	GetStringMapString(key string, defaultValue map[string]string) map[string]string
	GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string
	GetTime(key string, defaultValue time.Time) time.Time
}

type ConfigInterface interface {
	SetConfigInterface
	GetConfigInterface
}
