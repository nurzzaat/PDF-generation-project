package syllabus

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/PDF-generation-project/internal/models"
	"github.com/nurzzaat/PDF-generation-project/pkg"

	"github.com/unidoc/unipdf/v3/common/license"

	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type SyllabusController struct {
	SyllabusRepository models.SyllabusRepository
	Env                *pkg.Env
}

var (
	pageCount = 5
)

func init() {
	err := license.SetMeteredKey(`90308c219a04bac91fbc3ef50f27c988fbc8f4438cd10214e02cfb132113da0e`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// @Tags		Syllabus
// @Accept		json
// @Param		id	path	int	true	"id"
// @Security	ApiKeyAuth
// @Produce	json
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/syllabus/generate/{id} [post]
func (sc *SyllabusController) Generate(context *gin.Context) {
	log.Println("Enter to function")

	userID := context.GetUint("userID")
	id, _ := strconv.Atoi(context.Param("id"))

	log.Println(id)

	syllabus, err := sc.SyllabusRepository.GetByID(context, id, userID)
	if err != nil {
		context.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_SYLLABUS",
					Message: "Couldn't get syllabus",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	log.Println(syllabus)

	// font, _ := model.NewCompositePdfFontFromTTFFile("/home/ubuntu/PDF-generation-project/timesnrcyrmt.ttf")
	// fontBold, _ := model.NewCompositePdfFontFromTTFFile("/home/ubuntu/PDF-generation-project/TNR_Bold.ttf")
	font, _ := model.NewCompositePdfFontFromTTFFile("timesnrcyrmt.ttf")
	fontBold, _ := model.NewCompositePdfFontFromTTFFile("TNR_Bold.ttf")

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)

	log.Println(1)
	FirstPage(c, font, fontBold, syllabus)
	log.Println(2)
	Preface(c, font, fontBold, syllabus)
	log.Println(3)
	Text(c, font, fontBold, syllabus)
	log.Println(7)
	Topic(c, font, fontBold, syllabus)
	log.Println(4)
	GradesTable(c, font, fontBold)
	log.Println(5)
	Literature(c, font, fontBold, syllabus)
	log.Println(6)
	log.Println(2)

	if err := c.WriteToFile(fmt.Sprintf("syllabus_%d.pdf", id)); err != nil {
		fmt.Println(err.Error())
	}
	context.JSON(200, gin.H{"message": fmt.Sprintf("syllabus_%d.pdf", id)})
}

func FirstPage(c *creator.Creator, font, fontBold *model.PdfFont, syllabus models.Syllabus) {
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

	p = c.NewParagraph(fmt.Sprintf("\n\n%s\n\n%s", syllabus.MainInfo.SubjectInfo.SubjectName, syllabus.MainInfo.SubjectInfo.SpecialityName))
	p.SetFontSize(12)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)

	c.Draw(p)

	p = c.NewParagraph(fmt.Sprintf(`Факультет – %s
	Кафедра – %v
	Курс – %v
	Количество кредитов – %v
	Всего часов – %v
	Лекций – %v
	Семинарские (практические) занятия – %v
	СРО – %v
	СРОП – %v`, syllabus.MainInfo.FacultyName, syllabus.MainInfo.KafedraName, syllabus.MainInfo.CourseNumber, syllabus.MainInfo.CreditNumber,
		syllabus.MainInfo.AllHours, syllabus.MainInfo.LectureHours, syllabus.MainInfo.PracticeLessons, syllabus.MainInfo.SRO, syllabus.MainInfo.SROP))
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
	p.SetMargins(0, 0, 180, 0)

	c.Draw(p)
}
func Preface(c *creator.Creator, font, fontBold *model.PdfFont, syllabus models.Syllabus) {
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
	chunk = p.Append(fmt.Sprintf("Составитель:  %s ___________ %s", syllabus.Preface.MadeBy.Specialist, syllabus.Preface.MadeBy.FullName))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	chapter1.Add(p)

	chapter2 := c.NewChapter("ОБСУЖДЕНО")
	heading = chapter2.GetHeading()
	heading.SetColor(creator.ColorBlack)
	heading.SetFontSize(16)
	heading.SetFont(fontBold)

	subChapter := chapter2.NewSubchapter(`На заседании кафедры "Информационные системы" от ` + syllabus.Preface.Discussion1)
	heading = subChapter.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(font)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 30, 0)
	chunk = p.Append(fmt.Sprintf("Заведующий кафедрой ___________ %s", syllabus.Preface.Discussed1.FullName))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter.Add(p)

	subChapter2 := chapter2.NewSubchapter(`На заседании комиссии по обеспечению качества Технологического факультета от ` + syllabus.Preface.Discussion2)
	heading = subChapter2.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(font)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 30, 0)
	chunk = p.Append(fmt.Sprintf("Председатель ТФ  ___________ %s", syllabus.Preface.Discussed2.FullName))
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
	chunk = p.Append(fmt.Sprintf("Декан факультета ТФ ___________ %s", syllabus.Preface.ConfirmedBy.FullName))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	chapter3.Add(p)

	chapter1.Add(chapter2)
	chapter1.Add(chapter3)

	c.Draw(chapter1)
}
func Text(c *creator.Creator, font, fontBold *model.PdfFont, syllabus models.Syllabus) {
	c.NewPage()

	chapter2 := c.NewChapter("ОБЩИЕ ПОЛОЖЕНИЯ")
	chapter2.SetShowNumbering(false)
	heading := chapter2.GetHeading()
	heading.SetColor(creator.ColorBlack)
	heading.SetFontSize(14)
	heading.SetFont(fontBold)

	subChapter := chapter2.NewSubchapter(`Общие сведения о преподавателе и дисциплине`)
	heading = subChapter.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p := c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk := p.Append(fmt.Sprintf(`Ф.И.О преподавателя: %s
	Ученая степень, звание, должность:  %s
	Факультет  %s
	Контактная информация: тел. %s, %s
	Сроки и время для консультации обучающихся: %s `, syllabus.Preface.MadeBy.FullName, syllabus.Preface.MadeBy.Specialist, syllabus.Preface.MadeBy.Faculty,
		syllabus.Preface.MadeBy.Email, syllabus.Preface.MadeBy.Address, syllabus.Preface.MadeBy.TimeForConsultation))
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter.Add(p)

	subChapter2 := chapter2.NewSubchapter(`Краткое описание содержания дисциплины`)
	heading = subChapter2.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Дисциплина формирует навыки по разработке и использованию интеллектуальных систем в образовании, учебно-методических материалов, инструкций по внедрению средств и технологий информатизации и интеллектуализации в систему вузовского образования, а также автоматизированной образовательной системы вуза. `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter2.Add(p)

	subChapter3 := chapter2.NewSubchapter(`Цель преподавания дисциплины`)
	heading = subChapter3.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Цель преподавания дисциплины изучить современные инновационные направления в науке, позволяющие разрабатывать и использовать наукоемкие, интеллектуальные системы в образовании. `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter3.Add(p)

	subChapter4 := chapter2.NewSubchapter(`Задача дисциплины`)
	heading = subChapter4.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Задачей дисциплины является приобретение и применение магистрантами теоретических знаний и практических знаний последних достижений информационно-коммуникационных технологий  в образовательном процессе в том числе и индустрии 4.`)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter4.Add(p)

	subChapter5 := chapter2.NewSubchapter(`Ожидаемые результаты обучения и формируемые компетенции.`)
	heading = subChapter5.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`При завершении освоения дисциплины обучающийся: - знает  современных инновационных направлениях в науке, позволяющие разрабатывать и использовать наукоемкие, интеллектуальные системы в образовании; - умеет применять учебно-методические, материалы, инструкции по внедрению средств и технологии информатизации и интеллектуализации в систему вузовского образования; - имеет навыки внедрения средств и технологии информатизации и интеллектуализации образования; - обладает следующими  компетенциями:  разработка и использование наукоемких, интеллектуальных систем в образовании. `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter5.Add(p)

	subChapter6 := chapter2.NewSubchapter(`Пререквизиты курса:`)
	heading = subChapter6.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Информатика. Информационные технологии  в образовании  `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter6.Add(p)

	subChapter7 := chapter2.NewSubchapter(`Постреквизиты курса:`)
	heading = subChapter7.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Использование полученных знаний при выполнении  дисертационной работы, в  дальнейшей профессиональной деятельности.  При участии в грантовых и научных проектах. `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter7.Add(p)

	subChapter8 := chapter2.NewSubchapter(`Формат обучения`)
	heading = subChapter8.GetHeading()
	heading.SetFontSize(12)
	heading.SetFont(fontBold)
	heading.SetMargins(0, 0, 15, 0)

	p = c.NewStyledParagraph()
	p.SetMargins(0, 0, 15, 0)
	chunk = p.Append(`Оф –лайн.`)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter8.Add(p)

	c.Draw(chapter2)
}

func Topic(c *creator.Creator, font, fontBold *model.PdfFont, syllabus models.Syllabus) {
	c.NewPage()

	headTable(c, font, fontBold, 4)

	p := c.NewParagraph("РАБОЧАЯ УЧЕБНАЯ ПРОГРАММА ДИСЦИПЛИНЫ\n\n(СИЛЛАБУС)")
	p.SetFontSize(14)
	p.SetFont(fontBold)
	p.SetTextAlignment(creator.TextAlignmentCenter)
	p.SetMargins(0, 0, 20, 20)

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
	lk := 0
	spz := 0
	sro := 0
	for key, module := range syllabus.Topics {
		cell = table.MultiColCell(25)
		p = c.NewParagraph(fmt.Sprintf("Модуль %d. %s", key+1, module.ModuleName))
		p.SetMargins(0, 0, 0, 7)
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetTextAlignment(creator.TextAlignmentCenter)
		cell.SetContent(p)
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		for key, topic := range module.Topics {
			cell = table.MultiColCell(1)
			p = c.NewParagraph(fmt.Sprintf("%d", key+1))
			p.SetMargins(0, 0, 0, 7)
			p.SetFont(font)
			p.SetFontSize(12)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(14)
			p = c.NewParagraph(fmt.Sprintf("Тема %d. %s", key+1, topic.TopicName))
			p.SetMargins(0, 0, 0, 7)
			p.SetFont(font)
			p.SetFontSize(12)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

			cell = table.MultiColCell(2)
			p = c.NewParagraph(fmt.Sprintf("%v", topic.LK))
			p.SetFont(font)
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			lk += topic.LK

			cell = table.MultiColCell(2)
			p = c.NewParagraph(fmt.Sprintf("%v", topic.SPZ))
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFont(font)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			spz += topic.SPZ

			cell = table.MultiColCell(2)
			p = c.NewParagraph(fmt.Sprintf("%v", topic.SRO))
			p.SetTextAlignment(creator.TextAlignmentCenter)
			p.SetFont(font)
			p.SetFontSize(12)
			p.SetMargins(0, 0, 0, 7)
			cell.SetContent(p)
			cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
			sro += topic.SRO

			cell = table.MultiColCell(4)
			p = c.NewParagraph(fmt.Sprintf("%v", topic.Literature))
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
	p = c.NewParagraph(fmt.Sprintf("%v", lk))
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	p = c.NewParagraph(fmt.Sprintf("%v", spz))
	p.SetFont(font)
	p.SetFontSize(12)
	p.SetMargins(0, 0, 0, 7)
	cell.SetContent(p)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)

	cell = table.MultiColCell(2)
	p = c.NewParagraph(fmt.Sprintf("%v", sro))
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

	chapter := c.NewChapter("")
	chapter.SetShowNumbering(false)

	subChapter := chapter.NewSubchapter("Задания  1 - 2 рубежного контролей знаний. ")
	heading := subChapter.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 10, 10)
	heading.SetFont(fontBold)

	pa := c.NewStyledParagraph()
	pa.SetMargins(0, 0, 10, 0)
	chunk := pa.Append(`Рубежный контроль представляет собой промежуточную форму оценки усвоения теоретических знаний и практических умений. Рубежный контроль имеет целью установить качество усвоения учебного материала по 5 модулям тем дисциплины. Методические укзания по выполнению практических занятий размещены  в файловых ресурсах на  Платонус) 
	Рубежный контроль проводится в письменной форме по теоретическим вопросам и отчетам  по практическим занятиям.`)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	subChapter.Add(pa)

	pa = c.NewStyledParagraph()
	pa.SetMargins(0, 0, 30, 0)
	pa.SetTextAlignment(creator.TextAlignmentCenter)
	chunk = pa.Append(`Вопросы к рубежному контролю №1 `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = fontBold
	subChapter.Add(pa)
	for key, question := range syllabus.Question1.Questions {
		pa = c.NewStyledParagraph()
		pa.SetMargins(0, 0, 10, 0)
		chunk = pa.Append(strconv.Itoa(key+1) + `. ` + question)
		chunk.Style.FontSize = 12
		chunk.Style.Font = font
		subChapter.Add(pa)
	}

	pa = c.NewStyledParagraph()
	pa.SetMargins(0, 0, 30, 0)
	pa.SetTextAlignment(creator.TextAlignmentCenter)
	chunk = pa.Append(`Вопросы к рубежному контролю №2 `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = fontBold
	subChapter.Add(pa)
	for key, question := range syllabus.Question2.Questions {
		pa = c.NewStyledParagraph()
		pa.SetMargins(0, 0, 10, 0)
		chunk = pa.Append(strconv.Itoa(key+1) + `. ` + question)
		chunk.Style.FontSize = 12
		chunk.Style.Font = font
		subChapter.Add(pa)
	}

	chapter3 := c.NewChapter("3. ПОЛИТИКА КУРСА")
	chapter3.SetShowNumbering(false)
	heading = chapter3.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 20, 0)
	heading.SetFont(fontBold)

	pa = c.NewStyledParagraph()
	pa.SetMargins(0, 0, 30, 0)
	chunk = pa.Append(`	Посещение занятий строго обязательно. Если по какой-либо причине, студент не может посещать занятия, то он несет ответственность за весь неосвоенный материал.  

	Контрольные задания обязательны для выполнения и должны сдаваться в установленные сроки. Работы, выполненные с опозданием, будут автоматически оцениваться ниже.

	Итоги рубежной аттестации проставляются с учетом посещаемости, выполнение самостоятельных работ студента, в установленные сроки, ответов на занятиях в устной или письменной форме, результатов самого рубежного контроля. 

	Если студент пропустил занятия и не смог сдать рубежный контроль в установленные сроки по болезни или другим уважительным причинам, документально подтвержденным, соответствующей организацией, он имеет право на индивидуальное прохождение рубежного контроля. В этом случае ему устанавливают индивидуальные сроки сдачи рубежного контроля согласно предоставленным документам. 

	Любое списывание или плагиат (использование, копирование готовых заданий и решений других студентов) будет пресекаться в виде исключения из аудитории и/или наказания оценкой «неудовлетворительно».

	Сотовые телефоны отключать во время проведения занятий.
	
			Обучающиеся обязаны: 
- неукоснительно соблюдать Правила академической честности при выполнении 
учебных заданий; 
- использовать достоверные и надёжные  источники информации; 
- качественно выполнять письменные работы, предусмотренные курсом, рефераты, 
курсовые, эссе, отчеты по практическим-лабораторным занятиям, на основе 
собственных идей при указании на авторство и идеи других людей; 
- самостоятельно выполнять все виды оцениваемых работ; 
- соблюдать Правила академической честности.  `)
	chunk.Style.FontSize = 12
	chunk.Style.Font = font
	chapter3.Add(pa)

	c.Draw(table)
	c.Draw(chapter)
	c.Draw(chapter3)
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

func Literature(c *creator.Creator, font, fontBold *model.PdfFont, syllabus models.Syllabus) {
	c.NewPage()

	chapter := c.NewChapter("4. ЛИТЕРАТУРА И ИНТЕРНЕТ-РЕСУРСЫ")
	chapter.SetShowNumbering(false)
	heading := chapter.GetHeading()
	heading.SetFontSize(14)
	heading.SetMargins(0, 0, 30, 10)
	heading.SetFont(fontBold)

	subChapter := chapter.NewSubchapter("4.1. Основная литература")
	subChapter.SetShowNumbering(false)
	heading = subChapter.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 10, 10)
	heading.SetFont(fontBold)

	for key, literature := range syllabus.Literature.MainLiterature {
		p := c.NewParagraph(fmt.Sprintf("4.1.%d. %s", key+1, literature))
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetMargins(0, 0, 0, 7)
		subChapter.Add(p)
	}

	subChapter = chapter.NewSubchapter("4.2. Дополнительная литература.  ")
	subChapter.SetShowNumbering(false)
	heading = subChapter.GetHeading()
	heading.SetFontSize(13)
	heading.SetMargins(0, 0, 10, 10)
	heading.SetFont(fontBold)

	for key, literature := range syllabus.Literature.AdditionalLiterature {
		p := c.NewParagraph(fmt.Sprintf("4.2.%d. %s", key+1, literature))
		p.SetFont(font)
		p.SetFontSize(12)
		p.SetMargins(0, 0, 0, 7)
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

func addText(c *creator.Creator, fontBold *model.PdfFont, text string) *creator.Paragraph {
	p := c.NewParagraph(text)
	p.SetMargins(0, 0, 0, 5)
	p.SetFont(fontBold)
	p.SetFontSize(12)
	p.SetTextAlignment(creator.TextAlignmentCenter)

	return p
}
