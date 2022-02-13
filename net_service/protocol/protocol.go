package protocol

// 协议接口
type Protocol interface {
  // 计算数据长度
  DataLength(data []byte) (dataLength int, err error)
  // 数据解码
  DataDecode(data []byte) (decodeData []byte, err error)
  // 数据编码
  DataEncode(data []byte) (encodeData []byte, err error)
}
