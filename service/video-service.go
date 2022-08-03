package service

import "gin_chim/contents"

type VideoService interface {
	Save(contents.Video) contents.Video
	FindAll() []contents.Video
}

type videoService struct {
	videos []contents.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video contents.Video) contents.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []contents.Video {
	return service.videos
}
