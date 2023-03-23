package routers

import (
	"github.com/gofiber/fiber/v2"

	//----- SWAGGER -----
	swagger "github.com/arsmn/fiber-swagger/v2"
	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"

	"APITransactionGenerator/controller"
	_ "APITransactionGenerator/docs"
	"APITransactionGenerator/struct/publics"
)

func SetupPublicRoutes(app *fiber.App) {

	app.Get("/docs/*", swagger.HandlerDefault)

	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", publics.CheckServiceHealth)

	//DASHBOARD ENDPOINT
	dashboardEndpoint := v1Endpoint.Group("dashboard")
	// dashboardEndpoint.Get("/get_dashboardMenu", controller.Get_DashboardMenu)
	dashboardEndpoint.Post("/dashboardMenu", controller.DashboardMenu)

	//KPLUS ENDPOINT
	kPlusEndpoint := v1Endpoint.Group("kplus")
	kPlusEndpoint.Get("/kplus_upon_open", controller.KplusUponOpen)
	kPlusEndpoint.Get("/splash_screen", controller.SplashScreen)
	kPlusEndpoint.Get("/insti_param", controller.InstiParam)
	kPlusEndpoint.Get("/get_param", controller.GetParam)

	//JANUS TRANSACTION REPORT
	TransactionReport := v1Endpoint.Group("transaction")
	TransactionReport.Post("fetch_transaction", controller.TransactionCount)
	TransactionReport.Post("download_file", controller.GetPathFunc)

	//CREDENTIALS
	Credentials := v1Endpoint.Group("credentials")
	Credentials.Post("register_sign_up", controller.Registered)
	Credentials.Post("log_in", controller.Log_in)
}
