package utils

const fastStrMapSliceLimit = 1229
const fastStrMapSeed = 31

type fastStrMap struct {
	data [fastStrMapSliceLimit]string
	ref  [fastStrMapSliceLimit]string
}

func NewFastMap(m map[string]string) *fastStrMap {
	fm := &fastStrMap{}
	fm.setMap(m)
	return fm
}

func (fm *fastStrMap) setMap(m map[string]string) {
	for k, v := range m {
		fm.set(k, v)
	}
}

func (fm *fastStrMap) set(key, val string) {
	fm.data[strToInt(key)] = val
	fm.ref[strToInt(key)] = key
}

func (fm *fastStrMap) Get(key string) string {
	i := strToInt(key)
	if i >= fastStrMapSliceLimit || len(key) != len(fm.ref[i]) || fm.ref[i] != key {
		return ""
	}
	return fm.data[i]
}

// strToInt uses BKDRHash algorithm
func strToInt(s string) int {
	var hash int

	for _, b := range s {
		hash = hash*fastStrMapSeed + int(b)
	}
	return (hash & 0x7fffffff) % fastStrMapSliceLimit
}
