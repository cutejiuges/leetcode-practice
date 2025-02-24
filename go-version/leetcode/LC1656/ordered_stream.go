package LC1656

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/24 下午11:08
 * @FilePath: ordered_stream
 * @Description: 设计有序流
 */

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n+1),
		ptr:    1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	list := make([]string, 0)
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		list = append(list, this.stream[this.ptr])
		this.ptr++
	}
	return list
}
