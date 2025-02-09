package main

import (
	"C"
	"unsafe"

	myItem "github.com/Geonhyeong/go_test_lib/proto"
	"google.golang.org/protobuf/proto"
)
import v4 "github.com/Geonhyeong/go_test_lib/v4"

//export StarForceWrapper
func StarForceWrapper(input_ptr *C.uchar, input_len C.size_t, output_len *C.size_t) *C.uchar {
	// input_ptr에서 바이트 슬라이스 생성
	input := C.GoBytes(unsafe.Pointer(input_ptr), C.int(input_len))

	// proto 메시지 역직렬화 (오류 처리는 생략)
	item := &myItem.Item{}
	proto.Unmarshal(input, item)

	// item 값 변경 로직 (외부 라이브러리의 핵심 로직)
	v4.StarForce(item)

	// 변경된 item 직렬화
	out, _ := proto.Marshal(item)
	*output_len = C.size_t(len(out))

	// C에서 free()로 해제할 수 있는 메모리 할당
	return (*C.uchar)(C.CBytes(out))
}

func main() {
	// test_item := &myItem.Item{}
	// test_item.Id = 1
	// test_item.Value = "Park Gwiyeon"

	// test_ItemEventSourcing := &myItem.ItemEventSourcing{}
	// test_ItemEventSourcing.CHUC = 3
	// test_ItemEventSourcing.Option1 = 5

	// test_item.ItemEventSourcing = test_ItemEventSourcing

	// item_encode, _ := proto.Marshal(test_item)
	// len := C.size_t(len(item_encode))
	// output_len := C.size_t(0)

	// output_ptr := StarForceWrapper((*C.uchar)(C.CBytes(item_encode)), len, &output_len)
	// item := &myItem.Item{}

	// out := C.GoBytes(unsafe.Pointer(output_ptr), C.int(output_len))
	// proto.Unmarshal(out, item)

	// print(item)
}
