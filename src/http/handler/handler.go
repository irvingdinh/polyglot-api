package handler

import "github.com/irvingdinh/polyglot-api/src/repository"

type Handler interface {
	MeHandler() MeHandler
	PingHandler() PingHandler
	UserHandler() UserHandler
}

func New(
	repositoryObj repository.Repository,
) Handler {
	return &handlerImpl{
		meHandler:   NewMeHandler(repositoryObj),
		pingHandler: NewPingHandler(),
		userHandler: NewUserHandler(repositoryObj),
	}
}

type handlerImpl struct {
	meHandler   MeHandler
	pingHandler PingHandler
	userHandler UserHandler
}

func (i *handlerImpl) MeHandler() MeHandler {
	return i.meHandler
}

func (i *handlerImpl) PingHandler() PingHandler {
	return i.pingHandler
}

func (i *handlerImpl) UserHandler() UserHandler {
	return i.userHandler
}
