package usecases

import (
	"code-space-backend-api/app/Course/interfaces/dto/output"
	"code-space-backend-api/app/Course/interfaces/repository"
	"code-space-backend-api/common/errors"
)

const listContentsContext = "usecase/list-modules"

var (
	couldNotFindContentWithGivenChapterHash = errors.NotFound.
		WithContext(listContentsContext).
		WithName("could_not_find_modules_with_given_course_hash").
		WithTemplate("failed to get modules associated to the course: %v")
)

type ListContentFromChapter interface {
	Execute(moduleHash string) (output.ChapterOutputDTO, error)
}

type listContentFromModule struct {
	chapterRepository repository.ChapterRepository
}

func NewListContentFromChapter(chapterRepository repository.ChapterRepository) ListContentFromChapter {
	return &listContentFromModule{
		chapterRepository: chapterRepository,
	}
}

func (usecase *listContentFromModule) Execute(chapterHash string) (output.ChapterOutputDTO, error) {
	chapter, err := usecase.chapterRepository.ListContentByChapterHash(chapterHash)
	if err != nil {
		return chapter, couldNotFindContentWithGivenChapterHash.WithArgs(err.Error())
	}

	return chapter, nil
}
