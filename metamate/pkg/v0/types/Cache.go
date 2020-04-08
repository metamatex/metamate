package types

type Cache interface{
	Add(key string, value interface{})
	Get(key string) (interface{}, bool)
}