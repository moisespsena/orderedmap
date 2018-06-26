package orderedmap

func mapStringToInterface(src map[string]interface{}) map[interface{}]interface{} {
	dst := make(map[interface{}]interface{})
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
