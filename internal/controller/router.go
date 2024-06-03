package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/PDF-generation-project/middleware"
	"github.com/nurzzaat/PDF-generation-project/pkg"

	_ "github.com/nurzzaat/PDF-generation-project/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/nurzzaat/PDF-generation-project/internal/controller/auth"
	"github.com/nurzzaat/PDF-generation-project/internal/controller/syllabus"
	"github.com/nurzzaat/PDF-generation-project/internal/controller/user"
	"github.com/nurzzaat/PDF-generation-project/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	env := app.Env
	db := app.Pql

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/PDF-generation-project", "/home/ubuntu/PDF-generation-project")
	router.StaticFile("/PDF-generation-project", "/home/ubuntu/PDF-generation-project")


	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
		Env:            env,
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	syllabusController := &syllabus.SyllabusController{
		SyllabusRepository: repository.NewSyllabusRepository(db),
		Env:                env,
	}

	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)
	router.POST("/forgot-password", loginController.ForgotPassword)
	router.POST("/syllabus/generate/:id", syllabusController.Generate)

	router.Use(middleware.JWTAuth(env.AccessTokenSecret))
	router.POST("/logout", loginController.Logout)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
		userRouter.PUT("/profile", userController.UpdateProfile)
		userRouter.POST("/reset-password", loginController.ResetPassword)
	}

	syllabusRouter := router.Group("/syllabus")
	{
		syllabusRouter.POST("", syllabusController.Create)
		syllabusRouter.PUT("/main/:id", syllabusController.UpdateMain)
		syllabusRouter.PUT("/preface/:id", syllabusController.UpdatePreface)
		syllabusRouter.PUT("/topic/:id", syllabusController.UpdateTopic)
		syllabusRouter.PUT("/text/:id", syllabusController.UpdateText)
		syllabusRouter.PUT("/literature/:id", syllabusController.UpdateLiterature)
		syllabusRouter.PUT("/question/:id", syllabusController.UpdateQuestion)
		syllabusRouter.DELETE("/:id", syllabusController.Delete)
		syllabusRouter.GET("/:id", syllabusController.GetByID)
		syllabusRouter.GET("", syllabusController.GetAllOwn)
		syllabusRouter.GET("/others", syllabusController.GetAllOthers)
		//syllabusRouter.POST("/generate/:id", syllabusController.Generate)
	}
}
