package models

type Student struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Age   int32   `json:"age"`
}

type Exam struct {
	Id      string  `json:"id"`
	Name 		string  `json:"name"`
}

type Question struct {
	Id       string  `json:"id"`
	Question string  `json:"question"`
	Answer   string  `json:"answer"`
	ExamId	 string  `json:"exam_id"`
}

type Enrollment struct {
	StudentId string  `json:"student_id"`
	ExamId    string  `json:"exam_id"`
}
