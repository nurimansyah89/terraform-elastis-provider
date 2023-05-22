package client

type ElastisClient struct {
	EndpointURL string
	Token       string
}

func New(token string) *ElastisClient {
	return &ElastisClient{
		Token:       token,
		EndpointURL: "https://api.elastis.id/v1",
		// EndpointURL: "https://d8d1ab31-5589-46e0-99d3-9c9d6ce02909.mock.pstmn.io",
	}
}
