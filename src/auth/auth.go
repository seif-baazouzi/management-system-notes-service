package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func IsWorkspaceOwner(c *fiber.Ctx) error {
	conn := http.Client{Timeout: time.Duration(3) * time.Second}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/api/v1/workspaces/%s", os.Getenv("WORKSPACES_SERVICE"), c.Params("workspaceID")),
		nil,
	)
	if err != nil {
		return utils.ServerError(c, err)
	}

	req.Header.Add("X-Token", c.Get("X-Token"))

	resp, err := conn.Do(req)
	if err != nil {
		return utils.ServerError(c, err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	var body map[string]interface{}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	_, exist := body["workspace"]
	if !exist {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-token"})
	}

	workspace := body["workspace"].(map[string]interface{})

	c.Locals("uuid", workspace["userID"])
	return c.Next()
}

func IsLogin(c *fiber.Ctx) error {
	conn := http.Client{Timeout: time.Duration(3) * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/user", os.Getenv("ACCOUNTS_SERVICE")), nil)
	if err != nil {
		return utils.ServerError(c, err)
	}

	req.Header.Add("X-Token", c.Get("X-Token"))

	resp, err := conn.Do(req)
	if err != nil {
		return utils.ServerError(c, err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	var body map[string]string
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return utils.ServerError(c, err)
	}

	_, exist := body["userID"]
	if !exist {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-token"})
	}

	c.Locals("uuid", body["userID"])
	return c.Next()
}
