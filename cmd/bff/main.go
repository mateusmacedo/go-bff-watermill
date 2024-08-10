package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/mateusmacedo/bff-watermill/internal/slices/user"
	"github.com/mateusmacedo/bff-watermill/internal/slices/user/application"
	"github.com/mateusmacedo/bff-watermill/pkg/events"
	"github.com/mateusmacedo/bff-watermill/pkg/infrastructure"
)

func main() {
	config := infrastructure.LoadConfig()

	// Criação de um novo logger
	appLogger, err := infrastructure.NewZapAppLogger()
	if err != nil {
		panic(err)
	}

	// Configuração do adaptador de logger
	loggerAdapter := infrastructure.NewWatermillLoggerAdapter(appLogger)

	// Configuração do Redis
	redisClient := infrastructure.NewRedisClient()
	defer redisClient.Close()

	// Inicializar o publicador e assinante usando Redis Streams
	eventPublisher := infrastructure.NewWatermillRedisPublisher(redisClient, loggerAdapter)
	eventSubscriber := infrastructure.NewWatermillRedisSubscriber(redisClient, loggerAdapter)

	// Criar e registrar event handlers
	eventManager := events.NewEventManager()
	userCreatedHandler := application.NewUserCreatedHandler(appLogger)
	eventManager.RegisterHandler(userCreatedHandler)

	// Configurar o contexto com cancelamento
	ctx, cancel := context.WithCancel(context.Background())

	// Capturar sinais de término do sistema para cancelar o contexto
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigChan
		appLogger.Info(ctx, "Sinal capturado", map[string]interface{}{"signal": sig})
		cancel()
	}()

	// Registrar assinante antes de publicar delegando as mensagens para o eventManager realizar o roteamento
	go func() {
		messages, err := eventSubscriber.Subscribe(ctx, "user_events")
		if err != nil {
			log.Fatalf("Erro ao subscrever: %v", err)
		}

		for {
			select {
			case <-ctx.Done():
				appLogger.Info(ctx, "Encerrando assinante...", nil)
				return
			case msg := <-messages:
				eventManager.HandleMessage(ctx, msg)
			}
		}
	}()

	router := chi.NewRouter()

	// Inicializar o slice de usuário e registrar as rotas
	userSlice := user.NewUserSlice(eventPublisher, eventSubscriber)
	userSlice.RegisterRoutes(router)

	// Configurar e iniciar o servidor HTTP
	server := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: router,
	}

	// Goroutine para iniciar o servidor HTTP
	go func() {
		appLogger.Info(ctx, "Server starting on:"+config.ServerPort, nil)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
		appLogger.Info(ctx, "Server running on:"+config.ServerPort, nil)
	}()

	// Aguardar cancelamento e encerrar servidor HTTP
	<-ctx.Done()
	appLogger.Info(ctx, "Encerrando servidor...", nil)

	// Timeout para encerramento do servidor HTTP
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Erro ao encerrar servidor: %v", err)
	}

	appLogger.Info(ctx, "Servidor encerrado", nil)
}
