package models

import "context"

type Syllabus struct {
	SyllabusID int          `json:"syllabusID"`
	MainInfo   SyllabusInfo `json:"mainInfo"`
	Preface    PrefaceInfo  `json:"preface"`
	Topics     []Modules    `json:"topics"`
	Literature Litrature    `json:"literature"`
}

type SyllabusInfo struct {
	SubjectInfo     Header `json:"subjectInfo"`
	FacultyName     string `json:"facultyName"`
	KafedraName     string `json:"kafedraName"`
	CourseNumber    int `json:"courseNumber"`
	CreditNumber    int `json:"creditNumber"`
	AllHours        int `json:"allHours"`
	LectureHours    int `json:"lectureHours"`
	PracticeLessons int `json:"practiceLessons"`
	SRO             int `json:"sro"`
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

type Litrature struct {
	MainLiterature       []string `json:"mainLiterature"`
	AdditionalLiterature []string `json:"additionalLiterature"`
}

type Header struct {
	SubjectName    string `json:"subjectName"`
	SpecialityName string `json:"specialityName"`
}
type Confirmer struct {
	FullName   string `json:"fullName"`
	Specialist string `json:"specialist"`
}

type SyllabusRepository interface {
	Create(c context.Context, syllabusInfo SyllabusInfo, userID uint)(int , error)
	Update(c context.Context, userID uint, syllabus Syllabus) error
	Delete(c context.Context, syllabusID int) error
	GetByID(c context.Context, syllabusID int, userID uint) (Syllabus, error)
	GetAllOwn(c context.Context, userID uint) ([]Syllabus, error)
	GetAllOthers(c context.Context, userID uint , subject string) ([]Syllabus, error)
}
