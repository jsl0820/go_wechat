package wechat

import (
	
)

const 

const DEFAULT_CACHE = "file"

type Cache interface{
	Set(key string, value string, ex) (bool, error)
	Get(key string) string 
	Delete(key string) string
	Has(key string) bool
	Clear() bool
}






