package main

import (
	"fmt"
	"grpcdemo/service"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	pbfile := "./tmp/pb.bin"
	// jsonfile := "./tmp/pb.json"
	// err := writeBinfile(pbfile, 10)
	// if err != nil {
	// 	fmt.Println("bin write error", err)
	// 	return
	// }
	// err = writeJsonfile(jsonfile, 10)
	// if err != nil {
	// 	fmt.Println("json write error", err)
	// 	return
	// }
	resp := &service.OutputResponse{
		Stock: 1,
		Name:  []string{"hello", "world"},
	}
	err := writeBinfile(pbfile, resp, 1)
	if err != nil {
		fmt.Println("bin write error", err)
		return
	}
	resp1, err := readBinfile(pbfile)
	if err != nil {
		fmt.Println("bin read error", err)
		return
	}

	// fmt.Println(resp.Stock == resp1.Stock)
	fmt.Println(proto.Equal(resp, resp1))
	fmt.Println("done")
}

func writeBinfile(filename string, resp *service.OutputResponse, n int) error {
	pbfile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("crete file error", err)
	}
	// resp := &service.OutputResponse{
	// 	Stock: 1,
	// 	Name:  []string{"hello", "world"},
	// }
	res, err := proto.Marshal(resp)
	if err != nil {
		return fmt.Errorf("marshal pb data error", err)

	}
	for i := 0; i < n; i++ {
		_, err = pbfile.Write(res)
		if err != nil {
			return fmt.Errorf("write to file error ", err)
		}
	}

	return nil
}
func readBinfile(filename string) (*service.OutputResponse, error) {
	pbfile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open bin file error %s", err)
	}
	data, err := ioutil.ReadAll(pbfile)
	if err != nil {
		return nil, fmt.Errorf("read file error %s", err)
	}
	var resp = new(service.OutputResponse)
	err = proto.Unmarshal(data, resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal data error %s", err)
	}
	return resp, nil
}

func writeJsonfile(filename string, n int) error {
	jsonfile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("crete file error", err)
	}
	resp := &service.OutputResponse{
		Stock: 1,
		Name:  []string{"hello", "world"},
	}
	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EmitDefaults: true,
		Indent:       "----",
	}
	for i := 0; i < n; i++ {
		err = marshaler.Marshal(jsonfile, resp)
		if err != nil {
			return fmt.Errorf("marshal pb data error %s", err)

		}
	}

	return nil
}
