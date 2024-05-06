package syllabus

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/models"

	"github.com/unidoc/unipdf/v3/common/license"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type SyllabusController struct {
	SyllabusRepository models.SyllabusRepository
}

var (
	pageCount = 5
)

func init() {
	err := license.SetMeteredKey(`49976580bfcb30b60793dc96151a167a16bfc370f88dc092042bd1cd2fa25929`)
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

	font, _ := model.NewCompositePdfFontFromTTFFile("timesnrcyrmt.ttf")
	fontBold, _ := model.NewCompositePdfFontFromTTFFile("TNR_Bold.ttf")

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)

	FirstPage(c, font, fontBold)
	Preface(c, font, fontBold)
	Topic(c, font, fontBold)

	// Generate basic usage chapter.
	// if err := basicUsage(c, font, fontBold); err != nil {
	// 	log.Fatal(err)
	// }

	if err := c.WriteToFile("unipdf-tables.pdf"); err != nil {
		log.Fatal(err)
	}
	context.JSON(200, gin.H{"message": "Success"})
}

func FirstPage(c *creator.Creator, font, fontBold *model.PdfFont) {
	table := c.NewTable(2)

	cell := table.NewCell()
	p := c.NewParagraph("«Казахский университет технологии и бизнеса»")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(fontBold)
	p.SetFontSize(12)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	p = c.NewParagraph("УМКД 17/ 02-10-2021")
	p.SetMargins(0, 0, 0, 7)
	p.SetMargins(5, 5, 5, 5)
	p.SetFontSize(12)
	p.SetFont(font)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	p = c.NewParagraph("Учебно-методический комплекс дисциплины")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	p = c.NewParagraph("2-редакция")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	c.Draw(table)

	p = c.NewParagraph("РАБОЧАЯ УЧЕБНАЯ ПРОГРАММА ДИСЦИПЛИНЫ\n\n(СИЛЛАБУС)")
	p.SetFontSize(14)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 120, 0)

	c.Draw(p)

	p = c.NewParagraph(fmt.Sprintf("\n\n%s\n\n%s", "« Интеллектуализация образования, управления знаниями»", "7М06136  «Информационные системы»"))
	p.SetFontSize(12)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)

	c.Draw(p)

	p = c.NewParagraph(fmt.Sprintf(`Факультет – %s
	Кафедра – %s
	Курс – %s
	Количество кредитов – %s
	Всего часов – %s
	Лекций – %s
	Семинарские (практические) занятия – %s
	СРО – %s`, "« Интеллектуализация образования, управления знаниями»", "Iнформационные технологии", "2", "3", "4", "5", "6", "7"))
	p.SetFontSize(12)
	p.SetLineHeight(1.5)
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentLeft)
	p.SetMargins(0, 0, 40, 0)

	c.Draw(p)

	p = c.NewParagraph("Астана 2024")
	p.SetFontSize(12)
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 200, 0)

	c.Draw(p)
}

func Preface(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.NewPage()
	headTable(c, font, fontBold, 2)
	chapterCenter := c.NewStyledParagraph()
	chapterCenter.SetMargins(0, 0, 20, 40)
	chapterCenter.SetTextAlignment(creator.TextAlignmentCenter)
	chapterCenter.SetTextVerticalAlignment(creator.TextVerticalAlignmentTop)
	chunk := chapterCenter.Append("ПРЕДИСЛОВИЕ")
	chunk.Style.FontSize = 16
	chunk.Style.Font = fontBold

	c.Draw(chapterCenter)

	chapter1 := c.NewChapter("РАЗРАБОТАЛ")
	heading := chapter1.GetHeading()
	heading.SetColor(creator.ColorBlack)
	heading.SetFontSize(16)
	heading.SetFont(fontBold)

	p := c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 15)
	chunk = p.Append(fmt.Sprintf("Составитель:  %s ___________ %s", " prepod", " Arman"))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	chapter1.Add(p)

	chapter2 := c.NewChapter("ОБСУЖДЕНО")
	heading = chapter2.GetHeading()
	heading.SetColor(creator.ColorBlack)
	heading.SetFontSize(16)
	heading.SetFont(fontBold)

	subChapter := chapter2.NewSubchapter("Главное, что эта игра добрая и веселая. Второго плана в книге, как бы и нет. И этo достоинство. Авторское «я» весьма сильно")
	heading = subChapter.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(font)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 30, 0)
	chunk = p.Append(fmt.Sprintf("%s ___________ %s", " prepod", " Arman"))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter.Add(p)

	subChapter2 := chapter2.NewSubchapter("Главное, что эта игра добрая и веселая. Второго плана в книге, как бы и нет. И этo достоинство. Авторское «я» весьма сильно")
	heading = subChapter2.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(font)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 30, 0)
	chunk = p.Append(fmt.Sprintf("%s ___________ %s", " prepod", " Arman"))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter2.Add(p)

	chapter3 := c.NewChapter("УТВЕРЖДЕНО")
	heading = chapter3.GetHeading()
	heading.SetColor(creator.ColorBlack)
	heading.SetFontSize(16)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 15)
	chunk = p.Append(fmt.Sprintf("%s ___________ %s", " prepod", " Arman"))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	chapter3.Add(p)

	chapter1.Add(chapter2)
	chapter1.Add(chapter3)

	c.Draw(chapter1)
}

func Topic(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.NewPage()

	p := c.NewParagraph("РАБОЧАЯ УЧЕБНАЯ ПРОГРАММА ДИСЦИПЛИНЫ\n\n(СИЛЛАБУС)")
	p.SetFontSize(14)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 0, 20)

	c.Draw(p)

	table := c.NewTable(25)
	table.EnableRowWrap(false)

	cell := table.MultiCell(2, 1)
	p = c.NewParagraph("№")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	p.SetFontSize(12)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiCell(2, 14)
	p = c.NewParagraph("Модуль. Тема")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	p.SetFontSize(12)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	subtable := c.NewTable(6)
	cell = subtable.MultiCell(1, 6)
	p = c.NewParagraph("Количество часов ")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("ЛК")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("СПЗ")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("СРО")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	
	table.AddSubtable(1, 16, subtable)

	subtable = c.NewTable(4)
	cell = subtable.MultiCell(2, 4)
	p = c.NewParagraph("Литература")
	cell.SetContent(p)
	p.SetFont(font)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)

	table.AddSubtable(1, 22, subtable)

	c.Draw(table)

}

func headTable(c *creator.Creator, font, fontBold *model.PdfFont, pageNum int) {
	table := c.NewTable(3)

	cell := table.NewCell()
	p := c.NewParagraph("Силлабус 17/02-10-2021")
	p.SetFont(font)
	p.SetMargins(0, 0, 15, 15)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	p = c.NewParagraph("Ред.№ __ от __ ______ 2021 г.")
	p.SetFontSize(12)
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 15, 15)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	p = c.NewParagraph(fmt.Sprintf("стр. %d из %d", pageNum, pageCount))
	p.SetFontSize(12)
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 10, 15)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	c.Draw(table)
}
func headTable1(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
		table := c.NewTable(2)
		table.SetMargins(20, 20, 20, 0)

		cell := table.NewCell()
		p := c.NewParagraph("«Казахский университет технологии и бизнеса»")
		p.SetFont(fontBold)
		p.SetFontSize(12)
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

		cell = table.NewCell()
		p = c.NewParagraph("УМКД 17/ 02-10-2021")
		p.SetFontSize(12)
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

		cell = table.NewCell()
		p = c.NewParagraph("Учебно-методический комплекс дисциплины")
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

		cell = table.NewCell()
		p = c.NewParagraph("2-редакция")
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

		subtable := c.NewTable(1)
		cell = subtable.MultiRowCell(2)
		p = c.NewParagraph("5")
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)

		table.AddSubtable(1, 3, subtable)

		block.Draw(table)
	})
}
