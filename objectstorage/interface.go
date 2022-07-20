package objectstorage

type ObjectStorageClient interface {
	GetObject(bucket, path string)
	PutObject(bucket, path string)
}
