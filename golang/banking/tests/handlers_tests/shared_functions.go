package handlers_tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/shopspring/decimal"
)

func RegisterAndGetTokens(phone string, inn string) (Tokens, error) {
	err := signup(phone, inn)
	if err != nil {
		panic(fmt.Sprintf("Unable to signup %s", err))
	}
	return getAccessToken(phone)
}

func signup(phone string, inn string) error {
	url := "http://localhost:8000/api/banking/signup/"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	  "first_name": "Vasilii",
	  "last_name": "Popov",
	  "email": "dopefresh4000@gmail.com",
	  "password": "myhardpassword",
	  "phone": "%s",
	  "client": {
		  "residential_address": "Lesnaya 38",
		  "registration_address": "Lesnaya 38",
		  "client_type": "individual",
		  "ogrn": "%s",
		  "inn": "%s",
		  "kpp": "129581"
	  }
  }`, phone, inn, inn))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	return err
}

func getAccessToken(phone string) (Tokens, error) {
	tokens := Tokens{}

	url := "http://localhost:8000/api/token/"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
		"phone": "%s",
		"password": "myhardpassword"
	}`, phone))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return tokens, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return tokens, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return tokens, err
	}
	err = json.Unmarshal(body, &tokens)
	return tokens, err
}

func createCard(cardId string, balance decimal.Decimal, token string) error {
	url := "http://localhost:8080/api/v1/cards/"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	"cardId": "%s",
	"balance": %s,
	"validTo": "2023-07-21T17:32:28Z",
	"cvcCode": "010",
	"cardType": "credit",
	"currency": "RUB"
	}`, cardId, balance))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	return err
}

func getCard(cardId string, token string) (models.Card, error) {
	url := fmt.Sprintf("http://localhost:8080/api/v1/cards/%s/", cardId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return models.Card{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return models.Card{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Card{}, err
	}
	var card models.Card

	err = json.Unmarshal(body, &card)
	return card, err
}

func getNoCard(cardId string, token string) (int, error) {
	// runtime.Breakpoint()
	url := fmt.Sprintf("http://localhost:8080/api/v1/cards/%s/", cardId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, err

}

func updateCard(cardId string, cvcCode string, token string) error {
	url := fmt.Sprintf("http://localhost:8080/api/v1/cards/%s/", cardId)
	method := "PUT"

	payload := strings.NewReader(fmt.Sprintf(`{
	"validTo": "2017-07-21T17:32:28Z",
	"cvcCode": "%s"
  }`, cvcCode))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func deleteCard(cardId string, token string) (int, error) {
	url := fmt.Sprintf("http://localhost:8080/api/v1/cards/%s/", cardId)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, err
}

func getClient(token string) (models.Client, error) {
	url := "http://localhost:8080/api/v1/client/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return models.Client{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return models.Client{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Client{}, err
	}
	var clientModel models.Client
	err = json.Unmarshal(body, &clientModel)
	return clientModel, err
}

func updateClient(token string, phone string) (int, error) {
	url := "http://localhost:8080/api/v1/client/"
	method := "PUT"

	payload := strings.NewReader(fmt.Sprintf(`{
	"firstName": "Vasilii",
	"lastName": "Popov",
	"phone": "%s",
	"registrationAddress": "Russia, Moscow, Lev Tolstoy street 5",
	"residentialAddress": "Russia, Moscow, Pushkina street 4",
	"clientType": "jp"
  }`, phone))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, err
}

func transfer(cardFromId string, cardToId string, sum int64, token string) error {
	url := "http://localhost:8080/api/v1/client/transfer"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	"cardFromId": "%s",
	"cardToId": "%s",
	"summ": %d
	}`, cardFromId, cardToId, sum))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func deposit(cardId string, sum int, token string) error {
	url := "http://localhost:8080/api/v1/client/deposit"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	"cardId": "%s",
	"summ": %d
	}`, cardId, sum))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func withdraw(cardId string, sum int, token string) error {
	url := "http://localhost:8080/api/v1/client/withdraw"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	"cardId": "%s",
	"summ": %d
	}`, cardId, sum))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func getTransactions(token string) ([]models.Transaction, error) {
	url := "http://127.0.0.1:8080/api/v1/transactions/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return []models.Transaction{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return []models.Transaction{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []models.Transaction{}, err
	}
	var transactions []models.Transaction
	err = json.Unmarshal(body, &transactions)
	return transactions, err
}
