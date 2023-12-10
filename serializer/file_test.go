package serializer

import (
	"grpc_learn/sample"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	laptop_1 := sample.NewLaptop()
	laptop_2 := sample.NewLaptop()

	if err := WriteProtobufToBinaryFile(laptop_1, "../tmp/laptop_1.bin"); err != nil {
		t.Fatal("Cannot write binary file: ", err)
	}

	if err := WriteProtobuftoJSONFile(laptop_2, "../tmp/laptop_2.json"); err != nil {
		t.Fatal("Cannot write json file: ", err)
	}
}
