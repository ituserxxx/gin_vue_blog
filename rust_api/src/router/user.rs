use axum::{
    Router,
    routing::post,
};

use crate::mod_rs::service::user as user;

pub fn user_router() -> Router {
    return Router::new()
        .route("/", post(service::user::hello))
        .route("/users", post(service::user::create_user));
}

