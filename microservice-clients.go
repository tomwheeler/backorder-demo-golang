package app

// simple clients for calling the (local) microservices
// the implementation is never shown in the example
import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

type InventoryService struct {
	Hostname string // just for show, this is never used
}

func (client InventoryService) GetQuantityAvailable(productId string) (int, error) {
	params := map[string]string{
		"productId": productId,
	}
	baseUrl := "http://localhost:8009/get-quantity-available"
	resty := resty.New()
	resp, err := resty.R().SetQueryParams(params).Get(baseUrl)
	if err != nil {
		return -1, err
	}

	bodyAsString := resp.String()
	retval, err := strconv.Atoi(bodyAsString)
	if err != nil {
		return -1, err
	}
	return retval, nil
}

type NotificationService struct {
	Hostname string // just for show, this is never used
}

func (client NotificationService) SendText(message string, phoneNum string) (string, error) {
	params := map[string]string{
		"message": message,
		"number":  phoneNum,
	}
	baseUrl := "http://localhost:8008/send-text"
	resty := resty.New()
	resp, err := resty.R().SetQueryParams(params).Get(baseUrl)
	if err != nil {
		return "", err
	}

	bodyAsString := resp.String()
	return bodyAsString, nil
}
