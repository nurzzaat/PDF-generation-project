package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/PDF-generation-project/internal/models"
)

type SyllabusRepository struct {
	db *pgxpool.Pool
}

func NewSyllabusRepository(db *pgxpool.Pool) models.SyllabusRepository {
	return &SyllabusRepository{db: db}
}

func (sr *SyllabusRepository) Create(c context.Context, syllabusInfo models.SyllabusInfo, userID uint) (int, error) {
	var id int
	query := `INSERT INTO syllabus(
		userid, subject, faculty, kafedra, specialist, coursenumber, creditnumber, allhours, lecturehour, practicehour, sro , srop)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id;`
	err := sr.db.QueryRow(c, query, userID, syllabusInfo.SubjectInfo.SubjectName,
		syllabusInfo.FacultyName, syllabusInfo.KafedraName, syllabusInfo.SubjectInfo.SpecialityName, syllabusInfo.CourseNumber,
		syllabusInfo.CreditNumber, syllabusInfo.AllHours, syllabusInfo.LectureHours, syllabusInfo.PracticeLessons, syllabusInfo.SRO, syllabusInfo.SROP).Scan(&id)
	if err != nil {
		return id, err
	}
	additionQuery := `INSERT INTO public.addition(syllabusid) VALUES ($1);`
	_, err = sr.db.Exec(c, additionQuery, id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (sr *SyllabusRepository) UpdateMain(c context.Context, syllabus models.Syllabus) error {
	query := `UPDATE syllabus
	SET subject=$1, faculty=$2, kafedra=$3, specialist=$4, coursenumber=$5, creditnumber=$6, allhours=$7, 
	lecturehour=$8, practicehour=$9, sro=$10 , srop=$11 WHERE id = $12`
	_, err := sr.db.Exec(c, query, syllabus.MainInfo.SubjectInfo.SubjectName, syllabus.MainInfo.FacultyName,
		syllabus.MainInfo.KafedraName, syllabus.MainInfo.SubjectInfo.SpecialityName, syllabus.MainInfo.CourseNumber,
		syllabus.MainInfo.CreditNumber, syllabus.MainInfo.AllHours, syllabus.MainInfo.LectureHours,
		syllabus.MainInfo.PracticeLessons, syllabus.MainInfo.SRO, syllabus.MainInfo.SROP, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SyllabusRepository) UpdatePreface(c context.Context, syllabus models.Syllabus) error {
	query := `UPDATE syllabus
	SET madeby=$1, madebymajor=$2, discuss1=$3, discussby1=$4, discussby1major=$5, discuss2=$6, discussby2=$7, 
	discussby2major=$8, confirmedby=$9, confirmedbymajor=$10 , facultyOfProf = $11 , email=$12 , address=$13 , timeofcons=$14 WHERE id = $15`
	_, err := sr.db.Exec(c, query, syllabus.Preface.MadeBy.FullName, syllabus.Preface.MadeBy.Specialist,
		syllabus.Preface.Discussion1, syllabus.Preface.Discussed1.FullName, syllabus.Preface.Discussed1.Specialist,
		syllabus.Preface.Discussion2, syllabus.Preface.Discussed2.FullName, syllabus.Preface.Discussed2.Specialist,
		syllabus.Preface.ConfirmedBy.FullName, syllabus.Preface.ConfirmedBy.Specialist, syllabus.Preface.MadeBy.Faculty,
		syllabus.Preface.MadeBy.Email, syllabus.Preface.MadeBy.Address, syllabus.Preface.MadeBy.TimeForConsultation, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SyllabusRepository) UpdateTopic(c context.Context, syllabus models.Syllabus) error {
	deleteQuery := `delete from modules where syllabusid = $1`
	_, err := sr.db.Exec(c, deleteQuery, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	for key, module := range syllabus.Topics {
		var moduleID int
		moduleQuery := `INSERT INTO modules(
			syllabusid, orderid, title)
			VALUES ($1, $2, $3) returning id;`
		err = sr.db.QueryRow(c, moduleQuery, syllabus.SyllabusID, key+1, module.ModuleName).Scan(&moduleID)
		if err != nil {
			return err
		}
		for _, topic := range module.Topics {
			topicQuery := `INSERT INTO topic(
				moduleid, title, lk, spz, sro, literature)
				VALUES ($1, $2, $3, $4, $5, $6);`
			_, err := sr.db.Exec(c, topicQuery, moduleID, topic.TopicName, topic.LK, topic.SPZ, topic.SRO, topic.Literature)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (sr *SyllabusRepository) UpdateText(c context.Context, syllabus models.Syllabus) error {
	query := `UPDATE public.addition
	SET text2=$1, text3=$2, text4=$3, text5=$4, text6=$5, text7=$6, text8=$7
	WHERE syllabusid=$8`
	_, err := sr.db.Exec(c, query, syllabus.Text.Text2, syllabus.Text.Text3, syllabus.Text.Text4, syllabus.Text.Text5,
		syllabus.Text.Text6, syllabus.Text.Text7, syllabus.Text.Text8, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SyllabusRepository) UpdateQuestion(c context.Context, syllabus models.Syllabus) error {
	deleteQuery := `delete from questions where syllabusid = $1`
	_, err := sr.db.Exec(c, deleteQuery, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	for _, question := range syllabus.Question1.Questions {
		var questionID int
		questionQuery := `INSERT INTO questions(
			syllabusid, title , sequen)
			VALUES ($1, $2 , 1) returning id;`
		err = sr.db.QueryRow(c, questionQuery, syllabus.SyllabusID, question).Scan(&questionID)
		if err != nil {
			return err
		}
	}
	for _, question := range syllabus.Question2.Questions {
		var questionID int
		questionQuery := `INSERT INTO questions(
			syllabusid, title , sequen)
			VALUES ($1, $2 , 2) returning id;`
		err = sr.db.QueryRow(c, questionQuery, syllabus.SyllabusID, question).Scan(&questionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sr *SyllabusRepository) UpdateLiterature(c context.Context, syllabus models.Syllabus) error {
	deleteQuery := `delete from literature where syllabusid = $1`
	_, err := sr.db.Exec(c, deleteQuery, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	for _, main := range syllabus.Literature.MainLiterature {
		literatureQuery := `INSERT INTO literature(
			syllabusid, type, title)
			VALUES ($1, $2, $3);`
		_, err = sr.db.Exec(c, literatureQuery, syllabus.SyllabusID, "main", main)
		if err != nil {
			return err
		}
	}
	for _, additional := range syllabus.Literature.AdditionalLiterature {
		literatureQuery := `INSERT INTO literature(
			syllabusid, type, title)
			VALUES ($1, $2, $3);`
		_, err = sr.db.Exec(c, literatureQuery, syllabus.SyllabusID, "additional", additional)
		if err != nil {
			return err
		}
	}
	return nil
}
func (sr *SyllabusRepository) Update(c context.Context, syllabus models.Syllabus) error {
	query := `UPDATE syllabus
	SET subject=$1, faculty=$2, kafedra=$3, specialist=$4, coursenumber=$5, creditnumber=$6, allhours=$7, 
	lecturehour=$8, practicehour=$9, sro=$10, madeby=$11, madebymajor=$12, discuss1=$13, discussby1=$14, 
	discussby1major=$15, discuss2=$16, discussby2=$17, discussby2major=$18, confirmedby=$19, confirmedbymajor=$20 , srop=$21
	WHERE id = $22`
	_, err := sr.db.Exec(c, query, syllabus.MainInfo.SubjectInfo.SubjectName, syllabus.MainInfo.FacultyName,
		syllabus.MainInfo.KafedraName, syllabus.MainInfo.SubjectInfo.SpecialityName, syllabus.MainInfo.CourseNumber,
		syllabus.MainInfo.CreditNumber, syllabus.MainInfo.AllHours, syllabus.MainInfo.LectureHours,
		syllabus.MainInfo.PracticeLessons, syllabus.MainInfo.SRO, syllabus.Preface.MadeBy.FullName, syllabus.Preface.MadeBy.Specialist,
		syllabus.Preface.Discussion1, syllabus.Preface.Discussed1.FullName, syllabus.Preface.Discussed1.Specialist,
		syllabus.Preface.Discussion2, syllabus.Preface.Discussed2.FullName, syllabus.Preface.Discussed2.Specialist,
		syllabus.Preface.ConfirmedBy.FullName, syllabus.Preface.ConfirmedBy.Specialist, syllabus.MainInfo.SROP, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	deleteQuery := `delete from modules where syllabusid = $1`
	_, err = sr.db.Exec(c, deleteQuery, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	for key, module := range syllabus.Topics {
		var moduleID int
		moduleQuery := `INSERT INTO modules(
			syllabusid, orderid, title)
			VALUES ($1, $2, $3) returning id;`
		err = sr.db.QueryRow(c, moduleQuery, syllabus.SyllabusID, key+1, module.ModuleName).Scan(&moduleID)
		if err != nil {
			return err
		}
		for _, topic := range module.Topics {
			topicQuery := `INSERT INTO topic(
				moduleid, title, lk, spz, sro, literature)
				VALUES ($1, $2, $3, $4, $5, $6);`
			_, err := sr.db.Exec(c, topicQuery, moduleID, topic.TopicName, topic.LK, topic.SPZ, topic.SRO, topic.Literature)
			if err != nil {
				return err
			}
		}
	}

	deleteQuery = `delete from literature where syllabusid = $1`
	_, err = sr.db.Exec(c, deleteQuery, syllabus.SyllabusID)
	if err != nil {
		return err
	}
	for _, main := range syllabus.Literature.MainLiterature {
		literatureQuery := `INSERT INTO literature(
			syllabusid, type, title)
			VALUES ($1, $2, $3);`
		_, err = sr.db.Exec(c, literatureQuery, syllabus.SyllabusID, "main", main)
		if err != nil {
			return err
		}
	}
	for _, additional := range syllabus.Literature.AdditionalLiterature {
		literatureQuery := `INSERT INTO literature(
			syllabusid, type, title)
			VALUES ($1, $2, $3);`
		_, err = sr.db.Exec(c, literatureQuery, syllabus.SyllabusID, "additional", additional)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sr *SyllabusRepository) Delete(c context.Context, syllabusID int) error {

	deleteQuery := `delete from literature where syllabusid = $1`
	_, err := sr.db.Exec(c, deleteQuery, syllabusID)
	if err != nil {
		return err
	}

	deleteQuery = `delete from modules where syllabusid = $1`
	_, err = sr.db.Exec(c, deleteQuery, syllabusID)
	if err != nil {
		return err
	}

	query := `delete from syllabus where id = $1`
	_, err = sr.db.Exec(c, query, syllabusID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SyllabusRepository) GetByID(c context.Context, syllabusID int, userID uint) (models.Syllabus, error) {
	syllabus := models.Syllabus{}
	fmt.Println(1)
	query := `SELECT id, subject, faculty, kafedra, specialist, coursenumber, creditnumber, allhours, lecturehour, practicehour, 
	sro,srop , madeby, madebymajor,facultyOfProf , email , address , timeofcons, discuss1, discussby1, discussby1major, discuss2, discussby2, discussby2major, confirmedby, confirmedbymajor
	FROM syllabus where id = $1;`
	err := sr.db.QueryRow(c, query, syllabusID).Scan(&syllabus.SyllabusID, &syllabus.MainInfo.SubjectInfo.SubjectName,
		&syllabus.MainInfo.FacultyName, &syllabus.MainInfo.KafedraName, &syllabus.MainInfo.SubjectInfo.SpecialityName,
		&syllabus.MainInfo.CourseNumber, &syllabus.MainInfo.CreditNumber, &syllabus.MainInfo.AllHours, &syllabus.MainInfo.LectureHours,
		&syllabus.MainInfo.PracticeLessons, &syllabus.MainInfo.SRO, &syllabus.MainInfo.SROP,
		&syllabus.Preface.MadeBy.FullName, &syllabus.Preface.MadeBy.Specialist, &syllabus.Preface.MadeBy.Faculty, &syllabus.Preface.MadeBy.Email,
		&syllabus.Preface.MadeBy.Address, &syllabus.Preface.MadeBy.TimeForConsultation,
		&syllabus.Preface.Discussion1, &syllabus.Preface.Discussed1.FullName, &syllabus.Preface.Discussed1.Specialist,
		&syllabus.Preface.Discussion2, &syllabus.Preface.Discussed2.FullName, &syllabus.Preface.Discussed2.Specialist,
		&syllabus.Preface.ConfirmedBy.FullName, &syllabus.Preface.ConfirmedBy.Specialist)
	if err != nil {
		return syllabus, err
	}
	//third page
	fmt.Println(2)
	moduleQuery := `SELECT id, title FROM modules WHERE syllabusid = $1 order by orderid;`
	rows, err := sr.db.Query(c, moduleQuery, syllabusID)
	if err != nil {
		return syllabus, err
	}
	modules := []models.Modules{}
	for rows.Next() {
		var moduleID int
		module := models.Modules{}
		err := rows.Scan(&moduleID, &module.ModuleName)
		if err != nil {
			return syllabus, err
		}
		topicQuery := `SELECT title, lk, spz, sro, literature FROM topic WHERE moduleid = $1 order by id`
		rowss, err := sr.db.Query(c, topicQuery, moduleID)
		if err != nil {
			return syllabus, err
		}
		topics := []models.Topic{}
		for rowss.Next() {
			topic := models.Topic{}
			err := rowss.Scan(&topic.TopicName, &topic.LK, &topic.SPZ, &topic.SRO, &topic.Literature)
			if err != nil {
				return syllabus, err
			}
			topics = append(topics, topic)
		}
		module.Topics = topics
		modules = append(modules, module)
	}
	syllabus.Topics = modules
	//fourth page
	fmt.Println(3)
	literatureQuery := `SELECT type , title FROM literature where syllabusid = $1;`
	rows, err = sr.db.Query(c, literatureQuery, syllabusID)
	if err != nil {
		return syllabus, err
	}
	for rows.Next() {
		var literature, tip string
		err := rows.Scan(&tip, &literature)
		if err != nil {
			return syllabus, err
		}
		if tip == "main" {
			syllabus.Literature.MainLiterature = append(syllabus.Literature.MainLiterature, literature)
		} else {
			syllabus.Literature.AdditionalLiterature = append(syllabus.Literature.AdditionalLiterature, literature)
		}
	}
	//questions
	fmt.Println(4)
	questionQuery := `SELECT title , sequen FROM questions where syllabusid = $1;`
	rows, err = sr.db.Query(c, questionQuery, syllabusID)
	if err != nil {
		return syllabus, err
	}
	for rows.Next() {
		var question, tip string
		err := rows.Scan(&question, &tip)
		if err != nil {
			return syllabus, err
		}
		if tip == "1" {
			syllabus.Question1.Questions = append(syllabus.Question1.Questions, question)
		} else {
			syllabus.Question2.Questions = append(syllabus.Question2.Questions, question)
		}
	}
	//additional text
	fmt.Println(5)
	textQuery := `SELECT text2, text3, text4, text5, text6, text7, text8 FROM public.addition where syllabusid = $1;`
	err = sr.db.QueryRow(c, textQuery, syllabusID).Scan(&syllabus.Text.Text2, &syllabus.Text.Text3, &syllabus.Text.Text4,
		&syllabus.Text.Text5, &syllabus.Text.Text6, &syllabus.Text.Text7, &syllabus.Text.Text8)
	if err != nil {
		return syllabus, err
	}

	return syllabus, nil
}

func (sr *SyllabusRepository) GetAllOwn(c context.Context, userID uint) ([]models.Syllabus, error) {
	syllabuses := []models.Syllabus{}
	query := `SELECT id, subject, faculty, kafedra, specialist, coursenumber, allhours FROM syllabus where userid = $1`
	rows, err := sr.db.Query(c, query, userID)
	if err != nil {
		return syllabuses, err
	}
	for rows.Next() {
		syllabus := models.Syllabus{}
		err := rows.Scan(&syllabus.SyllabusID, &syllabus.MainInfo.SubjectInfo.SubjectName,
			&syllabus.MainInfo.FacultyName, &syllabus.MainInfo.KafedraName, &syllabus.MainInfo.SubjectInfo.SpecialityName,
			&syllabus.MainInfo.CourseNumber, &syllabus.MainInfo.AllHours)
		if err != nil {
			return syllabuses, err
		}
		syllabuses = append(syllabuses, syllabus)
	}
	return syllabuses, nil
}

func (sr *SyllabusRepository) GetAllOthers(c context.Context, userID uint, subject string) ([]models.Syllabus, error) {
	syllabuses := []models.Syllabus{}
	subject = `%` + subject + `%`
	query := `SELECT id, subject, faculty, kafedra, specialist, coursenumber, allhours FROM syllabus where userid <> $1 and subject ilike $2`
	rows, err := sr.db.Query(c, query, userID, subject)
	if err != nil {
		return syllabuses, err
	}
	for rows.Next() {
		syllabus := models.Syllabus{}
		err := rows.Scan(&syllabus.SyllabusID, &syllabus.MainInfo.SubjectInfo.SubjectName,
			&syllabus.MainInfo.FacultyName, &syllabus.MainInfo.KafedraName, &syllabus.MainInfo.SubjectInfo.SpecialityName,
			&syllabus.MainInfo.CourseNumber, &syllabus.MainInfo.AllHours)
		if err != nil {
			return syllabuses, err
		}
		syllabuses = append(syllabuses, syllabus)
	}
	return syllabuses, nil
}
