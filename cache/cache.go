package cache

import (
	"cloud-sek/globals"
	"cloud-sek/models"
)

func Init() {
	globals.Cache = &models.PostCache{
		PostCache: make(map[string]*models.Post),
	}
}
func SetPostCacheById(postId string, post models.Post) {
	globals.Cache.MuRW.Lock()
	globals.Cache.PostCache[postId] = &post
	globals.Cache.MuRW.Unlock()
}

func GetCouponCacheById(postId string) *models.Post {
	globals.Cache.MuRW.RLock()
	coupon, ok := globals.Cache.PostCache[postId]
	defer globals.Cache.MuRW.RUnlock()
	if !ok {
		return nil
	}
	return coupon
}
