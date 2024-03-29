package server

import (
	galleryHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/gallery/handler"
	headlineHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/headline/handler"
	newsCurriculumHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_curriculum/handler"
	newsGeneralHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_general/handler"
	newsStudentCouncilHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_student_council/handler"
	newsStudentshipHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/news_studentship/handler"
	organizationalStructureHandler "github.com/radityacandra/sekolahku-smanssa-api/internal/application/organizational_structure/handler"
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
	v1.GET("/news/headline/:id", headlineHandler.Get)

	// news curriculum
	newsCurriculumHandler := newsCurriculumHandler.NewHandler(deps)
	v1.GET("/news/curriculum", newsCurriculumHandler.List)
	v1.GET("/news/curriculum/:id", newsCurriculumHandler.Get)

	// news general
	newsGeneralHandler := newsGeneralHandler.NewHandler(deps)
	v1.GET("/news/general", newsGeneralHandler.List)
	v1.GET("/news/general/:id", newsGeneralHandler.Get)

	// news student council
	newsStudentCouncilHandler := newsStudentCouncilHandler.NewHandler(deps)
	v1.GET("/news/student-council", newsStudentCouncilHandler.List)
	v1.GET("/news/student-council/:id", newsStudentCouncilHandler.Get)

	// news studentship
	newsStudentshipHandler := newsStudentshipHandler.NewHandler(deps)
	v1.GET("/news/studentship", newsStudentshipHandler.List)
	v1.GET("/news/studentship/:id", newsStudentshipHandler.Get)

	// gallery
	galleryHandler := galleryHandler.NewHandler(deps)
	v1.GET("/gallery", galleryHandler.List)

	// organizational structure
	organizationalStructure := organizationalStructureHandler.NewHandler(deps)
	v1.GET("/organizational_structure", organizationalStructure.List)
}
