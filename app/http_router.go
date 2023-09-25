package app

import (
	"fmt"
	"net/http"
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	handler  http.HandlerFunc
}

type Router struct {
	root *TrieNode
}

func NewRouter() *Router {
	return &Router{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
			handler:  nil,
		},
	}
}

func (r *Router) AddRoute(path string, handler http.HandlerFunc) {
	node := r.root
	segments := splitPath(path)

	for _, segment := range segments {
		if _, exists := node.children[segment]; !exists {
			node.children[segment] = &TrieNode{
				children: map[string]*TrieNode{},
				handler:  nil,
			}
		}
		node = node.children[segment]
	}
	node.handler = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	node := r.root
	segments := splitPath(req.URL.Path)

	for _, segment := range segments {
		if child, exists := node.children[segment]; exists {
			node = child
		} else {
			http.NotFound(w, req)
			return
		}
	}

	if node.handler != nil {
		node.handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func splitPath(path string) []string {
	segments := make([]string, 0)
	for _, segment := range strings.Split(path, "/") {
		if segment != "" {
			segments = append(segments, segment)
		}
	}
	return segments
}

func Start(port string) {
	router := NewRouter()

	router.AddRoute("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	router.AddRoute("/user", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "User main page")
	})

	router.AddRoute("/user/profile", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "User Profile Page")
	})

	http.Handle("/", router)

	fmt.Printf("Application is listening on localhost%s", port)
	http.ListenAndServe(port, nil)
}
