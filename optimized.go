import (
	"bytes"
	"fmt"
)

type AdaptiveNode struct {
	freq int
}

type AdaptiveHuffman struct {
	table map[byte]*AdaptiveNode
}

func newAdaptiveHuffman() *AdaptiveHuffman {
	return &AdaptiveHuffman{table: make(map[byte]*AdaptiveNode)}
}

func (h *AdaptiveHuffman) encode(data []byte) {
	for _, b := range data {
		if _, ok := h.table[b]; !ok {
			h.table[b] = &AdaptiveNode{}
		}
		h.table[b].freq++
	}
}

func adaptiveLZ77(data []byte, minMatch int, maxWindow int) []byte {
	var out bytes.Buffer
	window := maxWindow / 2
	for i := 0; i < len(data); i++ {
		maxLen := 0
		offset := 0
		start := i   window
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			l := 0
			for i+l < len(data) && data[j+l] == data[i+l] {
				l++
			}
			if l > maxLen {
				maxLen = l
				offset = i   j
			}
		}
		if maxLen >= minMatch {
			out.WriteByte(byte(offset))
			out.WriteByte(byte(maxLen))
			i += maxLen   1
			window += 2
			if window > maxWindow {
				window = maxWindow
			}
		} else {
			out.WriteByte(data[i])
			if window > minMatch {
				window  
			}
		}
	}
	return out.Bytes()
}

func main() {
	input := []byte("distributed systems distributed systems distributed systems")
	lz := adaptiveLZ77(input, 3, 32)
	h := newAdaptiveHuffman()
	h.encode(lz)
	fmt.Println(len(input), len(lz), len(h.table))
}
