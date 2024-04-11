package serialization

import (
	"bytes"
	"compression-tool/priorityqueue"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"strings"
)




func  WriteEncodedStringToFile(outputFilePath string, encodedString string ) error {
	file, err := os.Create(outputFilePath)
	if err!=nil {
		fmt.Println("error creating file")
		return err
	}
	defer file.Close()

	var byteBuffer bytes.Buffer
	for i:=0;i<len(encodedString);i+=8 {
		var byteValue byte
		for j:=0;j<8 && i+j<len(encodedString); j++ {
			if encodedString[i+j]=='1' {
				byteValue |= 1 << (7-j)
			}
		}
		byteBuffer.WriteByte(byteValue)
	}

	_, err = file.Write(byteBuffer.Bytes())
	if err!= nil {
		return err
	}
	return nil
}

func ReadEncoededStringFromFile(filepath string) (string, error) {
	file, err:= os.Open(filepath)
	if err!=nil {
		return "",err
	}
	defer file.Close()

	var encodedSringBuilder strings.Builder
	bytevalue:= make([]byte,1)

	for {
		_, err:= file.Read(bytevalue)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		for i:=7;i>=0;i-- {
			if bytevalue[0]&(1<<i) == 0{
				encodedSringBuilder.WriteRune('0')
			} else {
				encodedSringBuilder.WriteRune('1')
			}
		}
	}
	encodedString:= encodedSringBuilder.String()
	// remove padding of zeros
	return encodedString,nil

}

func SerelizeTreeToFile(node *priorityqueue.HuffmanNode, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(node)
	if err != nil {
		return err
	}

	return nil
}

func DeserializeTreeFromFile(filepath string) (*priorityqueue.HuffmanNode,error) {
	file, err := os.Open(filepath)
	if err!= nil{
		return nil,err
	}
	defer file.Close()

	var node priorityqueue.HuffmanNode
	dec := gob.NewDecoder(file)
	err = dec.Decode(&node)
	if err != nil {
		return nil, err
	}
	return &node, nil
}