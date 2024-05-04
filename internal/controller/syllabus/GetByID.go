package syllabus

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)

//	@Tags		Syllabus
// @Security	ApiKeyAuth
//	@Accept		json
//	@Param		id	path	int	true	"id"
//	@Produce	json
//	@Success	200		{object}	models.Syllabus
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus/{id} [get]
func (sc *SyllabusController) GetByID(c *gin.Context) {
	userID := c.GetUint("userID")

	syllabusID , _ := strconv.Atoi(c.Param("id"))
	syllabus , err := sc.SyllabusRepository.GetByID(c , syllabusID , userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_SYLLABUS",
					Message: "Couldn't get syllabus",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200 , syllabus)
}
