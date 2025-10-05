package handler

import (
	"net/http"
	"ticket-app/internal/domain"
	"ticket-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type BuildingHandler struct {
	uc usecase.BuildingUsecase
}

func NewBuildingHandler(uc usecase.BuildingUsecase) *BuildingHandler { return &BuildingHandler{uc: uc} }

type createBuildingReq struct {
	BuildingName string   `json:"building_name"`
	Latitude     *float64 `json:"latitude,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty"`
}

type createBuildingRes struct {
	BuildingID uint64 `json:"building_id"`
}

type listBuildingRes struct {
	BuildingID   uint64  `json:"building_id"`
	BuildingName string  `json:"building_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

func (h *BuildingHandler) CreateBuilding(c echo.Context) error {
	var req createBuildingReq
	if err := c.Bind(&req); err != nil || req.BuildingName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	b := domain.Building{BuildingName: req.BuildingName, Latitude: req.Latitude, Longitude: req.Longitude}
	id, err := h.uc.CreateBuilding(c.Request().Context(), b)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createBuildingRes{BuildingID: id})
}
