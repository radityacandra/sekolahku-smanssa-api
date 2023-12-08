package server

import (
	headlineHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/handler"
	newsCurriculumHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/handler"
	newsGeneralHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/handler"
	newsStudentCouncilHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/handler"
	newsStudentshipHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/handler"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/application/user/handler"
	"github.com/radityacandra/sekolahku-smanssa-api/internal/server/types"
)

func InitRoutes(deps *types.Dependency) {
	v1 := deps.Echo.Group("/v1")

	// auth login group
	userHandler := handler.NewHandler(deps)
	v1.POST("/auth/login", userHandler.Auth)

	// headline
	headlineHandler := headlineHandler.NewHandler(deps)
	v1.GET("/news/headline", headlineHandler.List)

	// news curriculum
	newsCurriculumHandler := newsCurriculumHandler.NewHandler(deps)
	v1.GET("/news/curriculum", newsCurriculumHandler.List)

	// news general
	newsGeneralHandler := newsGeneralHandler.NewHandler(deps)
	v1.GET("/news/general", newsGeneralHandler.List)

	// news student council
	newsStudentCouncilHandler := newsStudentCouncilHandler.NewHandler(deps)
	v1.GET("/news/student-council", newsStudentCouncilHandler.List)

	// news studentship
	newsStudentshipHandler := newsStudentshipHandler.NewHandler(deps)
	v1.GET("/news/studentship", newsStudentshipHandler.List)
}
