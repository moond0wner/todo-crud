package web_service

type WebService struct {
	webRepository WebRepository
}

type WebRepository interface {
	GetFile(filepath string) ([]byte, error)
}

func NewWebService(wr WebRepository) *WebService {
	return &WebService{
		webRepository: wr,
	}
}
