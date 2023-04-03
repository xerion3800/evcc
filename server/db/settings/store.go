package settings

import "github.com/xerion3800/evcc/api/store"

type storer struct {
	key string
}

var _ store.Store = (*storer)(nil)

func NewStore(key string) store.Store {
	return &storer{key: key}
}

func (s *storer) Load(res any) error {
	return Json(s.key, &res)
}

func (s *storer) Save(val any) error {
	return SetJson(s.key, val)
}
