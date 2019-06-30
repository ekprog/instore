package instore

import (
	"errors"
	"fmt"
	"reflect"
)

type Settings struct {
	Postfix string
}

type Store struct {
	settings Settings
	items    map[string]reflect.Value
}

func NewStore(settings Settings) *Store {
	return &Store{
		settings: settings,
		items:    make(map[string]reflect.Value),
	}
}

func (s *Store) setItem(key string, i_v reflect.Value) {
	s.items[key+s.settings.Postfix] = i_v
}

func (s *Store) getItem(key string) reflect.Value {
	if v, ok := s.items[key+s.settings.Postfix]; ok {
		return v
	} else {
		return reflect.Value{}
	}
}

func (s *Store) itemExists(key string) bool {
	_, ok := s.items[key+s.settings.Postfix]
	return ok
}

func (s *Store) LoadItem(item interface{}) error {
	i_t := reflect.TypeOf(item)
	if i_t.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("store: item can be only a simple struct with config field: %v", i_t))
	}
	if s.itemExists(i_t.String()) {
		return errors.New(fmt.Sprintf("store: item '%v' is already exists", i_t))
	}
	s.setItem(i_t.String(), reflect.ValueOf(item))
	return nil
}

func (s *Store) UnloadItem(item interface{}) error {
	i_t := reflect.TypeOf(item)

	if i_t.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("store: Can unload only ptr config, not %v", i_t))
	}

	i_e := i_t.Elem()
	if i_e.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("store: Can unload only ptr config, not %v", i_t))
	}
	if !s.itemExists(i_e.String()) {
		return errors.New(fmt.Sprintf("store: item '%v' is not exists", i_t))
	}
	e_v := reflect.ValueOf(item).Elem()
	get_item := s.getItem(i_e.String())

	if !get_item.IsValid() {
		return errors.New(fmt.Sprintf("store: item '%v' cannot be resolve", i_t))
	}

	if !e_v.CanSet() {
		return errors.New(fmt.Sprintf("store: cannot fill struct %v", i_t))
	}

	e_v.Set(get_item)
	return nil
}
