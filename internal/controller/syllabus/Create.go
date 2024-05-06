package syllabus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)

//	@Tags		Syllabus
// @Security	ApiKeyAuth
//	@Param		syllabus	body	models.SyllabusInfo	true	"syllabus"
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus [post]
func (sc *SyllabusController) Create(c *gin.Context) {
	userID := c.GetUint("userID")

	var syllabus models.SyllabusInfo
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

	id , err := sc.SyllabusRepository.Create(c, syllabus, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_SYLLABUS",
					Message: "Couldn't create syllabus",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(200, models.SuccessResponse{Result: id})
}
