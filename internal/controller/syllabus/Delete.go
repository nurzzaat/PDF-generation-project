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
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus/{id} [delete]
func (sc *SyllabusController) Delete(c *gin.Context) {
	syllabusID , _ := strconv.Atoi(c.Param("id"))
	
	err := sc.SyllabusRepository.Delete(c , syllabusID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_DELETE_SYLLABUS",
					Message: "Couldn't delete syllabus",
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: "Success"})
}
