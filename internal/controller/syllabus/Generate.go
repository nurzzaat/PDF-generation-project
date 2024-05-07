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
	GradesTable(c, font, fontBold)
	Literature(c, font, fontBold)

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

	headTable(c, font, fontBold, 3)

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
	p.SetTextAlignment(creator.TextAlignmentCenter)
	cell.SetContent(p)
	cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	subtable := c.NewTable(6)
	cell = subtable.MultiCell(1, 6)
	p = c.NewParagraph("Количество часов ")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("ЛК")
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("СПЗ")
	p.SetFont(font)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = subtable.MultiColCell(2)
	p = c.NewParagraph("СРО")
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	table.AddSubtable(1, 16, subtable)

	subtable = c.NewTable(4)
	cell = subtable.MultiCell(2, 4)
	p = c.NewParagraph("Литература")
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	cell.SetContent(p)
	p.SetFont(font)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)

	table.AddSubtable(1, 22, subtable)
	for o := 0; o < 20; o++ {
		table.NewCell()
	}
	for i := 1; i <= 5; i++ {
		cell = table.MultiColCell(25)
		p = c.NewParagraph(fmt.Sprintf("Модуль %d История и предпосылки  возникновения современной педагогики, как науки ", i))
		p.SetMargins(0, 0, 0, 7)
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetTextAlignment(creator.TextAlignmentCenter)
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		for j := 1; j <= 3; j++ {
			cell = table.MultiColCell(1)
			p = c.NewParagraph(fmt.Sprintf("%d", j))
			p.SetMargins(0, 0, 0, 7)
			p.SetFont(font)
			p.SetFontSize(12)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(14)
			p = c.NewParagraph("Тема 1. Социально-исторические предпосылки развития проблемы ")
			p.SetMargins(0, 0, 0, 7)
			p.SetFont(font)
			p.SetFontSize(12)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(2)
			p = c.NewParagraph("3")
			p.SetFont(font)
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(2)
			p = c.NewParagraph("")
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFont(font)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(2)
			p = c.NewParagraph("10")
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFont(font)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(4)
			p = c.NewParagraph("5.1.1 -5.1.7\n5.2.1-5.2.5 ")
			cell.SetContent(p)
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFont(font)
			p.SetFontSize(12)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)
		}
	}
	cell = table.MultiColCell(1)
	p = c.NewParagraph("")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	p.SetFontSize(12)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(14)
	p = c.NewParagraph("ВСЕГО: ")
	p.SetMargins(0, 0, 0, 7)
	p.SetFont(font)
	p.SetFontSize(12)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	p = c.NewParagraph(fmt.Sprintf("%v", 30))
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	p = c.NewParagraph(fmt.Sprintf("%v", 15))
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	p = c.NewParagraph(fmt.Sprintf("%v", 105))
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(4)
	p = c.NewParagraph("")
	cell.SetContent(p)
	p.SetFont(font)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)

	c.Draw(table)
}

func GradesTable(c *creator.Creator, font, fontBold *model.PdfFont) {

	p := c.NewParagraph("4. Оценка знаний обучающихся определяется по шкале")
	p.SetFontSize(13)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 30, 20)

	c.Draw(p)

	table := c.NewTable(6)

	cell := table.NewCell()
	cell.SetContent(addText(c, fontBold, "Оценка по буквенной системе "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, fontBold, "Цифровой эквивалент баллов "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, fontBold, "Процентное cодержание"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, fontBold, "Оценка по традиционной системе "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, fontBold, "Критерии оценивания"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "A"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "4.0"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "95-100"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "Отлично"))
	cell.SetVerticalAlignment(creator.CellVerticalAlignmentBottom)
	cell.SetBorder(creator.CellBorderSideTop, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Ответ на вопрос изложен полно, системно, соответствует теме задания "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "A-"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "3.67"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "90-94"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideBottom, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Отличная работа, в которой может быть допущена одна незначительная оценка "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "B+"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "3.33"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "85-89"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Содержание ответа в целом соответствует теме задания, встречаются несущественные ошибки "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "B"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "3.00"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "80-84"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Работа среднего уровня  с несколькими незначительными ошибками "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "B-"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "2.67"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "75-79 "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Обыкновенная работа с несколькими ошибками "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "C+"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "2.33"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "70-74"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "Хорошо"))
	cell.SetBorder(creator.CellBorderSideBottom, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Содержание ответа в целом соответствует  теме задания, но продемонстрировано удовлетворительное знание фактического материала; есть 1-2 ошибки в использовании и трактовке терминов; примеры, приведенные в ответе, не в полной мере соответствуют излагаемому материалу; встречаются 3-5 орфографических ошибок; работа выполнена не очень аккуратно, встречаются помарки и исправления. "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "C"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "2.00"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "65-69"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Ответ на вопрос недостаточно полный; недостаточное владение понятийно терминологическим аппаратом; отсутствуют примеры; встречаются стилистические ошибки и более 5 орфографических ошибок;  "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "C-"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "1.67"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "60-64"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "Удовлетворительно"))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Неполный ответ, имеются нарушения в логике и последовательности изложения материала; допущены грубые ошибки при определении сущности понятий и использовании терминов; отсутствуют выводы; "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "D+"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "1.33"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "55-59"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideLeft, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Средняя неплохая работа с существенными недостатками "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "D"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "1.00"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "50-54 "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, ""))
	cell.SetBorder(creator.CellBorderSideBottom, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Удовлетворительная работа, соответствующая минимальной положительной оценке "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	//
	cell = table.NewCell()
	cell.SetContent(addText(c, font, "F"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "0.00"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "0-49"))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.NewCell()
	cell.SetContent(addText(c, font, "Неудовлетворительно "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	cell.SetContent(addText(c, font, "Отсутствуют ответы по базовым вопросам дисциплины "))
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	c.Draw(table)
}

func Literature(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.NewPage()

	chapter := c.NewChapter("ЛИТЕРАТУРА И ИНТЕРНЕТ-РЕСУРСЫ")
	heading := chapter.GetHeading()
	heading.SetFontSize(14)
	heading.SetMargins(0, 0, 30, 10)
	heading.SetFont(fontBold)

	subChapter := chapter.NewSubchapter("Основная литература")
	heading = subChapter.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 10, 10)
	heading.SetFont(fontBold)

	for i := 1; i <= 5; i++ {
		p := c.NewParagraph(fmt.Sprintf("4.1.%d. %s", i, "Заёнчик В. М. Основы"))
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetMargins(0 ,0 ,0, 7)
		subChapter.Add(p)
	}

	subChapter = chapter.NewSubchapter("Дополнительная литература.  ")
	heading = subChapter.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 10, 10)
	heading.SetFont(fontBold)

	for i := 1; i <= 5; i++ {
		p := c.NewParagraph(fmt.Sprintf("4.2.%d. %s", i, "Заёнчик В. М. Основы"))
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetMargins(0 ,0 ,0, 7)
		subChapter.Add(p)
	}

	c.Draw(chapter)

}
func headTable(c *creator.Creator, font, fontBold *model.PdfFont, pageNum int) {
	table := c.NewTable(3)

	cell := table.NewCell()
	p := c.NewParagraph("Силлабус 17/02-10-2021")
	p.SetFont(font)
	p.SetMargins(0, 0, 15, 15)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	cell.SetContent(addText(c, fontBold, ""))
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

func addText(c *creator.Creator, fontBold *model.PdfFont, text string) *creator.Paragraph {
	p := c.NewParagraph(text)
	p.SetMargins(0, 0, 0, 5)
	p.SetFont(fontBold)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)

	return p
}
