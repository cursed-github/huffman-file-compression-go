package main

import (
	"bufio"
	"compression-tool/priorityqueue"
	"compression-tool/serialization-tree"
	"fmt"
	"log"
	"os"
	"strings"
)

func logger(args ...interface{}) {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	log.SetOutput(file)

	// Create a message from all arguments
	message := ""
	for _, arg := range args {
		message += fmt.Sprintf("%v ", arg) // Convert each argument to string and concatenate
	}
	log.Println(message)
}

type HoffmanEncryption struct {
	filepath string
	EncodedFile string
	DecodedFile string
	FrequncyMap map[rune]*priorityqueue.HuffmanNode
	pq priorityqueue.PriorityQueue
	node *priorityqueue.HuffmanNode
}

func (hf *HoffmanEncryption) readFile() {
	hf.FrequncyMap = make(map[rune]*priorityqueue.HuffmanNode)
	file, err := os.Open(hf.filepath)

	if err != nil {
		//log error brotha
		logger(err)
	}
	defer file.Close()
	reader:= bufio.NewReader(file)

	for {
		r,_,err := reader.ReadRune()
		if err!= nil {
			break
		}

		if node, exists := hf.FrequncyMap[r]; exists {
			node.Frequency++
		} else {
			hf.FrequncyMap[r] = &priorityqueue.HuffmanNode{
				Char: r,
				Frequency: 1,
			}
		}
	}
}

func (hf *HoffmanEncryption) InitQueue() {
	hf.pq = make(priorityqueue.PriorityQueue,0)
	hf.pq.Init()

	for _ ,node := range hf.FrequncyMap {
		hf.pq.Push(node)
	}

	for hf.pq.Len() > 1 {
		left:= hf.pq.ExtractMinimum()
		right:= hf.pq.ExtractMinimum()

		hf.pq.Push(&priorityqueue.HuffmanNode{
			Left: left,
			Right: right,
			Frequency: left.Frequency + right.Frequency,

		})
	}
}

func(hf *HoffmanEncryption) EncodeChars(node *priorityqueue.HuffmanNode, prefix string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		node.CharEncoding = prefix
	}
	hf.EncodeChars(node.Left,prefix+"0")
	hf.EncodeChars(node.Right,prefix+"1")
}

func (hf *HoffmanEncryption) EncodeFile() {
	file, err := os.Open(hf.filepath)
	var encodedBuilder strings.Builder
	if err != nil {
		//log error brotha
		logger(err)
	}
	defer file.Close()
	reader:= bufio.NewReader(file)

	for {
		r,_,err := reader.ReadRune()
		if err!= nil {
			break
		}

		//hf.encryptedFile+=hf.FrequncyMap[r].CharEncoding
		encodedBuilder.WriteString(hf.FrequncyMap[r].CharEncoding)
	
		
	}
	hf.EncodedFile = encodedBuilder.String()
	//logger(hf.EncodedFile)
}




func (hf *HoffmanEncryption) DecodeFile(node *priorityqueue.HuffmanNode) {
	var decodedBuilder strings.Builder
	root := node // Start with the root of the Huffman tree

	for _, bit := range hf.EncodedFile {
		if bit == '0' {
			node = node.Left  // Move left for '0'
		} else if bit == '1' {
			node = node.Right  // Move right for '1'
		}

		if node.Left == nil && node.Right == nil {  // Check if it's a leaf node
			decodedBuilder.WriteRune(node.Char)  // Append the character of the leaf node
			node = root // Reset to root for the next character
		}
	}

	hf.DecodedFile = decodedBuilder.String()
	logger("Decoded content:", hf.DecodedFile)
}

func (hf *HoffmanEncryption) SerializeToFile() {
	err := serialization.SerelizeTreeToFile(hf.node, "output-tree.bin")
	if err != nil {
		logger("Error serializing Huffman Tree:", err)
	}

	// Write Encoded String to File
	err = serialization.WriteEncodedStringToFile( "encoded-file.bin",hf.EncodedFile,)
	if err != nil {
		logger("Error writing encoded string to file:", err)
	}
}

func (hf *HoffmanEncryption) DeserializeFromFile() {
	node, err := serialization.DeserializeTreeFromFile("output-tree.bin")
	if err != nil {
		logger("Error deserializing Huffman Tree:", err)
		return
	}
	hf.node = node
	
	encodedString, err := serialization.ReadEncoededStringFromFile("encoded-file.bin")
	if err != nil {
		logger("Error reading encoded string from file:", err)
		return
	}
	hf.EncodedFile = encodedString
}


func main() {
	hf:= &HoffmanEncryption{
		filepath:"book.txt",
	}
	hf.readFile()
	hf.InitQueue()
	hf.node = hf.pq.ExtractMinimum()
	hf.EncodeChars(hf.node,"")
	fmt.Println(hf.FrequncyMap[' '])
	hf.EncodeFile()
	

	
	
	//logger(node, encodedString)
	hf.DecodeFile(hf.node)
}