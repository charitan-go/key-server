package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/charitan-go/key-server/external/auth"
	"github.com/charitan-go/key-server/pkg/proto"
)

type KeyService interface {
	getPrivateKeyStr() string

	GenerateKeyPairs() error

	// GRPC Listener
	GetPrivateKeyGrpcHandler(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error)
}

type keyServiceImpl struct {
	privateKey           *rsa.PrivateKey
	publicKey            *rsa.PublicKey
	authRabbitmqProducer auth.AuthRabbitmqProducer
}

func NewKeyService(authRabbitmqProducer auth.AuthRabbitmqProducer) KeyService {
	return &keyServiceImpl{nil, nil, authRabbitmqProducer}
}

func (svc *keyServiceImpl) getPrivateKeyStr() string {
	// Convert the RSA private key to DER-encoded PKCS#1 format
	privateKeyDER := x509.MarshalPKCS1PrivateKey(svc.privateKey)

	// Create a PEM block with the DER-encoded key
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDER,
	}

	// Encode the PEM block to a []byte
	privateKeyPEM := pem.EncodeToMemory(pemBlock)

	// Convert the PEM []byte to a string
	privateKeyString := string(privateKeyPEM)
	return privateKeyString
}

func (svc *keyServiceImpl) GenerateKeyPairs() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Cannot generate private key: %v\n", err)
		return err
	}

	// Store both keys
	svc.privateKey = privateKey
	svc.publicKey = &privateKey.PublicKey

	// Send noti to auth server for get key pairs
	err = svc.authRabbitmqProducer.NotiGetPrivateKey()
	if err != nil {
		log.Fatalf("Cannot send noti to auth server: %v\n", err)
	}

	return nil
}

func (svc *keyServiceImpl) GetPrivateKeyGrpcHandler(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error) {
	if svc.privateKey == nil {
		return nil, fmt.Errorf("Private key not available")
	}

	return &proto.GetPrivateKeyResponseDto{PrivateKey: svc.getPrivateKeyStr()}, nil
}
