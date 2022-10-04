package controller

type UserController interface {
	SignIn()
	SignUp()
	SignOut()
}

// UserControllerは、自分自身のものではない
// AuthController
// SignUp
// POST /signup
// request email, password
// response token, user_id
// ユーザー作成とトークンの発行
// SignIn
// GET /signin
// request email, password
// response token, user_id
// 認証とトークンの発行
// SignOut
// DELETE /signout
// request なし
// response なし
// ProfileController
//

// Auth
// user_id
// token
// expired_at
