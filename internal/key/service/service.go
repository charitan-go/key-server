package service

import "github.com/charitan-go/key-server/pkg/proto"

type KeyService interface {
	GetPrivateKey(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error)
}

type keyServiceImpl struct {
}

func NewKeyService() KeyService {
	return &keyServiceImpl{}
}

func (svc *keyServiceImpl) GetPrivateKey(reqDto *proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error) {
	// TODO: Impl
	return nil, nil
}
