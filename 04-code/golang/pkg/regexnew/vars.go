package regexnew

import (
	"sync"
)

var (
	regexLock        = sync.Mutex{}
	lazyRegexLock    = sync.Mutex{}
	lazyRegexOnceMap = lazyRegexMap{
		items: make(
			map[string]*LazyRegex,
			DefaultCapacity),
	}

	New = newCreator{}
)
