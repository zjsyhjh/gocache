package hash

import "hash/crc32"
import "strconv"
import "sort"

/*
 * 默认为crc32.CheckSumIEEE
 */
type Algorithm func(data []byte) uint32

/*
 * 一致性hash算法, hash算法采用CheckSumIEEE
 * replicas是重复份数
 * sortedKey是一个保存key的有序数组
 */
type RingHash struct {
	algorithm  Algorithm
	replicas   int
	sortedKeys []int
	hashMap    map[int]string
}

/*
 * Initialize RingHash, parameter replicas and algorithm
 */
func NewRingHash(replicas int, algorithm Algorithm) *RingHash {
	ringHash := &RingHash{
		algorithm: algorithm,
		replicas:  replicas,
		hashMap:   make(map[int]string),
	}
	if ringHash.algorithm == nil {
		ringHash.algorithm = crc32.ChecksumIEEE
	}
	return ringHash
}

/*
 * put some keys into the RingHash
 * for example, keys are similar to server ip
 */
func (ringHash *RingHash) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < ringHash.replicas; i++ {
			hashValue := int(ringHash.algorithm([]byte(strconv.Itoa(i) + key)))
			ringHash.sortedKeys = append(ringHash.sortedKeys, hashValue)
			ringHash.hashMap[hashValue] = key
		}
	}
	sort.Ints(ringHash.sortedKeys)
}

/*
 * get the closest items in the ringhash  by key
 * for example, get the closest server ip
 */
func (ringHash *RingHash) Get(key string) string {
	if len(ringHash.sortedKeys) == 0 {
		return ""
	}
	//得到实例key的hash值
	hashValue := int(ringHash.algorithm([]byte(key)))
	//选择大于实例key的hash值中最小的那一个
	index := sort.Search(len(ringHash.sortedKeys), func(i int) bool {
		return ringHash.sortedKeys[i] >= hashValue
	})

	if index >= len(ringHash.sortedKeys) {
		index = 0
	}

	return ringHash.hashMap[ringHash.sortedKeys[index]]
}
