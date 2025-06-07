package utils

import (
	"io"
	"net/http"
	"strings"
)

func GetPublicIp() (string, error) {
	resp, err := http.Get("https://ifconfig.co")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(ip)), nil
}
