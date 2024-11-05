package edcclient

type AnyJSON map[string]interface{} // is a bit less generic than interface{} to be able to support unknown JSON input
