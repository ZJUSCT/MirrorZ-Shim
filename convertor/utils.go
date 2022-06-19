package convertor

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func readJson[T any](unmarshallDest T, filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error("Read cname.json failed")
		return err
	}
	err = json.Unmarshal(file, unmarshallDest)
	if err != nil {
		logrus.Error("Unmarshal cname.json failed")
		return err
	}
	return nil
}
