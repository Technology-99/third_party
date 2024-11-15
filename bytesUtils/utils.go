package bytesUtils

import (
	"fmt"
	"github.com/Technology-99/third_party/common"
	"sort"
	"strconv"
)

func Map2Bytes(data map[string]uint8) []byte {
	// note: 转换成buffer
	var result []byte
	keys := make([]int, 0, len(data))
	for k := range data {
		key, _ := strconv.Atoi(k) // 转换键为整数
		keys = append(keys, key)
	}
	sort.Ints(keys) // 按键升序排列

	for _, k := range keys {
		result = append(result, data[strconv.Itoa(k)]) // 将值转换为 byte 类型
	}
	return result
}

func Ten2sixteen2uint(param byte) uint {
	return common.Str2Uint(fmt.Sprintf("%x", param))
}

func SplitBytes(data []byte, segmentSize int) [][]byte {
	// 计算能分割的段数
	totalSegments := len(data) / segmentSize

	// 创建结果存储的二维切片
	result := make([][]byte, 0, totalSegments)

	// 按段切割数据
	for i := 0; i < totalSegments; i++ {
		start := i * segmentSize
		end := start + segmentSize
		result = append(result, data[start:end])
	}
	return result
}
