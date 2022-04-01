package clicker

import (
	"encoding/base64"
	"google.golang.org/protobuf/proto"
	homework "kadam_test/internal/proto"
)

func (s *Service) DecodeProtoClick(data string) (*homework.Click, error) {
	sec := "imSoVerySafe"

	raw, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(raw); i++ {
		raw[i] = raw[i] ^ sec[i%len(sec)]
	}

	click, err := unMarshalClick(raw)
	if err != nil {
		return nil, err
	}

	return click, nil
}

func unMarshalClick(data []byte) (*homework.Click, error) {
	click := &homework.Click{}

	if err := proto.Unmarshal(data, click); err != nil {
		return nil, err
	}

	return click, nil
}
