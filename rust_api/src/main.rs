use std::convert::Infallible;
use warp::{Filter, Reply};

#[tokio::main]
async fn main() {
    let hello = warp::path!("hello" / String)
        .map(|name| format!("Hello, {}!", name));

    let goodbye = warp::path!("goodbye" / String)
        .map(|name| format!("Goodbye, {}!", name));
    let routes = hello.or(goodbye);
    warp::serve(routes).run(([127, 0, 0, 1], 8000)).await;
}
