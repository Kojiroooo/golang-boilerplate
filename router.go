package main

import (
	"app/controller"
	"app/domain/repository"
	"app/domain/service"
	"fmt"
	"net/http"
	"strings"
)

func HandleRequest() {
	http.HandleFunc("/posts", handlePostRequest)
	http.HandleFunc("/signup", handleAuthRequest)
	http.HandleFunc("/signin", handleAuthRequest)
	http.HandleFunc("/signout", handleAuthRequest)
	http.HandleFunc("/", handleRootRequest)
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if !checkAuthentication(w, r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "認証エラー\n")
		return
	}
	// inject
	controller := injectPostDependency()
	// routing
	switch r.Method {
	case "GET":
		controller.PostList(w, r)
	case "POST":
		fmt.Println("HandlePostRequest POST")
		controller.CreatePost(w, r)
	// case "PUT":
	// 	ro.postController.PutTodo(w, r)
	// case "DELETE":
	// 	ro.postController.DeleteTodo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "リクエストエラー\n")
	}
}

func injectPostDependency() controller.PostController {
	repository := repository.NewPostRepository()
	return controller.NewPostController(repository)
}

func handleAuthRequest(w http.ResponseWriter, r *http.Request) {
	// inject
	controller := injectAuthDependency()
	// routing
	switch r.Method {
	case "GET": // signin
		controller.Signin(w, r)
	case "POST": // signup
		fmt.Println("")
		controller.Signup(w, r)
	// case "PUT":
	// 	ro.postController.PutTodo(w, r)
	case "DELETE": // signout
		controller.Signout(w, r)
	default:
		w.WriteHeader(405)
	}
}

func injectAuthDependency() controller.AuthController {
	authRepository := repository.NewAuthRepository()
	userRepository := repository.NewUserRepository()
	service := service.NewAuthService(authRepository, userRepository)
	return controller.NewAuthController(service)
}

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("********Got root access********")
	fmt.Fprint(w, "Hello golang api!!")
}

// 本当はエラーを返す
func checkAuthentication(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	fmt.Println("********Check authorization********")
	fmt.Println(token)

	const TOKEN_PREFIX = "Bearer"
	if strings.Contains(token, TOKEN_PREFIX) {
		tmp := strings.Split(token, " ")
		if len(tmp) != 2 {
			return false
		}
		token = tmp[1]
	} else {
		return false
	}

	fmt.Println(token)
	authRepository := repository.NewAuthRepository()
	_, err := authRepository.Authorize(token)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
