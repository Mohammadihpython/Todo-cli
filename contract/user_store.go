package contract

import "todocli/entity"

type UserWriteStore interface {
	save(u entity.User)
}
type UserReadStore interface {
	Load() []entity.User
}
