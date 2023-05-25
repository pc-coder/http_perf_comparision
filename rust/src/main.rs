use warp::{
    Filter,
    http::{Response},
};

#[tokio::main]
async fn main() {
    pretty_env_logger::init();

    let health = warp::get()
        .and(warp::path("health"))
        .map(|| {
            Response::builder().body("{ status = up }")
        });

    warp::serve(health)
        .run(([127, 0, 0, 1], 8000))
        .await
}