package concurrency

// 内存“泄露” - 非平常意义的内存泄露, 在 GO 中是一种设计选择。
// 通过 GC 的方式，提高处理吞吐

// Buffer 缓冲类型
type Buffer struct {
	data []byte
	size int
}

// 创建一个大小可容纳 100 个指针的缓冲池
var freeList = make(chan *Buffer, 100)

// 服务端通道
var serverChan = make(chan *Buffer)

// 客户端
func client() {
	for {
		var b *Buffer

		// 获取缓存池
		select {
		// 如果
		case b = <-freeList: //
			// Got one; nothing more to do.
		default:
			// None free, so allocate a new one.
			b = new(Buffer)
		}

		load(b) // Read next message from the net.

		serverChan <- b // Send to server.
	}
}

// 模拟数据加载到 buffer 中
func load(b *Buffer) {
}

// 服务端
func server() {
	for {
		b := <-serverChan // Wait for work.
		process(b)
		// Reuse buffer if there's room.
		select {
		case freeList <- b:
			// Buffer on free list; nothing more to do.
		default:
			// Free list full, just carry on.
		}
	}
}

// 读取 buffer
func process(b *Buffer) {
	// 读取数据
	// reset buffer ，为了复用
}
