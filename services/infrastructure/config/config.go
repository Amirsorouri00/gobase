package config

// Adapter is the interface for configuration adapters.
type Adapter interface {

	// String returns the string value for a given key.
	String(key string) (value string, ok bool)

	// Int returns the int value for a given key.
	Int(key string) (value int, ok bool)

	// Bool returns the bool value for a given key.
	Bool(key string) (value bool, ok bool)
}

var adapters []Adapter

// Init sets up the configuration adapters.
func Init(ads ...Adapter) {
	adapters = ads
}

// String returns the string value for a given key.
func String(key string) string {
	return StringWithDefault(key, "")
}

// StringWithDefault returns the string value for a given key. If the key is
// not found, the default value is returned.
func StringWithDefault(key, def string) string {
	for _, adapter := range adapters {
		if v, ok := adapter.String(key); ok {
			return v
		}
	}
	return def
}

// Int returns the int value for a given key.
func Int(key string) int {
	return IntWithDefault(key, 0)
}

// IntWithDefault returns the int value for a given key. If the key is not
// found, the default value is returned.
func IntWithDefault(key string, def int) int {
	for _, adapter := range adapters {
		if v, ok := adapter.Int(key); ok {
			return v
		}
	}
	return def
}

// Bool returns the bool value for a given key.
func Bool(key string) bool {
	return BoolWithDefault(key, false)
}

// BoolWithDefault returns the bool value for a given key. If the key is not
// found, the default value is returned.
func BoolWithDefault(key string, def bool) bool {
	for _, adapter := range adapters {
		if v, ok := adapter.Bool(key); ok {
			return v
		}
	}
	return def
}
