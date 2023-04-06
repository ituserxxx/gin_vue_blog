fn routes() -> impl Filter<Extract=impl warp::Reply, Error=warp::Rejection> + Clone {
    let users = warp::path!("users");

    let get_all_users = users.and(warp::get()).
        and_then(get_users);
    let get_single_user = users.and(warp::path!(i32)).
        and(warp::get()).
        and_then(get_user);
    let create_new_user = users.and(warp::post()).
        and(warp::body::json()).
        and_then(create_user);
    let update_existing_user = users.and(warp::put()).
        and(warp::path!(i32)).
        and(warp::body::json()).
        and_then(update_user);
    let delete_user_by_id = users.and(warp::delete()).
        and(warp::path!(i32)).
        and_then(delete_user);
    get_all_users
        .or(get_single_user)
        .or(create_new_user)
        .or(update_existing_user)
        .or(delete_user_by_id)
        .recover(handle_rejection);
}