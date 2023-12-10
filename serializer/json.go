package serializer

import (
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobuftoJSONFile(message proto.Message, filename string) error {
	marshaller := protojson.MarshalOptions{
		UseProtoNames:  true,
		Indent:         "    ",
		UseEnumNumbers: true,
	}
	marshaler, err := marshaller.Marshal(message)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, marshaler, 0644)

	if err != nil {
		return err
	}

	return nil
}
