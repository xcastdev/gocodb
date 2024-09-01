package gocodb

import (
	"encoding/json"
	"fmt"
)

type Table struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	TableName   string                 `json:"table_name"`
	BaseID      string                 `json:"base_id"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Columns     []Column               `json:"columns"`
	Views       []View                 `json:"views"`
	Meta        map[string]interface{} `json:"meta"`
	Schema      map[string]interface{} `json:"schema"`
	Enabled     bool                   `json:"enabled"`
	Order       int                    `json:"order"`
	Deleted     bool                   `json:"deleted"`
	Type        string                 `json:"type"`
	Tags        []string               `json:"tags"`
	ColumnsByID map[string]Column      `json:"columnsById"`
}

type Column struct {
	ID                string `json:"id"`
	ColumnName        string `json:"column_name"`
	DataType          string `json:"dt"`
	PrimaryKey        bool   `json:"pk"`
	AutoIncrement     bool   `json:"ai"`
	Unique            bool   `json:"unique"`
	NotNull           bool   `json:"rqd"`
	Length            string `json:"clen"`
	DefaultValue      string `json:"cdf"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Comment           string `json:"cc"`
	SpecificType      string `json:"dtx"`
	SpecificTypeParam string `json:"dtxp"`
}

type View struct {
	ID        string                 `json:"id"`
	Title     string                 `json:"title"`
	TableName string                 `json:"table_name"`
	BaseID    string                 `json:"base_id"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	Meta      map[string]interface{} `json:"meta"`
	LockType  string                 `json:"lock_type"`
	IsDefault bool                   `json:"is_default"`
}

func (c *Client) CreateTable(baseID string, tableData interface{}) (*Table, error) {
	path := fmt.Sprintf("/api/v2/meta/bases/%s/tables", baseID)
	resp, err := c.post(path, tableData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var table Table
	err = json.NewDecoder(resp.Body).Decode(&table)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &table, nil
}

func (c *Client) GetTableInfo(tableID string) (*Table, error) {
	path := fmt.Sprintf("/api/v2/meta/tables/%s", tableID)
	resp, err := c.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var table Table
	err = json.NewDecoder(resp.Body).Decode(&table)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &table, nil
}

func (c *Client) UpdateTable(tableID string, updateData interface{}) (*Table, error) {
	path := fmt.Sprintf("/api/v2/meta/tables/%s", tableID)
	resp, err := c.patch(path, updateData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var table Table
	err = json.NewDecoder(resp.Body).Decode(&table)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &table, nil
}

func (c *Client) DeleteTable(tableID string) (bool, error) {
	path := fmt.Sprintf("/api/v2/meta/tables/%s", tableID)
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
