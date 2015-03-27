package sources

import (
	"net/http"
	"io/ioutil"
	"github.com/golang/glog")

func PostRequestAndGetValue(client *http.Client, req *http.Request, value interface{}) error {
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	glog.Info(string(bytes))

//	dec := json.NewDecoder(response.Body)
//	dec.UseNumber()
//	err = dec.Decode(value)
//	if err != nil {
//		return err
//	}
	return nil
}
