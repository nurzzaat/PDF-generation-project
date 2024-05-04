package syllabus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)

//	@Tags		Syllabus
// @Security	ApiKeyAuth
//	@Param		subject	query	string	false	"subject"
//	@Accept		json
//	@Produce	json
//	@Success	200		{array}		models.Syllabus
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/syllabus/others [get]
func (sc *SyllabusController) GetAllOthers(c *gin.Context) {
	userID := c.GetUint("userID")
	subjectName := c.Query("subject")

	syllabuses , err := sc.SyllabusRepository.GetAllOthers(c , userID , subjectName)
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
	c.JSON(200 , syllabuses)
}
