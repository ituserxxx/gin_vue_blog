use axum::{
    Router,
};
mod user;
pub fn load_router() -> Router {
    return user::user_router();
}