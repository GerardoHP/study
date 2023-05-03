package random

var hashLength = 10

type bucketElement struct {
	key   string
	value string
}

type Dictionary struct {
	buckets [][]bucketElement
	count   int
}

func NewDictionary() *Dictionary {
	buckets := make([][]bucketElement, hashLength)
	return &Dictionary{
		buckets: buckets,
	}
}

func (d *Dictionary) Get(key string) string {
	value, _, _ := d.internalGet(key)
	return value
}

func (d *Dictionary) Push(k, v string) string {
	currentV, hash, index := d.internalGet(k)
	bucket := &d.buckets[hash]

	if currentV != "" {
		(*bucket)[index].value = v
	}

	if currentV == "" {
		*bucket = append(*bucket, bucketElement{key: k, value: v})
		d.count++
	}

	return v
}

func (d *Dictionary) Delete(key string) (val string) {
	val, hash, index := d.internalGet(key)
	if val != "" {
		bucket := &d.buckets[hash]
		*bucket = append((*bucket)[:index], (*bucket)[index+1:]...)
		d.count--
	}

	return
}

func (d *Dictionary) Count() int {
	return d.count
}

func (d *Dictionary) internalGet(key string) (val string, hash int, index int) {
	hash = getHash(key)
	bucket := d.buckets[hash]
	for i, v := range bucket {
		if v.key == key {
			val = v.value
			index = i
			return
		}
	}

	return
}

func getHash(key string) int {
	return len(key) % hashLength
}
