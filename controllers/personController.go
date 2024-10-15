package controllers

import (
	"Mereb-V1/constants"
	"Mereb-V1/helpers"
	"Mereb-V1/models"
	"Mereb-V1/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePersonController(c *gin.Context) {

	var person models.Person

	//  1, Bind Request Body To Go Struct

	if err := c.ShouldBindJSON(&person); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"message": "Error In Binding Request Body Binding",

			"error": err.Error()})

		return
	}

	// 2 Validate The Go Struct and The Data we need by looking Model

	personValidator := helpers.NewValidatorService()

	validationError := personValidator.ValidateData(person)

	if validationError != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": " Invalide Data Format",
			"error":   validationError.Error()})
		return
	}

	// 3, Create The Person If There is No Isuue

	if err := services.CreatePersonService(person); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": constants.USER_ALREADY_EXIST,
			"error":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)
}



func GetAllPersonController(c *gin.Context) {
	persons := services.FetchAllPersonsService()
	c.JSON(http.StatusOK, persons)
}

func GetPersonController(c *gin.Context) {
	id := c.Param("id")
	person, err := services.FetchPersonService(id)
	// fmt.Println("Here In Get Person Cont : ", err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": constants.USER_NOT_EXIST,
			"error":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func DeletePersonController(c *gin.Context) {
	personID := c.Param("id")
	if _, err := services.DeletePersonService(personID); err != nil {
		fmt.Println("I Am Error : ", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": constants.USER_NOT_EXIST,

			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": constants.USER_DELETED})
}

func UpdatePersonController(c *gin.Context) {
	id := c.Param("id")
	var updatedPerson models.Person

	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check Validity

	// fmt.Println("ID", updatedPerson.ID)
	// fmt.Println("id :", id)

	// Ensure the ID matches the updated person's ID
	if updatedPerson.ID.String() != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UpdateId is not same as the actual id of user"})
		return
	}

	if err := services.UpdatePersonService(id, updatedPerson); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPerson)
}

func RouteNotFoundController(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Route not found"})
}
