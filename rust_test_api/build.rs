fn main() -> Result<(), Box<dyn std::error::Error>> {
    // 라이브러리 파일이 위치한 디렉터리를 지정합니다.
    println!("cargo:rustc-link-search=native=.");
    // 동적 라이브러리(goitemlib)를 링크하도록 지정합니다.
    println!("cargo:rustc-link-lib=dylib=starforce");

    tonic_build::configure()
        // gRPC 서버와 클라이언트 코드를 모두 생성하려면 아래 옵션들을 사용합니다.
        .build_server(true)
        .build_client(true)
        // proto 파일들의 경로와 import 경로를 지정합니다.
        .compile_protos(
            &["../proto/item.proto"],
            &["../proto"], // proto 파일들을 찾을 경로
        )?;
    Ok(())
}
