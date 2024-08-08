package events

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// EventManager gerencia handlers de eventos e despacha mensagens
type EventManager struct {
	handlers []EventHandler
}

// NewEventManager cria uma nova instância de EventManager
func NewEventManager() *EventManager {
	return &EventManager{handlers: []EventHandler{}}
}

// RegisterHandler registra um handler de eventos
func (m *EventManager) RegisterHandler(handler EventHandler) {
	m.handlers = append(m.handlers, handler)
}

// HandleMessage despacha uma mensagem para o handler apropriado
func (m *EventManager) HandleMessage(msg *message.Message) {
	var event Event
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		log.Printf("Erro ao desserializar mensagem: %v", err)
		msg.Nack() // Não confirmar o processamento da mensagem
		return
	}

	for _, handler := range m.handlers {
		if handler.CanHandle(event) {
			if err := handler.Handle(event); err != nil {
				log.Printf("Erro ao processar evento: %v", err)
				msg.Nack() // Não confirmar o processamento em caso de erro
				return
			}
			msg.Ack() // Confirmar o processamento da mensagem após o sucesso
			return
		}
	}

	log.Printf("Nenhum handler registrado pode processar o evento: %s", event.Event)
	msg.Ack() // Confirma o processamento se nenhum handler puder lidar com o evento
}
