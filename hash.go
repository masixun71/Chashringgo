package Chashringgo

const P = 16777619;

type hashAlgorithm interface {
	Hash(val string) uint32
}

type FNV1_32_HASH struct {
}

func (f *FNV1_32_HASH) Hash(str string) uint32 {

	hash := 2166136261;
	len := len(str)
	for i := 0; i < len; i++ {
		hash = (hash ^ int(str[i])) * P
	}
	hash += hash << 13
	hash ^= hash >> 7
	hash += hash << 3
	hash ^= hash >> 17
	hash += hash << 5

	if hash < 0 {
		hash = -hash
	}

	return uint32(hash)
}
