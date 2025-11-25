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

// Function that returns a course by the id
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

// Function that returns all the courses
func GetCourses(c *gin.Context) {
	var coursesDto dto.CoursesDTO

	coursesDto, err := service.CourseService.GetCourses()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

// Function that returns courses based on search query and optional categories
func SearchCourses(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q")) // Puede venir vacío
	categoriesParam := c.Query("categories") // IDs separados por comas

	var categoryIDs []int
	if categoriesParam != "" {
		idsStr := strings.Split(categoriesParam, ",")
		for _, idStr := range idsStr {
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err == nil {
				categoryIDs = append(categoryIDs, id)
			}
		}
	}

	// Validación: al menos un filtro
	if query == "" && len(categoryIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debes proporcionar al menos 'q' o 'categories'"})
		return
	}

	courses, err := service.CourseService.SearchCourses(query, categoryIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// Function that insert a course
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

// Function that modified a course by the id
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

// Function that delete a course by the id
func DeleteCourseById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.CourseService.DeleteCourseById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}

// Function that returns all the comments about a course by the id
func GetCourseComments(c *gin.Context) {
	var commentsDto dto.CommentsDTO
	id, _ := strconv.Atoi(c.Param("id"))

	commentsDto, err := service.CourseService.GetCourseComments(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, commentsDto)
}
