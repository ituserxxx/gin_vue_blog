mod router;

#[tokio::main]
async fn main() {
    let app = router::load_router();
    let address = "127.0.0.1:3019";
    println!("server listening on {}", address);
    axum::Server::bind(&address.parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();
}