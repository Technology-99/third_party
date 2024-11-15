package bytesUtils

import (
	"fmt"
	"github.com/Technology-99/third_party/common"
)

func ten2sixteen2uint(param byte) uint {
	return common.Str2Uint(fmt.Sprintf("%x", param))
}

func splitBytes(data []byte, segmentSize int) [][]byte {
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
