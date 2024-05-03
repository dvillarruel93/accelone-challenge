package rest

import (
	"accelone-challenge/interfaces"
	"accelone-challenge/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactHandler struct {
	ContactUsecase interfaces.ContactUsecase
}

// NewContactHandler set endpoint URLs in echo framework and its handler functions.
func NewContactHandler(e *echo.Echo, contactUsecase interfaces.ContactUsecase) {
	handler := &ContactHandler{
		ContactUsecase: contactUsecase,
	}

	e.POST("/contacts", handler.AddContact)
	e.PUT("/contacts/:ID", handler.PutContact)
	e.GET("/contacts/:ID", handler.GetContact)
	e.DELETE("/contacts/:ID", handler.DeleteContact)
}

// AddContact saves a contact to DB
func (h *ContactHandler) AddContact(c echo.Context) error {
	// bind body from request
	contact := model.Contact{}
	if err := c.Bind(&contact); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error binding contact body")
	}

	// validate body with the contact struct that we want to save
	if err := model.Validate(contact); err != nil {
		return err
	}

	response, err := h.ContactUsecase.AddContact(contact)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}

// PutContact updates a contact to DB according to ID received in path parameter
func (h *ContactHandler) PutContact(c echo.Context) error {
	// bind body from request
	contact := model.Contact{}
	if err := c.Bind(&contact); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error binding contact body")
	}

	// validate body with the contact struct that we want to save
	if err := model.Validate(contact); err != nil {
		return err
	}

	// get ID param from path
	id := c.Param("ID")
	if len(id) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ID param can't be empty")
	}

	contact.ID = id
	response, err := h.ContactUsecase.PutContact(contact)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// GetContact returns the contact from DB according to ID received in path parameter
func (h *ContactHandler) GetContact(c echo.Context) error {
	// get ID param from path
	contactID := c.Param("ID")
	if len(contactID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ID param can't be empty")
	}

	response, err := h.ContactUsecase.GetContact(contactID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// DeleteContact removes the contact from DB according to ID received in path parameter
func (h *ContactHandler) DeleteContact(c echo.Context) error {
	// get ID param from path
	contactID := c.Param("ID")
	if len(contactID) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ID param can't be empty")
	}

	err := h.ContactUsecase.DeleteContact(contactID)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
