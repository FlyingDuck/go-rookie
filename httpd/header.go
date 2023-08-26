package httpd

type Header map[string][]string

func (h Header) Add(key, value string) {
	h[key] = append(h[key], value)
}

func (h Header) Set(key, value string) {
	h[key] = []string{value}
}

func (h Header) Get(key string) string {
	if vals, existing := h[key]; existing && len(vals) != 0 {
		return vals[0]
	}
	return ""
}

func (h Header) Del(key string) {
	delete(h, key)
}
