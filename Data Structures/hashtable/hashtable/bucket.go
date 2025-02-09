package hashtable

type BucketNode[K comparable, V any] struct {
	key K
	value V
	next *BucketNode[K, V]
}
	

type Bucket[K comparable, V any] struct {
	head *BucketNode[K, V]	
}