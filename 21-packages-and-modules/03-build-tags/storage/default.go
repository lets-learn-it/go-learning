//go:build !az && !aws

package storage

func New() Store {
	panic("use az or aws implemention of Store.")
}
