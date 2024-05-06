package syllabus

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/nurzzaat/ZharasDiplom/internal/models"

	"log"

	"github.com/unidoc/unipdf/v3/common/license"
	

	//"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type SyllabusController struct {
	SyllabusRepository models.SyllabusRepository
}

func init() {
	err := license.SetMeteredKey(`5b78d4db382544315c5ae45eb3e5041431e199757514c99aed959b06db0935fa`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// @Tags		Syllabus
// @Security	ApiKeyAuth
// @Accept		json
// @Produce	json
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/syllabus/generate [post]
func (sc *SyllabusController) Generate(context *gin.Context) {
	//	@Param		id	path	int	true	"id"
	
	font, err := model.NewPdfFontFromTTFFile("times-new-roman-cyr.otf")
	
	if err != nil {
		log.Fatal(err)
	}
	
	fontBold, err := model.NewPdfFontFromTTFFile("times-new-roman-cyr.otf")
	if err != nil {
		log.Fatal(err)
	}

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)

	// Generate front page.
	drawFrontPage(c, font, fontBold)

	// Generate footer for pages.
	//drawFooter(c, font, fontBold)

	// Customize table of contents style.
	//customizeTOC(c, font, fontBold)

	// Generate basic usage chapter.
	// if err := basicUsage(c, font, fontBold); err != nil {
	// 	log.Fatal(err)
	// }

	// Generate styling content chapter.
	// if err := stylingContent(c, font, fontBold); err != nil {
	// 	log.Fatal(err)
	// }

	// Generate advanced usage chapter.
	// if err := advancedUsage(c, font, fontBold); err != nil {
	// 	log.Fatal(err)
	// }

	// Write to output file.
	if err := c.WriteToFile("unipdf-tables.pdf"); err != nil {
		log.Fatal(err)
	}

	context.JSON(200, gin.H{"message": "Success"})
}
func basicUsage(c *creator.Creator, font, fontBold *model.PdfFont) error {
	// Create chapter.
	ch := c.NewChapter("ФА")
	ch.SetMargins(0, 0, 50, 0)
	ch.GetHeading().SetFont(font)
	ch.GetHeading().SetFontSize(18)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Draw subchapters.
	contentAlignH(c, ch, font, fontBold)
	contentAlignV(c, ch, font, fontBold)
	contentWrapping(c, ch, font, fontBold)

	// Draw chapter.
	if err := c.Draw(ch); err != nil {
		return err
	}

	return nil
}

func contentAlignH(c *creator.Creator, ch *creator.Chapter, font, fontBold *model.PdfFont) {
	// Create subchapter.
	sc := ch.NewSubchapter("Content horizontal alignment")
	sc.SetMargins(0, 0, 30, 0)
	sc.GetHeading().SetFont(font)
	sc.GetHeading().SetFontSize(13)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Create subchapter description.
	desc := c.NewStyledParagraph()
	desc.SetMargins(0, 0, 10, 0)
	desc.Append("Cell content can be aligned horizontally left, right or it can be centered.")

	sc.Add(desc)

	// Create table.
	table := c.NewTable(3)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, align creator.CellHorizontalAlignment) {
		p := c.NewStyledParagraph()
		p.Append(text).Style.Font = font

		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(align)
		cell.SetContent(p)
	}

	// Draw table header.
	drawCell("Align left", fontBold, creator.CellHorizontalAlignmentLeft)
	drawCell("Align center", fontBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Align right", fontBold, creator.CellHorizontalAlignmentRight)

	// Draw table content.
	for i := 0; i < 5; i++ {
		num := i + 1

		drawCell(fmt.Sprintf("Product #%d", num), font, creator.CellHorizontalAlignmentLeft)
		drawCell(fmt.Sprintf("Description #%d", num), font, creator.CellHorizontalAlignmentCenter)
		drawCell(fmt.Sprintf("$%d", num*10), font, creator.CellHorizontalAlignmentRight)
	}

	sc.Add(table)
}

func contentAlignV(c *creator.Creator, ch *creator.Chapter, font, fontBold *model.PdfFont) {
	// Create subchapter.
	sc := ch.NewSubchapter("Content vertical alignment")
	sc.SetMargins(0, 0, 30, 0)
	sc.GetHeading().SetFont(font)
	sc.GetHeading().SetFontSize(13)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Create subchapter description.
	desc := c.NewStyledParagraph()
	desc.SetMargins(0, 0, 10, 0)
	desc.Append("Cell content can be positioned vertically at the top, bottom or in the middle of the cell.")

	sc.Add(desc)

	// Create table.
	table := c.NewTable(3)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, fontSize float64, align creator.CellVerticalAlignment) {
		p := c.NewStyledParagraph()
		chunk := p.Append(text)
		chunk.Style.Font = font
		chunk.Style.FontSize = fontSize

		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetVerticalAlignment(align)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetContent(p)
	}

	// Draw table header.
	drawCell("Align top", fontBold, 10, creator.CellVerticalAlignmentMiddle)
	drawCell("Align bottom", fontBold, 10, creator.CellVerticalAlignmentMiddle)
	drawCell("Align middle", fontBold, 10, creator.CellVerticalAlignmentMiddle)

	// Draw table content.
	for i := 0; i < 5; i++ {
		num := i + 1
		fontSize := float64(num) * 2

		drawCell(fmt.Sprintf("Product #%d", num), font, fontSize, creator.CellVerticalAlignmentTop)
		drawCell(fmt.Sprintf("$%d", num*10), font, fontSize, creator.CellVerticalAlignmentBottom)
		drawCell(fmt.Sprintf("Description #%d", num), font, fontSize, creator.CellVerticalAlignmentMiddle)
	}

	sc.Add(table)
}

func contentWrapping(c *creator.Creator, ch *creator.Chapter, font, fontBold *model.PdfFont) {
	// Create subchapter.
	sc := ch.NewSubchapter("Content wrapping")
	sc.SetMargins(0, 0, 30, 0)
	sc.GetHeading().SetFont(font)
	sc.GetHeading().SetFontSize(13)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Create subchapter description.
	desc := c.NewStyledParagraph()
	desc.SetMargins(0, 0, 10, 0)
	desc.Append("Cell text content is automatically broken into lines, depeding on the cell size.")

	sc.Add(desc)

	// Create table.
	table := c.NewTable(4)
	table.SetColumnWidths(0.25, 0.2, 0.25, 0.3)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, align creator.TextAlignment) {
		p := c.NewStyledParagraph()
		p.SetTextAlignment(align)
		p.SetMargins(2, 2, 0, 0)
		p.Append(text).Style.Font = font

		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetContent(p)
		cell.SetIndent(0)
	}

	// Draw table header.
	drawCell("Align left", fontBold, creator.TextAlignmentLeft)
	drawCell("Align center", fontBold, creator.TextAlignmentCenter)
	drawCell("Align right", fontBold, creator.TextAlignmentRight)
	drawCell("Align justify", fontBold, creator.TextAlignmentCenter)

	// Draw table content.
	content := "Maecenas tempor nibh gravida nunc laoreet, ut rhoncus justo ultricies. Mauris nec purus sit amet purus tincidunt efficitur tincidunt non dolor. Aenean nisl eros, volutpat vitae dictum id, facilisis ac felis. Integer lacinia, turpis at fringilla posuere, erat tortor ultrices orci, non tempor neque mauris ac neque. Morbi blandit ante et lacus ornare, ut vulputate massa dictum."

	drawCell(content, font, creator.TextAlignmentLeft)
	drawCell(content, font, creator.TextAlignmentCenter)
	drawCell(content, font, creator.TextAlignmentRight)
	drawCell(content, font, creator.TextAlignmentJustify)

	sc.Add(table)
}

func drawFrontPage(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.CreateFrontPage(func(args creator.FrontpageFunctionArgs) {
		p := c.NewStyledParagraph()
		p.SetMargins(0, 0, 300, 0)
		p.SetTextAlignment(creator.TextAlignmentCenter)

		chunk := p.Append("РаботаHello")
		chunk.Style.Font = font
		chunk.Style.FontSize = 56
		chunk.Style.Color = creator.ColorRGBFrom8bit(56, 68, 77)

		chunk = p.Append("\n")

		chunk = p.Append("Тудай")
		chunk.Style.Font = fontBold
		chunk.Style.FontSize = 40
		chunk.Style.Color = creator.ColorRGBFrom8bit(45, 148, 215)

		c.Draw(p)
	})
}

func (sc *SyllabusController) CreateSyllabuss(c *gin.Context) {
	//userID := c.GetUint("userID")

	//	@Security	ApiKeyAuth

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

	pdf := newReport()
	pdf = addHeader(pdf)

	subjectInfo := models.Header{}
	subjectInfo.SubjectName = `« Интеллектуализация образования, управления знаниями »`
	subjectInfo.SpecialityName = `7М06136  «Информационные системы»`

	pdf.SetFont("TNR_Bold", "", 11)
	pdf.CellFormat(200, 10, subjectInfo.SubjectName, "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, 10, subjectInfo.SpecialityName, "0", 0, "C", false, 0, "")
	pdf.Ln(-1)

	//pdf = SetSyllabusInfo(pdf, c)
	pdf = SetPreface(pdf, c)
	pdf = ConstructSyllabusTable(pdf, c)

	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully created"})
}

func newReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.AddUTF8Font("timesnrcyrmt", "", "timesnrcyrmt.ttf")
	pdf.AddUTF8Font("TNR_Bold", "", "TNR_Bold.ttf")

	return pdf
}
func addHeader(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.SetFont("TNR_Bold", "", 12)
	pdf.CellFormat(140, 10, "«Казахский университет технологии и бизнеса»", "1", 0, "L", false, 0, "")
	pdf.SetFont("timesnrcyrmt", "", 11)
	pdf.CellFormat(50, 10, "УМКД 17/ 02-10-2021", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.CellFormat(140, 10, "Учебно-методический комплекс дисциплины", "1", 0, "L", false, 0, "")
	pdf.CellFormat(50, 10, "2-редакция", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	pdf.SetY(80)
	pdf.SetFont("TNR_Bold", "", 14)
	pdf.CellFormat(200, 10, "РАБОЧАЯ УЧЕБНАЯ ПРОГРАММА ДИСЦИПЛИНЫ", "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(200, 10, "(СИЛЛАБУС)", "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	return pdf
}

// func SetSyllabusInfo(pdf *gofpdf.Fpdf, c *gin.Context) *gofpdf.Fpdf {
// 	info := models.SyllabusInfo{}
// 	info.FacultyName = `Технологический`
// 	info.KafedraName = `Информационные технологии`
// 	info.CourseNumber = "2"
// 	info.CreditNumber = "5"
// 	info.AllHours = "140"
// 	info.LectureHours = "30"
// 	info.PracticeLessons = "15"
// 	info.SRO = "105"

//		pdf.SetFont("timesnrcyrmt", "", 11)
//		pdf.Ln(20)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Факультет – `+info.FacultyName, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Кафедра – `+info.KafedraName, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Курс – `+info.CourseNumber, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Количество кредитов – `+info.CreditNumber, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Всего часов – `+info.AllHours, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Лекций – `+info.LectureHours, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `Семинарские (практические) занятия – `+info.PracticeLessons, "0", 0, "L", false, 0, "")
//		pdf.Ln(5)
//		pdf.SetX(25)
//		pdf.CellFormat(200, 10, `СРО – `+info.SRO, "0", 0, "L", false, 0, "")
//		pdf.Ln(80)
//		pdf.CellFormat(200, 10, `Астана  2024`, "0", 0, "C", false, 0, "")
//		return pdf
//	}
func SetPreface(pdf *gofpdf.Fpdf, c *gin.Context) *gofpdf.Fpdf {
	info := models.PrefaceInfo{}
	info.MadeBy.FullName = `Арман`
	info.MadeBy.Specialist = `препод`

	info.Discussion1 = `Главное, что эта игра добрая и веселая. Второго плана в книге, как бы и нет. И это достоинство. Авторское «я» весьма сильно`
	info.Discussed1.FullName = `Гулнур`
	info.Discussed1.Specialist = `зам.каф`

	info.Discussion2 = `Не эта строчка, сегодня звезда данного шедевра, а друзья. Вернее тема дружбы. Ну, или то, что это значит для других`
	info.Discussed2.FullName = `Гулсая`
	info.Discussed2.Specialist = `директор`

	info.ConfirmedBy.FullName = `Болат`
	info.ConfirmedBy.Specialist = `декан`

	pdf.AddPage()
	pdf = PageHeader(pdf, c)

	pdf.Ln(20)
	pdf.SetFont("TNR_Bold", "", 14)
	pdf.CellFormat(200, 10, `ПРЕДИСЛОВИЕ`, "0", 0, "C", false, 0, "")

	pdf.Ln(20)
	pdf.SetX(25)
	pdf.CellFormat(200, 10, `1. РАЗРАБОТАЛ`, "0", 0, "L", false, 0, "")
	pdf.Ln(5)
	pdf.SetX(25)
	pdf.SetFont("timesnrcyrmt", "", 12)
	pdf.CellFormat(200, 10, `Составитель: `+info.MadeBy.Specialist+`  ____________  `+info.MadeBy.FullName, "0", 0, "L", false, 0, "")

	pdf.Ln(10)
	pdf.SetX(25)
	pdf.SetFont("TNR_Bold", "", 14)
	pdf.CellFormat(200, 10, `2. ОБСУЖДЕНО`, "0", 0, "L", false, 0, "")
	pdf.Ln(10)
	pdf.SetX(25)
	pdf.SetFont("timesnrcyrmt", "", 12)
	pdf.MultiCell(180, 5, `2.1. `+info.Discussion1, "0", "L", false)
	pdf.Ln(5)
	pdf.SetX(25)
	pdf.CellFormat(200, 10, info.Discussed1.Specialist+`  ____________  `+info.Discussed1.FullName, "0", 0, "L", false, 0, "")

	pdf.Ln(10)
	pdf.SetX(25)
	pdf.MultiCell(180, 5, `2.2. `+info.Discussion2, "0", "L", false)
	pdf.Ln(5)
	pdf.SetX(25)
	pdf.CellFormat(200, 10, info.Discussed2.Specialist+`  ____________  `+info.Discussed2.FullName, "0", 0, "L", false, 0, "")

	pdf.Ln(15)
	pdf.SetX(25)
	pdf.SetFont("TNR_Bold", "", 14)
	pdf.CellFormat(200, 10, `2. УТВЕРЖДЕНО`, "0", 0, "L", false, 0, "")
	pdf.Ln(10)
	pdf.SetX(25)
	pdf.SetFont("timesnrcyrmt", "", 12)
	pdf.MultiCell(200, 10, info.ConfirmedBy.Specialist+`  ____________  `+info.ConfirmedBy.FullName, "0", "L", false)

	return pdf
}
func PageHeader(pdf *gofpdf.Fpdf, c *gin.Context) *gofpdf.Fpdf {
	pdf.Ln(50)
	pdf.SetX(25)
	pdf.CellFormat(60, 20, "Силлабус 17/02-10-2021", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 20, "Ред.№ __ от __ ______ 2021 г.", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 20, `стр. `+strconv.Itoa(pdf.PageNo())+` из `+strconv.Itoa(pdf.PageCount()), "1", 0, "C", false, 0, "")
	return pdf
}
func ConstructSyllabusTable(pdf *gofpdf.Fpdf, c *gin.Context) *gofpdf.Fpdf {
	pdf.AddPage()
	pdf.SetFont("TNR_Bold", "", 14)
	pdf.SetX(20)
	pdf.MultiCell(180, 5, `2. ТЕМАТИЧЕСКОЕ СОДЕРЖАНИЕ ДИСЦИПЛИНЫ И РАСПРЕДЕЛЕНИЕ ЧАСОВ ПО ВИДАМ ЗАНЯТИЙ. ССЫЛКИ НА ЛИТЕРАТУРУ`, "0", "C", false)

	pdf.Ln(15)
	pdf.SetX(20)
	pdf.CellFormat(7, 20, "№", "1", 0, "CT", false, 0, "")
	pdf.CellFormat(100, 20, "Модуль. Тема", "1", 0, "CT", false, 0, "")
	pdf.SetX(127)
	pdf.CellFormat(42, 15, "Количество часов", "1", 0, "CT", false, 0, "")
	pdf.SetY(50)
	pdf.SetX(127)
	pdf.CellFormat(14, 5, "ЛК", "1", 0, "CT", false, 0, "")
	pdf.CellFormat(14, 5, "СПЗ", "1", 0, "CT", false, 0, "")
	pdf.CellFormat(14, 5, "СРО", "1", 0, "CT", false, 0, "")
	pdf.SetY(35)
	pdf.SetX(169)
	pdf.CellFormat(30, 20, "Литература", "1", 0, "CT", false, 0, "")

	line := 65.0
	pdf.Ln(20)
	for i := 1; i <= 5; i++ {
		pdf.SetX(20)
		pdf.MultiCell(179, 5, `Mодуль `+strconv.Itoa(i)+`.  Информационно-коммуникационное пространство и его использование в образовании`, "1", "CT", false)
		pdf.SetTopMargin(0)
		for j := 0; j < 3; j++ {
			pdf.SetX(20)
			pdf.CellFormat(7, 10, strconv.Itoa(j+1), "1", 0, "CT", false, 0, "")
			pdf.MultiCell(100, 5, `Тема `+strconv.Itoa(j+1)+`. Социально-исторические предпосылки развития проблемы`, "1", "CT", false)
			pdf.SetY(line)
			line += 10
			pdf.SetX(127)
			pdf.CellFormat(14, 10, strconv.Itoa(int(line)), "1", 0, "CT", false, 0, "")
			pdf.CellFormat(14, 10, "", "1", 0, "CT", false, 0, "")
			pdf.CellFormat(14, 10, "10", "1", 0, "CT", false, 0, "")
			pdf.MultiCell(30, 5, "5.1.1 -5.1.7 5.1.1 -5.1.7", "1", "CT", false)
		}
		line += 10
	}

	return pdf
}
