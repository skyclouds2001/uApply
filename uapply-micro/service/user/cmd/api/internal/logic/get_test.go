package logic

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetOrgEnrolling(t *testing.T) {
	url := "http://localhost:9003/user/org-enrolling"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Error(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(body))
}

func BenchmarkGetOrgEnrolling(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {

		err := GetOrgEnrollingB()
		if err != nil {
			b.Error(err)
		}

	}
	b.StopTimer()
}

func GetOrgEnrollingB() error {
	url := "http://localhost:9003/user/org-enrolling"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}
