package syllabus

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)


//	@Tags		Syllabus
// @Security	ApiKeyAuth
//	@Param		id			path	int				true	"id"
//	@Param		syllabus	body	models.Syllabus	true	"syllabus"
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus/topic/{id} [put]
func (sc *SyllabusController) UpdateTopic(c *gin.Context) {
	var syllabus models.Syllabus
	if err := c.ShouldBind(&syllabus); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signin",
				},
			},
		})
		return
	}
	syllabus.SyllabusID , _ = strconv.Atoi(c.Param("id"))
	err := sc.SyllabusRepository.UpdateTopic(c,  syllabus)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_UPDATE_SYLLABUS",
					Message: "Couldn't update syllabus",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: "Success"})
}