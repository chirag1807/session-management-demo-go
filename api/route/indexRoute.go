package route

import (
	"sessionmanagement/api/controller"
	"sessionmanagement/api/middleware"
	"sessionmanagement/api/repository"
	"sessionmanagement/api/service"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func UsersRoutes(conn *pgx.Conn, sessionManager *scs.SessionManager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(sessionManager.LoadAndSave)
	//This middleware takes care of loading and committing session data to the session store.

	//By default SCS uses an in-memory store for session data. This is convenient (no setup!) and very fast,
	//but all session data will be lost when your application is stopped or restarted.
	//Therefore it's useful for applications where data loss is an acceptable trade off for fast performance, or for prototyping and testing purposes.
	//In most production applications you will want to use a persistent session store like PostgreSQL or MySQL instead.

	//to use different session store (redis, postgres etc...) go to official github page of scs.
	//https://github.com/alexedwards/scs?tab=readme-ov-file#configuring-the-session-store

	authRepository := repository.NewAuthRepo(conn)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService, sessionManager)

	articleRepository := repository.NewArticleRepo(conn)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService, sessionManager)

	topicRepository := repository.NewTopicRepo(conn)
	topicService := service.NewTopicService(topicRepository)
	topicController := controller.NewTopicController(topicService, sessionManager)

	r.Route("/api/auth", func(r chi.Router) {

		r.Post("/login", authController.UserLogin)
	})

	r.Route("/api/article", func(r chi.Router) {
		r.Use(middleware.VerifyidentityAndSession(0, sessionManager))

		r.Post("/add-article", articleController.AddArticle)

		r.Get("/get-my-articles", articleController.GetMyArticles)

		r.Get("/get-article-by-id/{ID}", articleController.GetArticleById)

		r.Put("/update-article", articleController.UpdateArticle)

		r.Delete("/delete-article/{ID}", articleController.DeleteArticle)
	})

	r.Route("/api/admin", func(r chi.Router) {
		r.Use(middleware.VerifyidentityAndSession(1, sessionManager))

		r.Post("/add-topic", topicController.AddTopic)

		r.Get("/get-all-topics", topicController.GetAllTopics)

		r.Put("/update-topic", topicController.UpdateTopic)

		r.Delete("/delete-topic/{ID}", topicController.DeleteTopic)
	})

	return r
}
