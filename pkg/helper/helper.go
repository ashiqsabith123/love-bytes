package helper

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/ashiqsabith123/api-gateway/pkg/config"
	"github.com/ashiqsabith123/api-gateway/pkg/models/responce"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"

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

func CreateResponse(code int32, message string, err any, data any) responce.Response {

	return responce.Response{
		Code:    int(code),
		Message: message,
		Error:   err,
		Data:    data,
	}

}

func IsValidPhoneNumber(phoneNumber string) bool {

	regex := `^[1-9][0-9]{9}$`

	pattern := regexp.MustCompile(regex)

	return pattern.MatchString(phoneNumber)
}

func Validator(data interface{}) error {
	validte := validator.New()

	err := validte.Struct(data)

	if err != nil {
		return err
	}

	return nil
}

func ValidateJWTTokens(token string) (jwt.MapClaims, error) {

	parseToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method: ")
		}
		return []byte(config.GetSecretKey()), nil

	})

	if err != nil {
		return nil, err
	}

	if !parseToken.Valid {
		return nil, errors.New("token not valid")

	}

	claim, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	if float64(time.Now().Unix()) > claim["exp"].(float64) {
		return nil, errors.New("token expired")
	}

	return claim, nil

}
