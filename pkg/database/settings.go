package database

import (
	r "github.com/dancannon/gorethink"
	m "github.com/lmdtfy/lmdtfy/pkg/model"
)

type SettingStore struct {
	TableName string
}

func (ss *SettingStore) Get(settings *m.Settings) error {
	res, err := r.Table("settings").Run(sess)

	if err != nil {
		return err
	}
	return res.One(settings)
}
