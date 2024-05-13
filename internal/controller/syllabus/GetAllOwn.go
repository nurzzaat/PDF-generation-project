package syllabus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/PDF-generation-project/internal/models"
)

//	@Tags		Syllabus
//
// @Security	ApiKeyAuth
//
//	@Accept		json
//	@Produce	json
//	@Success	200		{array}		models.Syllabus
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus [get]
func (sc *SyllabusController) GetAllOwn(c *gin.Context) {
	userID := c.GetUint("userID")

	syllabuses, err := sc.SyllabusRepository.GetAllOwn(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_SYLLABUSES",
					Message: "Couldn't get syllabuses",
				},
			},
		})
		return
	}
	c.JSON(200, syllabuses)
}
