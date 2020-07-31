package result

import "github.com/sofyan48/boilerplate/src/appctx/apps"

// Result ...
type Result interface {
	New(code int, name, msg string, data interface{}) *apps.Result
	List(data interface{}, meta interface{}) *apps.Result
	Detail(data interface{}, msg string) *apps.Result
	Message(name, msg string, details map[string]string) *apps.Result
	Accepted(name, msg string) *apps.Result
}
