package service

import "music_player/model"

type InfoService struct {
}

func (s *InfoService) FindByKey(key string, key2 string) (model.ApiResult, error) {
	return model.Dba.FindByKey(key, key2)
}
