package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Miguelburitica/mini-projects/go/protobuffers/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)

	student := models.Student{}
	err := row.Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (repo *PostgresRepository) SetExam(ctx context.Context, exam *models.Exam) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO exams (id, name) VALUES ($1, $2)", exam.Id, exam.Name)
	
	return err
}

func (repo *PostgresRepository) GetExam(ctx context.Context, id string) (*models.Exam, error) {
	row := repo.db.QueryRowContext(ctx, "SELECT id, name FROM exams WHERE id = $1", id)

	exam := models.Exam{}
	err := row.Scan(&exam.Id, &exam.Name)
	if err != nil {
		return nil, err
	}

	return &exam, nil
}

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, question, answer, exam_id) VALUES ($1, $2, $3, $4)", question.Id, question.Question, question.Answer, question.ExamId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepository) GetStudentsPerExam(ctx context.Context, ExamId string) ([]*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT students.id, students.name, students.age FROM students INNER JOIN enrollments ON students.id = enrollments.student_id WHERE enrollments.exam_id = $1", ExamId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	students := []*models.Student{}
	for rows.Next() {
		student := models.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	
	return students, nil
}
