package aci

import (
	"fmt"
	"regexp"

	"github.com/ciscoecosystem/aci-go-client/container"
)

func toStrMap(inputMap map[string]interface{}) map[string]string {
	rt := make(map[string]string)
	for key, value := range inputMap {
		rt[key] = value.(string)
	}

	return rt
}

func preparePayload(className string, inputMap map[string]string) (*container.Container, error) {
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
			}
		}
	}`, className))
	cont, err := container.ParseJSON(containerJSON)

	if err != nil {
		return nil, err
	}
	for key, value := range inputMap {
		cont.Set(value, className, "attributes", key)
	}
	return cont, nil

}

func GetMOName(dn string) string {
	re := regexp.MustCompile(".*/\\S+-(\\S+.*)$")
	match := re.FindStringSubmatch(dn)
	return match[1]

}

func GetParentDn(childDn string) string {
	re := regexp.MustCompile("(.*)/\\S+-\\S+.*$")
	match := re.FindStringSubmatch(childDn)
	return match[1]

}
