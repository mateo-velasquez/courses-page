package controller

import (
	"net/http"
	"project/dto"
	"project/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCourseById(c *gin.Context) {
	var courseDto dto.CourseDTO

	id, _ := strconv.Atoi(c.Param("id"))
	courseDto, err := service.CourseService.GetCourseById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func GetCourses(c *gin.Context) {
	var coursesDto dto.CoursesDTO

	coursesDto, err := service.CourseService.GetCourses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func SearchCourses(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))                    // It may come empty.
	categoriesParam := strings.TrimSpace(c.Query("categories")) // It may come empty.

	var categories []string
	if categoriesParam != "" {
		categories = strings.Split(categoriesParam, ",")
	}

	// Validación: al menos un filtro
	if query == "" && len(categories) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debes proporcionar al menos 'q' o 'categories'"})
		return
	}

	courses, err := service.CourseService.SearchCourses(query, categories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func InsertCourse(c *gin.Context) {
	var courseDto dto.CourseDTO

	err := c.BindJSON(&courseDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseDto, er := service.CourseService.InsertCourse(courseDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, courseDto) // Si está todo bien devuelve un 201 y el userDto
}

func PutCourseById(c *gin.Context) {
	var courseDto dto.CourseDTO
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&courseDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseDto.IDCourse = id

	courseDto, er := service.CourseService.PutCourseById(courseDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func DeleteCourseById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.CourseService.DeleteCourseById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
