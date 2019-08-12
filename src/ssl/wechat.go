package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"log"
)
const(
wechatCertPath = "/path/to/wechat/cert.pem"
wechatKeyPath = "/path/to/wechat/key.pem"
wechatCAPath = "/path/to/wechat/ca.pem"
wechatRefundURL = "https://wechat/refund/url"
)
var _tlsConfig *tls.Config

func getTLSConfig() (*tls.Config, error) {
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	// load cert
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		log.Fatalln("load wechat keys fail", err)
		return nil, err
	}

	// load root ca
	caData, err := ioutil.ReadFile(wechatCAPath)
	if err != nil {
		log.Fatalln("read wechat ca fail", err)
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return _tlsConfig, nil
}

func SecurePost(url string, xmlContent []byte) (*http.Response, error) {
	tlsConfig, err := getTLSConfig()
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(
		wechatRefundURL,
		"text/xml",
		bytes.NewBuffer(xmlContent))
}
