package models

import "context"

type Syllabus struct {
	SyllabusID int          `json:"syllabusID"`
	MainInfo   SyllabusInfo `json:"mainInfo"`
	Preface    PrefaceInfo  `json:"preface"`
	Topics     []Modules    `json:"topics"`
	Text       Text         `json:"text"`
	Literature Litrature    `json:"literature"`
	Question1  Question     `json:"question1"`
	Question2  Question     `json:"question2"`
}

type SyllabusInfo struct {
	SubjectInfo     Header `json:"subjectInfo"`
	FacultyName     string `json:"facultyName"`
	KafedraName     string `json:"kafedraName"`
	CourseNumber    int    `json:"courseNumber"`
	CreditNumber    int    `json:"creditNumber"`
	AllHours        int    `json:"allHours"`
	LectureHours    int    `json:"lectureHours"`
	PracticeLessons int    `json:"practiceLessons"`
	SRO             string `json:"sro"`
	SROP            string `json:"srop"`
}

type PrefaceInfo struct {
	MadeBy      Confirmer `json:"madeBy"`
	Discussion1 string    `json:"discussion1"`
	Discussed1  Confirmer `json:"discussedBy1"`
	Discussion2 string    `json:"discussion2"`
	Discussed2  Confirmer `json:"discussedBy2"`
	ConfirmedBy Confirmer `json:"confirmedBy"`
}

type Modules struct {
	ModuleName string  `json:"moduleName"`
	Topics     []Topic `json:"topic"`
}

type Topic struct {
	TopicName  string `json:"topicName"`
	LK         int    `json:"LK"`
	SPZ        int    `json:"SPZ"`
	SRO        int    `json:"SRO"`
	Literature string `json:"literature"`
}

type Text struct {
	Text2 string `json:"text2"`
	Text3 string `json:"text3"`
	Text4 string `json:"text4"`
	Text5 string `json:"text5"`
	Text6 string `json:"text6"`
	Text7 string `json:"text7"`
	Text8 string `json:"text8"`
}

type SyllabusMaker struct {
	Faculty             string `json:"faculty"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	TimeForConsultation string `json:"consultation"`
}

type Question struct {
	Questions []string `json:"questions"`
}

type Litrature struct {
	MainLiterature       []string `json:"mainLiterature"`
	AdditionalLiterature []string `json:"additionalLiterature"`
}

type Header struct {
	SubjectName    string `json:"subjectName"`
	SpecialityName string `json:"specialityName"`
}
type Confirmer struct {
	FullName            string `json:"fullName"`
	Specialist          string `json:"specialist"`
	Faculty             string `json:"faculty"`
	Email               string `json:"email"`
	Address             string `json:"address"`
	TimeForConsultation string `json:"consultation"`
}

type SyllabusRepository interface {
	Create(c context.Context, syllabusInfo SyllabusInfo, userID uint) (int, error)
	UpdateMain(c context.Context, syllabus Syllabus) error
	UpdatePreface(c context.Context, syllabus Syllabus) error
	UpdateTopic(c context.Context, syllabus Syllabus) error
	UpdateText(c context.Context, syllabus Syllabus) error
	UpdateLiterature(c context.Context, syllabus Syllabus) error
	UpdateQuestion(c context.Context, syllabus Syllabus) error
	Delete(c context.Context, syllabusID int) error
	GetByID(c context.Context, syllabusID int, userID uint) (Syllabus, error)
	GetAllOwn(c context.Context, userID uint) ([]Syllabus, error)
	GetAllOthers(c context.Context, userID uint, subject string) ([]Syllabus, error)
}
