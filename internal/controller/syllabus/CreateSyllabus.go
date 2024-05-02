package syllabus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"

	"github.com/unidoc/unioffice/document"
)

type SyllabusController struct {
	SyllabusRepository models.SyllabusRepository
}

// @Tags		Syllabus
// @Accept		json
// @Produce	json

// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/syllabus/generate [get]
func (sc *SyllabusController) CreateSyllabus(c *gin.Context) {
	//userID := c.GetUint("userID")
	
	// @Security	ApiKeyAuth
	
	// err := sc.SyllabusRepository.Create(c)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, models.ErrorResponse{
	// 		Result: []models.ErrorDetail{
	// 			{
	// 				Code:    "ERROR_GET_USER_PROFILE",
	// 				Message: "Can't get profile from db",
	// 				Metadata: models.Properties{
	// 					Properties1: err.Error(),
	// 				},
	// 			},
	// 		},
	// 	})
	// 	return
	// }

	doc := document.New()

	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText("This is a paragraph of text.")

	heading := doc.AddParagraph()
	heading.SetStyle("Heading1")
	run = heading.AddRun()
	run.AddText("This is a Heading 1")

	doc.SaveToFile("output.docx")

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully created"})
}
