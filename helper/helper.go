package helper

import (
	"crypto/tls"
	"net/http"
)

type CertAuth struct {
	CertPath string
	KeyPath  string
}

func NewHttpClient(certAuth *CertAuth, isTLS bool) (*http.Client, error) {
	httpClient := &http.Client{}
	var tlsConfig *tls.Config
	if certAuth != nil {
		cert, err := tls.LoadX509KeyPair(certAuth.CertPath, certAuth.KeyPath)
		if err != nil {
			return nil, err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}

	if isTLS || certAuth != nil {
		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}
		tlsConfig.InsecureSkipVerify = true
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		httpClient = &http.Client{Transport: transport}
	}

	return httpClient, nil
}
