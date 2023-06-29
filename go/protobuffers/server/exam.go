package server

import (
	"context"
	"io"

	"github.com/Miguelburitica/mini-projects/go/protobuffers/exampb"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/models"
	"github.com/Miguelburitica/mini-projects/go/protobuffers/repository"
)

type ExamServer struct {
	repo repository.Repository
	exampb.UnimplementedExamServiceServer
}

func NewExamServer(repo repository.Repository) *ExamServer {
	return &ExamServer{repo, exampb.UnimplementedExamServiceServer{}}
}

func (s *ExamServer) GetExam(ctx context.Context, req *exampb.GetExamRequest) (*exampb.Exam, error) {
	exam, err := s.repo.GetExam(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &exampb.Exam{
		Id:   exam.Id,
		Name: exam.Name,
	}, nil
}

func (s *ExamServer) SetExam(ctx context.Context, req *exampb.Exam) (*exampb.SetExamResponse, error) {
	exam := &models.Exam{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetExam(ctx, exam)
	if err != nil {
		return nil, err
	}

	return &exampb.SetExamResponse{
		Id: exam.Id,
		Name: exam.Name,
	}, nil
}

func (s *ExamServer) SetQuestions(stream exampb.ExamService_SetQuestionsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Question: msg.GetQuestion(),
			Answer:   msg.GetAnswer(),
			ExamId:   msg.GetExamId(),
		}

		err = s.repo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&exampb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}
