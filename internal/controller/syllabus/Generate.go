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

	// Generate basic usage chapter.
	// if err := basicUsage(c, font, fontBold); err != nil {
	// 	log.Fatal(err)
	// }

	// Generate advanced usage chapter.
	if err := advancedUsage(c, font, fontBold); err != nil {
		log.Fatal(err)
	}

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
func Check(context *gin.Context) {
	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)

	headingFont, err := model.NewStandard14Font(model.HelveticaBoldName)
	if err != nil {
		log.Fatal(err)
	}

	if err := rowWrapDisabled(c, headingFont); err != nil {
		log.Fatal(err)
	}

	if err := c.WriteToFile("unipdf-tables.pdf"); err != nil {
		log.Fatal(err)
	}
	context.JSON(200, gin.H{"message": "Success"})
}

func headTable(c *creator.Creator, font, fontBold *model.PdfFont , pageNum int) {
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
	p = c.NewParagraph(fmt.Sprintf("стр. %d из %d",pageNum , pageCount ))
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

		// subtable := c.NewTable(1)
		// cell = subtable.MultiRowCell(2)
		// p = c.NewParagraph("5")
		// cell.SetContent(p)
		// cell.SetBorder(creator.CellBorderSideAll , creator.CellBorderStyleSingle , 1)
		// cell.SetVerticalAlignment(creator.CellVerticalAlignmentMiddle)

		// table.AddSubtable(1, 3 , subtable)

		block.Draw(table)
	})
}
func Preface(c *creator.Creator, font, fontBold *model.PdfFont) {
	c.NewPage()
	headTable(c, font, fontBold , 2 )
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

func basicUsage(c *creator.Creator, font, fontBold *model.PdfFont) error {
	// Create chapter.
	ch := c.NewChapter("РАБОЧАЯ УЧЕБНАЯ ПРОГРАММА ДИСЦИПЛИНЫ\n(СИЛЛАБУС)")
	ch.GetHeading().SetTextAlignment(creator.TextAlignmentCenter)
	ch.GetHeading().SetFont(fontBold)
	ch.GetHeading().SetFontSize(16)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	p := c.NewStyledParagraph()
	p.SetTextAlignment(creator.TextAlignmentCenter)

	chunk := p.Append("Астана 2024")
	chunk.Style.Font = font
	chunk.Style.FontSize = 12
	chunk.Style.Color = creator.ColorRGBFrom8bit(63, 68, 76)

	contentWrapping(c, ch, font, fontBold)

	// Draw chapter.
	if err := c.Draw(ch); err != nil {
		return err
	}

	return nil
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

func contentBackground(c *creator.Creator, ch *creator.Chapter, font, fontBold *model.PdfFont) {
	// Create subchapter.
	sc := ch.NewSubchapter("Cell background")
	sc.SetMargins(0, 0, 30, 0)
	sc.GetHeading().SetFont(font)
	sc.GetHeading().SetFontSize(13)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Create subchapter description.
	desc := c.NewStyledParagraph()
	desc.SetMargins(0, 0, 10, 0)
	desc.Append("The background color of the cells is also customizable.")

	sc.Add(desc)

	// Create table.
	table := c.NewTable(4)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, bgColor creator.Color) {
		p := c.NewStyledParagraph()
		p.SetMargins(2, 2, 0, 0)
		chunk := p.Append(text)
		chunk.Style.Font = font
		chunk.Style.Color = creator.ColorWhite

		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetBackgroundColor(bgColor)
		cell.SetContent(p)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetIndent(0)
	}

	// Draw table content.
	for i := 0; i < 15; i++ {
		drawCell("Content", fontBold, creator.ColorRGBFrom8bit(byte(i*20), byte(i*7), byte(i*4)))
		drawCell("Content", fontBold, creator.ColorRGBFrom8bit(byte(i*10), byte(i*20), byte(i*4)))
		drawCell("Content", fontBold, creator.ColorRGBFrom8bit(byte(i*15), byte(i*6), byte(i*9)))
		drawCell("Content", fontBold, creator.ColorRGBFrom8bit(byte(i*6), byte(i*7), byte(i*25)))
	}

	sc.Add(table)
}

func advancedUsage(c *creator.Creator, font, fontBold *model.PdfFont) error {
	c.NewPage()

	// Create chapter.
	ch := c.NewChapter("Advanced usage")
	ch.SetMargins(0, 0, 50, 0)
	ch.GetHeading().SetFont(font)
	ch.GetHeading().SetFontSize(18)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Draw subchapters.
	columnSpan(c, ch, font, fontBold)

	// Draw chapter.
	if err := c.Draw(ch); err != nil {
		return err
	}

	return nil
}

func columnSpan(c *creator.Creator, ch *creator.Chapter, font, fontBold *model.PdfFont) {
	// Create subchapter.
	sc := ch.NewSubchapter("Column span")
	sc.SetMargins(0, 0, 30, 0)
	sc.GetHeading().SetFont(font)
	sc.GetHeading().SetFontSize(13)
	sc.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	// Create subchapter description.
	desc := c.NewStyledParagraph()
	desc.SetMargins(0, 0, 10, 0)
	desc.Append("Table content can be configured to span a specified number of cells.")

	sc.Add(desc)

	// Create table.
	table := c.NewTable(5)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, colspan int, color, bgColor creator.Color) {
		p := c.NewStyledParagraph()
		p.SetMargins(2, 2, 0, 0)
		chunk := p.Append(text)
		chunk.Style.Font = font
		chunk.Style.Color = color

		cell := table.MultiColCell(colspan)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetBackgroundColor(bgColor)
		cell.SetContent(p)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentCenter)
		cell.SetIndent(0)
	}

	// Draw table content.

	// Colspan 1 + 1 + 1 + 1 + 1.
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)

	// Colspan 2 + 3.
	drawCell("2", fontBold, 2, creator.ColorWhite, creator.ColorRed)
	drawCell("3", fontBold, 3, creator.ColorWhite, creator.ColorRed)

	// Colspan 4 + 1.
	drawCell("4", fontBold, 4, creator.ColorBlack, creator.ColorGreen)
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorGreen)

	// Colspan 2 + 2 + 1.
	drawCell("2", fontBold, 2, creator.ColorBlack, creator.ColorYellow)
	drawCell("2", fontBold, 2, creator.ColorBlack, creator.ColorYellow)
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorYellow)

	// Colspan 5.
	drawCell("5", fontBold, 5, creator.ColorWhite, creator.ColorBlack)

	// Colspan 1 + 2 + 1 + 1.
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorYellow)
	drawCell("2", fontBold, 2, creator.ColorBlack, creator.ColorYellow)
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorYellow)
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorYellow)

	// Colspan 1 + 4.
	drawCell("1", fontBold, 1, creator.ColorBlack, creator.ColorGreen)
	drawCell("4", fontBold, 4, creator.ColorBlack, creator.ColorGreen)

	// Colspan 3 + 2.
	drawCell("3", fontBold, 3, creator.ColorWhite, creator.ColorRed)
	drawCell("2", fontBold, 2, creator.ColorWhite, creator.ColorRed)

	// Colspan 1 + 2 + 2.
	drawCell("1", fontBold, 2, creator.ColorBlack, creator.ColorYellow)
	drawCell("2", fontBold, 2, creator.ColorBlack, creator.ColorYellow)
	drawCell("2", fontBold, 1, creator.ColorBlack, creator.ColorYellow)

	// Colspan 1 + 1 + 1 + 2.
	drawCell("2", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("1", fontBold, 1, creator.ColorWhite, creator.ColorBlue)
	drawCell("2", fontBold, 2, creator.ColorWhite, creator.ColorBlue)

	sc.Add(table)
}

func rowWrapDisabled(c *creator.Creator, headingFont *model.PdfFont) error {
	heading := c.NewStyledParagraph()
	chunk := heading.Append("1. Table row wrap disabled")
	chunk.Style.Font = headingFont
	chunk.Style.FontSize = 20

	if err := c.Draw(heading); err != nil {
		return err
	}

	description := c.NewStyledParagraph()
	description.SetMargins(0, 0, 10, 20)
	chunk = description.Append("When table row wrapping is disabled, if one of the cells of a row does not fit in the available space of the current page, the whole row will be moved on the next one.")
	chunk.Style.FontSize = 14
	chunk.Style.Color = creator.ColorRGBFromHex("#777")

	if err := c.Draw(description); err != nil {
		return err
	}

	return fillTable(c, 22, false, func(table *creator.Table) error {
		sp1 := c.NewStyledParagraph()
		sp1.Append("This is a styled paragraph which will not fit on the current page. All its content should be moved on the next page, along with the entire row, as row wrapping is disabled.").Style.FontSize = 14

		p1 := c.NewParagraph("This is a regular paragraph which will fit on the current page. However, it will be moved on the next page.")
		p1.SetFontSize(14)

		sp2 := c.NewStyledParagraph()
		sp2.Append("This is a styled paragraph which will fit on the current page. However, it will be moved on the next page.").Style.FontSize = 14

		p2 := c.NewParagraph("This is a regular paragraph which will not fit on the current page. All its content should be moved on the next page, along with the entire row.")
		p2.SetFontSize(14)

		// Draw table row.
		for _, d := range []creator.VectorDrawable{sp1, p1, sp2, p2} {
			if err := drawCell(table, d); err != nil {
				return err
			}
		}

		return nil
	})
}

func drawCell(table *creator.Table, content creator.VectorDrawable) error {
	cell := table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	return cell.SetContent(content)
}

func fillTable(c *creator.Creator, lineCount int, enableRowWrap bool,
	addExtraRows func(table *creator.Table) error) error {
	// Create table.
	table := c.NewTable(4)
	table.SetMargins(0, 0, 10, 0)
	table.EnableRowWrap(enableRowWrap)

	for i := 0; i < lineCount; i++ {
		for j := 0; j < 4; j++ {
			sp := c.NewStyledParagraph()
			chunk := sp.Append(fmt.Sprintf("Row %d - Cell %d", i+1, j+1))
			chunk.Style.FontSize = 14

			if err := drawCell(table, sp); err != nil {
				return err
			}
		}
	}

	if err := addExtraRows(table); err != nil {
		return err
	}

	// Draw table.
	if err := c.Draw(table); err != nil {
		return err
	}

	return nil
}
