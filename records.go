package gocodb

import (
	"encoding/json"
	"fmt"
)

type Record map[string]interface{}

type RecordList struct {
	List     []Record `json:"list"`
	PageInfo PageInfo `json:"pageInfo"`
}

type RecordCount struct {
	Count int `json:"count"`
}

func (c *Client) ListTableRecords(tableID string, queryParams map[string]string) (*RecordList, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", tableID)
	resp, err := c.get(path + formatQueryParams(queryParams))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recordList RecordList
	err = json.NewDecoder(resp.Body).Decode(&recordList)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &recordList, nil
}

func (c *Client) CreateTableRecords(tableID string, recordData []interface{}) ([]Record, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", tableID)
	resp, err := c.post(path, recordData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var records []Record
	err = json.NewDecoder(resp.Body).Decode(&records)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return records, nil
}

func (c *Client) UpdateTableRecords(tableID string, recordData interface{}) ([]Record, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", tableID)
	resp, err := c.patch(path, recordData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var records []Record
	err = json.NewDecoder(resp.Body).Decode(&records)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return records, nil
}

func (c *Client) DeleteTableRecords(tableID string, recordIDs []interface{}) (bool, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records", tableID)
	resp, err := c.deleteWithBody(path, recordIDs)
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

func (c *Client) GetTableRecord(tableID, recordID string, queryParams map[string]string) (*Record, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records/%s", tableID, recordID)
	resp, err := c.get(path + formatQueryParams(queryParams))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var record Record
	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &record, nil
}

func (c *Client) CountTableRecords(tableID string, queryParams map[string]string) (*RecordCount, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/records/count", tableID)
	resp, err := c.get(path + formatQueryParams(queryParams))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var count RecordCount
	err = json.NewDecoder(resp.Body).Decode(&count)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &count, nil
}

func (c *Client) ListLinkedRecords(tableID, linkFieldID, recordID string, queryParams map[string]string) (*RecordList, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%s", tableID, linkFieldID, recordID)
	resp, err := c.get(path + formatQueryParams(queryParams))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recordList RecordList
	err = json.NewDecoder(resp.Body).Decode(&recordList)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &recordList, nil
}

func (c *Client) LinkRecords(tableID, linkFieldID, recordID string, linkData interface{}) (bool, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%s", tableID, linkFieldID, recordID)
	resp, err := c.post(path, linkData)
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

func (c *Client) UnlinkRecords(tableID, linkFieldID, recordID string, unlinkData interface{}) (bool, error) {
	path := fmt.Sprintf("/api/v2/tables/%s/links/%s/records/%s", tableID, linkFieldID, recordID)
	resp, err := c.deleteWithBody(path, unlinkData)
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
