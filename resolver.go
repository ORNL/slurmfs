package slurmfs

import (
    "encoding/json"
    "fmt"
    "bytes"
    "io"
    "io/ioutil"
    "net/http"
)

func (client Client) Call(method, path string, data []byte) ([]byte, error) {
	var inp io.Reader
	if data != nil {
		inp = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method,
		"http://unix/slurm/v0.0.38/"+path, inp)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (client Client) Jobs() (resp JobsResponse, err error) {
	// recommended parameter: update_time: int64
	body, err := client.Call("GET", "jobs", nil)
	if err != nil {
		return
	}
	json.Unmarshal(body, &resp)
	return
}

func (client Client) Create(script string, prop JobProps) (resp SubmitResponse, err error) {
	req, err := json.Marshal(&JobSubmission{
                Script: script,
                Job:	prop,
	})
	//fmt.Printf("Sending %s\n", string(req))
	if err != nil {
		return
	}
	body, err := client.Call("POST", "job/submit", req)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

func (client Client) Update(job_id int, prop JobProps) (resp SlurmResponse, err error) {
	path := fmt.Sprintf("job/%d", job_id)
	req, err := json.Marshal(&prop)
	if err != nil {
		return
	}
	body, err := client.Call("POST", path, req)
	if err != nil {
		return
	}
        if err = json.Unmarshal(body, &resp); err != nil {
		return
	}
	return
}

func (client Client) Remove(job_id int) error {
	path := fmt.Sprintf("job/%d", job_id)
	// Note: can send signal: int as a parameter in msg body
	_, err := client.Call("DELETE", path, nil)
	return err
}

