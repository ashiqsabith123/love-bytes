package helper

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"regexp"

	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	"google.golang.org/grpc/credentials"
)

func GetCertificate(ca_cert, client_cert, client_key string) (credentials.TransportCredentials, error) {
	// read ca's cert
	caCert, err := os.ReadFile(ca_cert)
	if err != nil {
		return nil, err
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, err
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair(client_cert, client_key)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	return tlsCredential, nil
}

func CreateResponse(code int32, message, err string) responce.Response {

	return responce.Response{
		Code:    int(code),
		Message: message,
		Error:   err,
	}

}

func IsValidPhoneNumber(phoneNumber string) bool {

	regex := `^[1-9][0-9]{9}$`

	pattern := regexp.MustCompile(regex)

	return pattern.MatchString(phoneNumber)
}
