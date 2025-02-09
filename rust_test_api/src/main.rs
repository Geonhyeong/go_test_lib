mod proto_mod;

use prost::Message;
use proto_mod::item;

use std::slice;
use tonic::Status;

#[link(name = "starforce.dll")]
extern "C" {
    fn StarForceWrapper(input_ptr: *const u8, input_len: usize, output_len: *mut usize) -> *mut u8;
}

fn main() {
    let item_event_sourcing = item::ItemEventSourcing {
        chuc: 1,
        option1: 3,
    };

    let item = item::Item {
        id: 1,
        value: "Park GwiYeon".to_string(),
        item_event_sourcing: Some(item_event_sourcing),
    };

    // prost를 이용해 Item을 바이트 배열로 직렬화
    let mut buf = Vec::new();
    item.encode(&mut buf).unwrap();

    let mut output_len: usize = 0;

    unsafe {
        let output_ptr = StarForceWrapper(buf.as_ptr(), buf.len(), &mut output_len as *mut usize);

        if output_ptr.is_null() {
            eprintln!("StarForceWrapper 호출 실패");
            return;
        }

        // 반환된 버퍼를 slice로 변환
        let output_slice = slice::from_raw_parts(output_ptr, output_len);

        // 반환된 바이트 배열을 역직렬화하여 새 Item 생성
        let new_item = item::Item::decode(output_slice).unwrap();

        // 반환된 데이터를 사용 (예시: 단순히 출력)
        println!("변경된 데이터: {:?}", new_item);
    }
}
