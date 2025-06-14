package api

import (
	"fmt"
	"net/http"

	"github.com/daniel-pantoja/api-estudante/db"
	"github.com/labstack/echo/v4"
)

func (api *API) getEstudantes(c echo.Context) error {
	estudantes, err := api.DB.GetEstudantes()
	if err != nil {
		return c.String(http.StatusNotFound, "Falha em obter os estudantes")
	}

	return c.JSON(http.StatusOK, estudantes)
}

func (api *API) createEstudante(c echo.Context) error {
	estudante := db.Estudante{}
	if err := c.Bind(&estudante); err != nil {
		return err
	}

	if err := api.DB.AddEstudante(estudante); err != nil {
		return c.String(http.StatusInternalServerError, "Erro para cadastrar estudante")
	}

	return c.String(http.StatusOK, "Estudante cadastrado")
}

func (api *API) getEstudante(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Pega %s estudante", id)
	return c.String(http.StatusOK, getStud)
}

func (api *API) updateEstudante(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Atualiza %s estudante", id)
	return c.String(http.StatusOK, updateStud)
}

func (api *API) deleteEstudante(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Deleta %s estudante", id)
	return c.String(http.StatusOK, deleteStud)
}