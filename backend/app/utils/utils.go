package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fi44er/ton-backend/utils/response"
	"github.com/gofiber/fiber/v2"
)

func getTransaction(hash string) (response.Data, error) {
	url := "https://tonapi.io/v2/traces/" + hash

	var data response.Data
	res, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func isTransactionEmpty(tx response.Data) bool {
	return tx.Transaction.Hash == ""
}

func CheckValidTransaction(hash string) error {
	data, err := getTransaction(hash)
	if err != nil {
		return err
	}

	if isTransactionEmpty(data) {
		time.Sleep(3 * time.Minute)
		data, err = getTransaction(hash)
		if err != nil {
			return err
		}

		if isTransactionEmpty(data) {
			return fiber.NewError(404, "Transaction not found")
		}
	}

	if !isValid(data.Transaction) {
		return fiber.NewError(400, "Transaction not valid")
	}

	for i := 0; i < len(data.Children); i++ {
		if !isValid(data.Children[i].Transaction) {
			return fiber.NewError(400, "Transaction not valid")
		}
	}

	return nil
}

func isValid(tx response.Transaction) bool {
	if !tx.Success || tx.Destroyed || tx.Account.IsScam {
		return false
	}
	return true
}
