package apiserver

type APIserver struct{}

func New() *APIserver{
	return &APIserver{}
}

func (s *APIserver) Start()error{
	return nil
}