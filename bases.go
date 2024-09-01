package gocodb

import (
	"encoding/json"
	"fmt"
)

// Structs

type Base struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Color       string                 `json:"color"`
	Prefix      string                 `json:"prefix"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Deleted     bool                   `json:"deleted"`
	IsMeta      bool                   `json:"is_meta"`
	Meta        map[string]interface{} `json:"meta"`
	Order       int                    `json:"order"`
	Status      string                 `json:"status"`
}

type BaseList struct {
	List     []Base   `json:"list"`
	PageInfo PageInfo `json:"pageInfo"`
}

type PageInfo struct {
	IsFirstPage bool `json:"isFirstPage"`
	IsLastPage  bool `json:"isLastPage"`
	Page        int  `json:"page"`
	PageSize    int  `json:"pageSize"`
	TotalRows   int  `json:"totalRows"`
}

// Client Functions

func (c *Client) ListBases() (*BaseList, error) {
	resp, err := c.get("/api/v2/meta/bases/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var baseList BaseList
	err = json.NewDecoder(resp.Body).Decode(&baseList)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &baseList, nil
}

func (c *Client) CreateBase(baseData interface{}) (*Base, error) {
	resp, err := c.post("/api/v2/meta/bases/", baseData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var base Base
	err = json.NewDecoder(resp.Body).Decode(&base)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &base, nil
}

func (c *Client) GetBase(baseID string) (*Base, error) {
	path := fmt.Sprintf("/api/v2/meta/bases/%s", baseID)
	resp, err := c.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var base Base
	err = json.NewDecoder(resp.Body).Decode(&base)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &base, nil
}

func (c *Client) UpdateBase(baseID string, updateData interface{}) (*Base, error) {
	path := fmt.Sprintf("/api/v2/meta/bases/%s", baseID)
	resp, err := c.patch(path, updateData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var base Base
	err = json.NewDecoder(resp.Body).Decode(&base)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &base, nil
}

func (c *Client) DeleteBase(baseID string) (bool, error) {
	path := fmt.Sprintf("/api/v2/meta/bases/%s", baseID)
	resp, err := c.delete(path)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result bool
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, fmt.Errorf("error decoding response: %v", err)
	}
	return result, nil
}

func (c *Client) ListTablesInBase(baseID string) ([]Table, error) {
	path := fmt.Sprintf("/api/v2/meta/bases/%s/tables", baseID)
	resp, err := c.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tables []Table
	err = json.NewDecoder(resp.Body).Decode(&tables)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return tables, nil
}
