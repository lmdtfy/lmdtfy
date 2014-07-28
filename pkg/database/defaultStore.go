package database

import (
	r "github.com/dancannon/gorethink"

	"errors"
)

// Storer interface
type Storer interface {
	Store(interface{}) (string, error)
	Update(interface{}, string) error
	FindAll(interface{}) error
	FindByID(interface{}, string) error
}

// DefaultStore ...
type DefaultStore struct {
	TableName string
}

// NewDefaultStore ...
func NewDefaultStore(tableName string) DefaultStore {
	return DefaultStore{TableName: tableName}
}

// Store ...
func (d *DefaultStore) Store(v interface{}) (string, error) {

	res, err := r.Table(d.TableName).Insert(v).RunWrite(sess)
	if err != nil {
		return "", err
	}
	if len(res.GeneratedKeys) == 1 {
		return res.GeneratedKeys[0], nil
	}
	return "", nil
}

// Update ...
func (d *DefaultStore) Update(v interface{}, id string) error {

	res, err := r.Table(d.TableName).Get(id).Update(v).RunWrite(sess)
	if err != nil {
		return err
	}

	if res.Replaced == 0 {
		return ErrNotFound
	}
	return nil
}
func (d *DefaultStore) Filter(data interface{}, filter interface{}) error {
	res, err := r.Table(d.TableName).Filter(filter).Run(sess)
	if err != nil {
		return err
	}
	return res.All(data)
}

// FindAll ...
func (d DefaultStore) FindAll(data interface{}) error {

	res, err := r.Table(d.TableName).Run(sess)
	if err != nil {
		return err
	}
	return res.All(data)
}

// Delete ...
func (d DefaultStore) Delete(id string) error {
	if id == "" {
		return errors.New("ID required")
	}
	_, err := r.Table(d.TableName).Get(id).Delete().Run(sess)
	if err != nil {
		return err
	}
	return nil
}

// FindByID ...
func (d DefaultStore) FindByID(data interface{}, id string) error {

	res, err := r.Table(d.TableName).Get(id).Run(sess)
	if err != nil {
		return err
	}

	if res.IsNil() {
		return ErrNotFound
	}

	return res.One(data)
}
