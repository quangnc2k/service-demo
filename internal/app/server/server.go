package server

import (
	"context"
	"git.cyradar.com/phinc/my-awesome-project/internal/config"
	"git.cyradar.com/phinc/my-awesome-project/internal/persistance"
	"git.cyradar.com/phinc/my-awesome-project/internal/persistance/mongo"
	"git.cyradar.com/phinc/my-awesome-project/internal/services"
	"log"
	"net/http"
	"sync"
	"time"

	"git.cyradar.com/phinc/my-awesome-project/pkg/hxxp"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func ServeServer(ctx context.Context, addr string) (err error) {
	ctx, cancel := context.WithCancel(ctx)

	config.Env = config.Init()
	persistance.DefaultRepository = mongo.NewMongoDB()
	services.DefaultDispatcher = services.NewHTTPDispatcher()

	r := chi.NewRouter()
	r.Use(hxxp.ChiRootContext(ctx))
	r.Route("/", func(r chi.Router) {
		r.Use(chiMiddleware.Recoverer)
		r.Use(chiMiddleware.URLFormat)
		r.Use(chiMiddleware.Timeout(30 * time.Second))
		r.Group(func(r chi.Router) {
			r.Route("/staff", func(r chi.Router) {
				r.Post("/", addStaff)
				r.Get("/{id}", viewStaff)
				r.Put("/{id}", updateStaff)
				r.Delete("/{id}", removeStaff)
			})
		})
	})

	errChan := make(chan error, 1)
	var exitErr error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) {
		defer wg.Done()
		srv := &http.Server{Addr: addr, IdleTimeout: 30 * time.Second, Handler: r}
		go func() {
			<-ctx.Done()
			srv.Shutdown(context.Background()) // nolint:errcheck
		}()
		log.Println("Listening on", addr)
		err := srv.ListenAndServe()
		select {
		case errChan <- err:
		default:
		}
	}(ctx, wg, errChan)

	select {
	case <-ctx.Done():
		exitErr = ctx.Err()
	case exitErr = <-errChan:
	}

	cancel()
	wg.Wait()
	return exitErr

}
