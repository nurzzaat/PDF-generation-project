package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)

type SyllabusRepository struct {
	db *pgxpool.Pool
}

func NewSyllabusRepository(db *pgxpool.Pool) models.SyllabusRepository {
	return &SyllabusRepository{db: db}
}

func (sr *SyllabusRepository) Create(c context.Context, syllabusInfo models.SyllabusInfo, userID uint) (int , error) {
	var id int
	query := `INSERT INTO syllabus(
		userid, subject, faculty, kafedra, specialist, coursenumber, creditnumber, allhours, lecturehour, practicehour, sro)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) retuning id;`
	err := sr.db.QueryRow(c, query, userID, syllabusInfo.SubjectInfo.SubjectName,
		syllabusInfo.FacultyName, syllabusInfo.KafedraName, syllabusInfo.SubjectInfo.SpecialityName, syllabusInfo.CourseNumber,
		syllabusInfo.CreditNumber, syllabusInfo.AllHours, syllabusInfo.LectureHours, syllabusInfo.PracticeLessons, syllabusInfo.SRO).Scan(&id)
	if err != nil {
		return id , err
	}
	return id ,nil
}

func (sr *SyllabusRepository) Update(c context.Context, userID uint, syllabus models.Syllabus) error {
	query := `UPDATE syllabus
	SET subject=$1, faculty=$2, kafedra=$3, specialist=$4, coursenumber=$5, creditnumber=$6, allhours=$7, 
	lecturehour=$8, practicehour=$9, sro=$10, madeby=$11, madebymajor=$12, discuss1=$13, discussby1=$14, 
	discussby1major=$15, discuss2=$16, discussby2=$17, discussby2major=$18, confirmedby=$19, confirmedbymajor=$20
	WHERE id = $21`
	_, err := sr.db.Exec(c, query, syllabus.MainInfo.SubjectInfo.SubjectName, syllabus.MainInfo.FacultyName,
		syllabus.MainInfo.KafedraName, syllabus.MainInfo.SubjectInfo.SpecialityName, syllabus.MainInfo.CourseNumber,
		syllabus.MainInfo.CreditNumber, syllabus.MainInfo.AllHours, syllabus.MainInfo.LectureHours,
		syllabus.MainInfo.PracticeLessons, syllabus.MainInfo.SRO, syllabus.Preface.MadeBy.FullName, syllabus.Preface.MadeBy.Specialist,
		syllabus.Preface.Discussion1, syllabus.Preface.Discussed1.FullName, syllabus.Preface.Discussed1.Specialist,
		syllabus.Preface.Discussion2, syllabus.Preface.Discussed2.FullName, syllabus.Preface.Discussed2.Specialist,
		syllabus.Preface.ConfirmedBy.FullName, syllabus.Preface.ConfirmedBy.Specialist, syllabus.SyllabusID)
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
	query := `delete from syllabus where id = $1`
	_, err := sr.db.Exec(c, query, syllabusID)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SyllabusRepository) GetByID(c context.Context, syllabusID int, userID uint) (models.Syllabus, error) {
	syllabus := models.Syllabus{}
	query := `SELECT id, subject, faculty, kafedra, specialist, coursenumber, creditnumber, allhours, lecturehour, practicehour, 
	sro, madeby, madebymajor, discuss1, discussby1, discussby1major, discuss2, discussby2, discussby2major, confirmedby, confirmedbymajor
	FROM syllabus where id = $1;`
	err := sr.db.QueryRow(c, query, syllabusID).Scan(&syllabus.SyllabusID, &syllabus.MainInfo.SubjectInfo.SubjectName,
		&syllabus.MainInfo.FacultyName, &syllabus.MainInfo.KafedraName, &syllabus.MainInfo.SubjectInfo.SpecialityName,
		&syllabus.MainInfo.CourseNumber, &syllabus.MainInfo.CreditNumber, &syllabus.MainInfo.AllHours, &syllabus.MainInfo.LectureHours,
		&syllabus.MainInfo.PracticeLessons, &syllabus.MainInfo.SRO, &syllabus.Preface.MadeBy.FullName, &syllabus.Preface.MadeBy.Specialist,
		&syllabus.Preface.Discussion1, &syllabus.Preface.Discussed1.FullName, &syllabus.Preface.Discussed1.Specialist,
		&syllabus.Preface.Discussion2, &syllabus.Preface.Discussed2.FullName, &syllabus.Preface.Discussed2.Specialist,
		&syllabus.Preface.ConfirmedBy.FullName, &syllabus.Preface.ConfirmedBy.Specialist)
	if err != nil {
		return syllabus, err
	}

	//third page
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
	rows, err := sr.db.Query(c, query, userID , subject)
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
