package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/zdecoder"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/xxl6097/clink-go-tcp-base-lib/iface/impl"
	"github.com/xxl6097/clink-go-tcp-base-lib/server"
	"github.com/xxl6097/clink-go-tcp-base-lib/util/httputil"
	"github.com/xxl6097/go-glog/glog"
	"math"
	"time"
)

type TestTimer struct {
}

func (this *TestTimer) Run() {
	glog.Debug("test Run..")
}

func test() {
	tt := &TestTimer{}
	timer := impl.New(time.Second*2, tt.Run)
	timer.Start()
}

func main() {
	test()
	httputil.Get("http://baidu.com", nil, time.Second)
	fmt.Println("hello main..")
	server.ServeTCP(func(begin bool, conn ziface.IConnection) {
		glog.Info("Connection ...", begin, conn.GetConnID(), conn.RemoteAddr())
	}, func(conf *zconf.Config) {
		conf.Name = "main"
		conf.Host = "0.0.0.0"
		conf.TCPPort = 8899
		conf.MaxConn = 10000
		conf.WorkerPoolSize = 10
	}, func(f func(ziface.IDecoder)) {
		f(NewDecoder())
	}, func(f func(uint32, ziface.IRouter)) {
		f(0x15, &Message0x15{})
	})
}

// //////////////////////////////
const HEADER_SIZE = 6

type TestDecoder struct {
	Head    byte   //头码
	Funcode byte   //功能码
	Length  uint16 //数据长度
	Body    []byte //数据内容
	Crc     []byte //CRC校验
	Data    []byte //原始数据内容
}

func NewDecoder() ziface.IDecoder {
	return &TestDecoder{}
}

func (hcd *TestDecoder) GetLengthField() *ziface.LengthField {
	//+------+-------+---------+--------+--------+
	//| 头码  | 功能码 | 数据长度 | 数据内容 | CRC校验 |
	//| 1字节 | 1字节  | 1字节   | N字节   |  2字节  |
	//+------+-------+---------+--------+--------+
	//头码   功能码 数据长度      Body                         CRC
	//A2      10     0E        0102030405060708091011121314 050B
	//说明：
	//   1.数据长度len是14(0E),这里的len仅仅指Body长度;
	//
	//   lengthFieldOffset   = 2   (len的索引下标是2，下标从0开始) 长度字段的偏差
	//   lengthFieldLength   = 1   (len是1个byte) 长度字段占的字节数
	//   lengthAdjustment    = 2   (len只表示Body长度，程序只会读取len个字节就结束，但是CRC还有2byte没读呢，所以为2)
	//   initialBytesToStrip = 0   (这个0表示完整的协议内容，如果不想要A2，那么这里就是1) 从解码帧中第一次去除的字节数
	//   maxFrameLength      = 255 + 4(起始码、功能码、CRC) (len是1个byte，所以最大长度是无符号1个byte的最大值)
	return &ziface.LengthField{
		MaxFrameLength:      math.MaxUint16 + 4,
		LengthFieldOffset:   2,
		LengthFieldLength:   2,
		LengthAdjustment:    2,
		InitialBytesToStrip: 0,
	}
}

func (hcd *TestDecoder) decode(data []byte) *TestDecoder {
	datasize := len(data)

	htlvData := TestDecoder{
		Data: data,
	}

	//4. 解析头
	htlvData.Head = data[0]
	htlvData.Funcode = data[1]

	var value uint16
	buffer := bytes.NewBuffer(data[2:4])
	binary.Read(buffer, binary.BigEndian, &value)
	htlvData.Length = value

	htlvData.Body = data[4 : datasize-2]
	htlvData.Crc = data[datasize-2 : datasize]

	//5. CRC校验
	if !zdecoder.CheckCRC(data[:datasize-2], htlvData.Crc) {
		glog.Debugf("crc校验失败 %s %s\n", hex.EncodeToString(data), hex.EncodeToString(htlvData.Crc))
		return nil
	}

	//glog.Debugf("2htlvData %s \n", hex.EncodeToString(htlvData.data))
	//glog.Debugf("HTLVCRC-DecodeData size:%d data:%+v\n", unsafe.Sizeof(htlvData), htlvData)

	return &htlvData
}

func (hcd *TestDecoder) Intercept(chain ziface.IChain) ziface.IcResp {
	//1. 获取zinx的IMessage
	iMessage := chain.GetIMessage()
	if iMessage == nil {
		//进入责任链下一层
		return chain.ProceedWithIMessage(iMessage, nil)
	}

	//2. 获取数据
	data := iMessage.GetData()
	//glog.Debugf("HTLVCRC-RawData size:%d data:%s\n", len(data), hex.EncodeToString(data))

	//3. 读取的数据不超过包头，直接进入下一层
	if len(data) < HEADER_SIZE {
		return chain.ProceedWithIMessage(iMessage, nil)
	}

	//4. HTLV+CRC 解码
	htlvData := hcd.decode(data)

	//5. 将解码后的数据重新设置到IMessage中, Zinx的Router需要MsgID来寻址
	iMessage.SetMsgID(uint32(htlvData.Funcode))

	//6. 将解码后的数据进入下一层
	return chain.ProceedWithIMessage(iMessage, *htlvData)
}

/////////////////

type Message0x15 struct {
	znet.BaseRouter
}

type JsonData struct {
	Ciphertext string `json:"ciphertext"`
}

func (this *Message0x15) Handle(request ziface.IRequest) {
	_response := request.GetResponse()
	if _response != nil {
		switch _response.(type) {
		case TestDecoder:
			_data := _response.(TestDecoder)
			glog.Debugf("body:%s\n", string(_data.Body))

			var res JsonData
			errs := json.Unmarshal(_data.Body, &res)
			if errs != nil {
				glog.Debugf("json unmarshal error:%s", errs.Error())
			}

			decoded, _ := base64.StdEncoding.DecodeString(res.Ciphertext)

			glog.Debugf("ciphertext:%s\n", decoded)
			//glog.Debugf("do Data0x15Router data business %+v\n", _data)
			buffer := pack15(_data)
			request.GetConnection().Send(buffer)
		}
	} else {
		glog.Debugf("Message0x15 Handle %s \n", hex.EncodeToString(request.GetMessage().GetData()))
	}
}

// 头码   功能码 数据长度      Body                         CRC
// A2      10   000E        0102030405060708091011121314 050B
func pack15(_data TestDecoder) []byte {
	buffer := bytes.NewBuffer([]byte{})
	buffer.WriteByte(0xA1)
	buffer.WriteByte(_data.Funcode)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, _data.Length)

	buffer.Write(bytesBuffer.Bytes())
	buffer.Write(_data.Body)
	return buffer.Bytes()
}
