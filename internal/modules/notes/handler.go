package notes


func NewHandler(service Service) *handler {
	return &handler{service: service}
}

