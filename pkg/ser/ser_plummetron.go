package ser

import (
	"fmt"
	"log"
	"time"

	"net/http"

	"encoding/json"

	"go.bug.st/serial.v1"

	"github.com/leehayford/go_serial/pkg/dat"
)

var portName string = "COM3"

// var portName string = "/dev/ttyUSB0"

func RegisterPlummetronRoutes() {

	http.HandleFunc("/serial/ports", GetSerialPortsHandler)
	http.HandleFunc("/plummetron/tool-id", GetToolID)
	http.HandleFunc("/plummetron/sample", GetSample)
	http.HandleFunc("/plummetron/time", GetTime)
	http.HandleFunc("/plummetron/tool-info", Get64Bytes)

}

func GetSerialPortsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("\tGetting serial device list....")
	GetSerialPorts()
}

type SerialResponse struct {
	Type string
	Hex  string
	Data interface{}
}

func GetResponse(port serial.Port, buffsize int) SerialResponse {
	res := SerialResponse{}
	buff := make([]byte, buffsize)
	for {
		_, err := port.Read(buff)
		if err != nil {
			log.Println("Error reading serial port:\n", err)
			break
		}

		time.Sleep(1)
		log.Println("Response =>\n", fmt.Sprintf("%x", buff))
		start := dat.BytesToUInt16(buff[2:3])
		code := dat.BytesToUInt16(buff[3:4])
		log.Println("Received =>\n", fmt.Sprintf("Start: %d\tCode: %d", start, code))
		if start == 170 {
			if code == 48 { // Get Sample Response
				res.Type = "Sample"
				res.Data = DecodeSample(buff)
				break
			}

			if code == 64 { // Get Tool ID Response
				res.Type = "Tool ID"
				res.Data = DecodeToolID(buff)
				break
			}

			if code == 40 { // Get Time 0x28
				res.Type = "Get Time"
				// res.Data = DecodeToolInfo(buff)
				break
			}

			if code == 01 { // Read 64 Bytes 0x01
				res.Type = "Tool Information"
				res.Data = DecodeToolInfo(buff)
				break
			}

			if code == 00 { // Null Response 0x01
				res.Type = "Null Response"
				res.Data = fmt.Sprintf("%x", buff)
				break
			}
		}
	}
	res.Hex = fmt.Sprintf("%x", buff)
	log.Println("Response =>\n", res)
	return res
}

/*GET TOOL ID*/
type ToolID struct {
	ID    uint8 // [06:07]
	SubID uint8 // [11:12]

	MemCount    uint8 // [07:08]
	MemType     uint8 // [08:09]
	MemCapacity uint8 //[12:13]

	FWYear uint8 // [09:10]
	FWWeek uint8 // [10:11]

}

func DecodeToolID(buff []byte) ToolID {
	i := ToolID{}

	i.ID = buff[6:7][0]
	i.SubID = buff[11:12][0]

	i.MemCount = buff[7:8][0]
	i.MemType = buff[8:9][0]
	i.MemCapacity = buff[12:13][0]

	i.FWYear = buff[9:10][0]
	i.FWWeek = buff[10:11][0]

	return i
}

func GetToolID(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Tool ID...")
	arr := []byte{0xAA, 0xAA, 0xAA, 0x55, 0x40, 0x00, 0x01, 0xEE, 0xFA} // Get Tool Info
	port := OpenSerialPort(portName)

	_, err := port.Write(arr)
	if err != nil {
		log.Fatal("Error writing to serial port\n", err)
	} // log.Println("Bytes written:\t", arr)

	res := GetResponse(port, 15)

	port.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/*GET TOOL INFORMATION*/
type ToolInfo struct {
	SerialNumber string // [06:26]

	MaxPressure  string // [26:45]
	PressuerUnit uint8  // [45:46]

	MaxTemp  string // [46:62]
	TempUnit uint8  //[65:66]

	DateFormat     uint8 // [62:63]
	TimeColumnUnit uint8 // [63:64]
	TimeUnit       uint8 // [64:65]
}

func DecodeToolInfo(buff []byte) ToolInfo {
	info := ToolInfo{}

	info.SerialNumber = string(dat.StripFFs(buff[6:26]))

	info.MaxPressure = string(dat.StripFFs(buff[26:45]))
	info.PressuerUnit = buff[45:46][0]

	info.MaxTemp = string(dat.StripFFs(buff[46:62]))
	info.TempUnit = buff[65:66][0]

	info.DateFormat = buff[62:63][0]
	info.TimeColumnUnit = buff[63:64][0]
	info.TimeUnit = buff[64:65][0]
	return info
}

func Get64Bytes(w http.ResponseWriter, r *http.Request) {

	arr := []byte{0xAA, 0xAA, 0xAA, 0x55, 0x02, 0x00, 0x04, 0x00, 0x0A, 0x00, 0x00, 0x59} // Get 64 Bytes Starting at 2560
	port := OpenSerialPort(portName)

	_, err := port.Write(arr)
	if err != nil {
		log.Fatal("Error writing to serial port\n", err)
	} // log.Println("Bytes written:\t", arr)

	res := GetResponse(port, 70)

	port.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/*GET RAW SAMPLE*/
type RawSample struct {
	Pressure    uint32
	Temperature uint32
	AccX        uint32
	AccY        uint32
	AccZ        uint32
}

func DecodeSample(buff []byte) RawSample {
	log.Println("Decoding sample...")
	s := RawSample{}
	s.Pressure = dat.BytesToUInt32(buff[6:9])
	s.Temperature = dat.BytesToUInt32(buff[9:12])
	s.AccZ = dat.BytesToUInt32(buff[12:15])
	s.AccY = dat.BytesToUInt32(buff[15:18])
	s.AccX = dat.BytesToUInt32(buff[18:21])
	return s
}

func GetSample(w http.ResponseWriter, r *http.Request) {

	arr := []byte{0xAA, 0xAA, 0xAA, 0x55, 0x30, 0x00, 0x01, 0xEE, 0x8A} // Get Sample
	port := OpenSerialPort(portName)

	_, err := port.Write(arr)
	if err != nil {
		log.Fatal("Error writing to serial port\n", err)
	} // log.Println("Bytes written:\t", arr)

	res := GetResponse(port, 23)

	port.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

/*** NOT IMPLEMENTED ***/
func GetTime(w http.ResponseWriter, r *http.Request) {
	arr := []byte{0xAA, 0xAA, 0xAA, 0x55, 0x28, 0x00, 0x01, 0xAA, 0xD6} // Get Time
	port := OpenSerialPort(portName)

	_, err := port.Write(arr)
	if err != nil {
		log.Fatal("Error writing to serial port\n", err)
	} // log.Println("Bytes written:\t", arr)

	res := GetResponse(port, 22)

	port.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
