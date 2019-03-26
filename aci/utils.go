package aci

import (
	"fmt"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/container"
)

func toStrMap(inputMap map[string]interface{}) map[string]string {
	rt := make(map[string]string)
	for key, value := range inputMap {
		rt[key] = value.(string)
	}

	return rt
}

func toStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
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
	arr := strings.Split(dn, "/")
	hashedName := arr[len(arr)-1]
	nameArr := strings.Split(hashedName, "-")
	name := strings.Join(nameArr[1:], "-")
	return name

}

func GetParentDn(childDn string) string {
	arr := strings.Split(childDn, "/")
	// in case of cidr blocks we have extra / in the ip range so let's catch it and remove. This will have extra part.
	lastElement := arr[len(arr)-1]
	if strings.Contains(lastElement, "]") {
		backSlashedDn := strings.Join(arr[:len(arr)-1], "/")
		// split on - to remove last element.
		arr = strings.Split(backSlashedDn, "/")
		// remove last 2 elements as that will contain the extra part like rn - ip

		dnWithSlash := strings.Join(arr[:len(arr)-1], "/")

		return strings.TrimSuffix(dnWithSlash, "/")

	}

	return strings.Join(arr[:len(arr)-1], "/")

}
