package RESTHandlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/s21platform/gateway-service/internal/useCase/auth"
)

func AttachHandlers(r chi.Router) {
	fmt.Println("work attached")
	r.Route("/auth", func(authRouter chi.Router) {
		fmt.Println("Registered route")
		authRouter.Get("/login", auth.GetAuth)
	})
}
