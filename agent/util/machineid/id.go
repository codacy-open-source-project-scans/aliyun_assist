package machineid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetMachineID() (string, error) {
	id, err := machineID()
	if err != nil {
		return "", fmt.Errorf("machineid: %v", err)
	}
	return id, nil
}


func protect(appID, id string) string {
	mac := hmac.New(sha256.New, []byte(id))
	mac.Write([]byte(appID))
	return hex.EncodeToString(mac.Sum(nil))
}


func trim(s string) string {
	return strings.TrimSpace(strings.Trim(s, "\n"))
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}