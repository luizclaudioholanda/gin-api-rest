package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luizclaudioholanda/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()
	r.GET("/:nome", controllers.Saudacao)

	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.DELETE("/alunos/:id", controllers.RemoveAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.POST("/alunos", controllers.CriaAluno)

	r.Run(":8888")
}
