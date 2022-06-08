package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	var magic byte = 0x80
	var opcode byte = 0xb5
	var keyLength uint16 = 0x0000
	var extrasLength byte = 0x0
	var dataType byte = 0x00
	var vBucketId uint16 = 0x0000
	//var totalBodyLength uint32 = 0x0
	var opaque uint32 = 0x0f000000
	var CAS uint64 = 0x0000000000000000

	//var key string = "ID::1"
	//b_key := []byte(key)

	b_keyLength := make([]byte, 2)
	b_vBucketId := make([]byte, 2)

	byteRequest := make([]byte, 1024)
	byteRequest[0] = magic
	byteRequest[1] = opcode

	binary.LittleEndian.PutUint16(b_keyLength, keyLength)
	byteRequest[2] = b_keyLength[0]
	byteRequest[3] = b_keyLength[1]
	byteRequest[4] = extrasLength
	byteRequest[5] = dataType

	binary.LittleEndian.PutUint16(b_vBucketId, vBucketId)
	byteRequest[6] = b_vBucketId[0]
	byteRequest[7] = b_vBucketId[1]

	b_totalBodyLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(b_totalBodyLength, 0) //uint32(len(b_key))
	byteRequest[8] = b_totalBodyLength[0]
	byteRequest[9] = b_totalBodyLength[1]
	byteRequest[10] = b_totalBodyLength[2]
	byteRequest[11] = b_totalBodyLength[3]

	b_opaque := make([]byte, 4)
	binary.LittleEndian.PutUint32(b_opaque, opaque)
	byteRequest[12] = b_opaque[0]
	byteRequest[13] = b_opaque[1]
	byteRequest[14] = b_opaque[2]
	byteRequest[15] = b_opaque[3]

	b_cas := make([]byte, 8)
	binary.LittleEndian.PutUint64(b_cas, CAS)
	byteRequest[16] = b_cas[0]
	byteRequest[17] = b_cas[1]
	byteRequest[18] = b_cas[2]
	byteRequest[19] = b_cas[3]
	byteRequest[20] = b_cas[4]
	byteRequest[21] = b_cas[5]
	byteRequest[22] = b_cas[6]
	byteRequest[23] = b_cas[7]

	/*byteRequest[24] = b_key[0]
	byteRequest[25] = b_key[1]
	byteRequest[26] = b_key[2]
	byteRequest[27] = b_key[3]
	byteRequest[28] = b_key[4]*/

	dial, err := net.Dial("tcp4", "10.0.1.18:11210")
	if err != nil {
		panic(err)
	}

	write, err := dial.Write(byteRequest[0:28])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %v bytes", write)
	fmt.Println("")
	bytes := make([]byte, 1024)
	read, err := dial.Read(bytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %v bytes ", read)
	fmt.Println(string(bytes[0:read]))

}
