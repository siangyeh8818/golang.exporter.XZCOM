package tool

import (
	"github.com/patrickmn/go-cache"
)

type Mycache struct {
	seleiumcache *cache.Cache
}

func (c *Mycache) New(assigncache *cache.Cache) {

	c.seleiumcache = assigncache

}
