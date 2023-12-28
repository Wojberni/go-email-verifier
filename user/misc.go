package user

import (
	"sync"
)

var (
	users = map[int]*User{}
	seq   = 1
	lock  = sync.Mutex{}
)
