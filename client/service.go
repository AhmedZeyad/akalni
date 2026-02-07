package client

type ClientService struct {
	client ClientRepo
}

func NewClientService(client ClientRepo) *ClientService {
	return &ClientService{client: client}
}

func (s ClientService) GetProfile(id int) (client Client, err error) {
	client, err = s.client.GetByID(id)
	if err != nil {
		return client, err
	}
	return client, nil
}
