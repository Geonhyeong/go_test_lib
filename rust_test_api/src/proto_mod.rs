// build.rs에서 compile()로 생성한 코드가 OUT_DIR에 들어있습니다.
pub mod item {
    include!(concat!(env!("OUT_DIR"), "/item.rs"));
}
