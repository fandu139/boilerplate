package orders

import (
	"sync"

	"github.com/gin-gonic/gin"
	logics "github.com/ralali/event-api/src/business/v1/orders"
	dbEntity "github.com/ralali/event-api/src/entity/db/v1"
	httpEntity "github.com/ralali/event-api/src/entity/http/v1"
	repository "github.com/ralali/event-api/src/repository/db/v1"
)

// V1OrderService struct
type V1OrderService struct {
	tokenRepository    repository.TokenRepositoryInterface
	bookFreeEventLogic logics.BookFreeEventLogicInterface
	eventRepository    repository.EventRepositoryInterface
	ticketRepository   repository.TicketsRepositoryInterface
}

// OrderServiceHandler params
func OrderServiceHandler() *V1OrderService {
	return &V1OrderService{
		bookFreeEventLogic: logics.BookFreeEventLogicHandler(),
		tokenRepository:    repository.TokenRepositoryHandler(),
		eventRepository:    repository.EventRepositoryHandler(),
		ticketRepository:   repository.TicketsRepositoryHandler(),
	}
}

// OrderServiceInterface interface
type OrderServiceInterface interface {
	BookFreeEvent(context *gin.Context, payloads httpEntity.BookFreeEventRequest) error
	GetCheckoutOrder(eventID uint, context *gin.Context, waitGroup *sync.WaitGroup) (interface{}, error)
}

// BookFreeEvent paramas
// @context: gin.Context
// @payloads: BookFreeEventRequest
func (service *V1OrderService) BookFreeEvent(context *gin.Context, payloads httpEntity.BookFreeEventRequest) error {
	user, err := service.tokenRepository.GetSessionVisitor(context)
	if err != nil {
		return err
	}

	return service.bookFreeEventLogic.Do(user, payloads)
}

// GetCheckoutOrder params
func (service *V1OrderService) GetCheckoutOrder(eventID uint, context *gin.Context, waitGroup *sync.WaitGroup) (interface{}, error) {
	eventData := &dbEntity.Events{}
	ticketData := []dbEntity.Ticket{}
	waitGroup.Add(1)
	go service.eventRepository.GetEventByID(eventID, eventData, waitGroup)
	waitGroup.Wait()
	data, err := service.ticketRepository.GetTicketByEventID(eventID, ticketData)
	if err != nil {
		return data, err
	}
	return data, nil
}
