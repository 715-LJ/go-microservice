syntax = "v1"

info (
	title: "title"
	desc: "desc"
	author: "{{.gitUser}}"
	email: "{{.gitEmail}}"
)

type request {
	// TODO: add members here and delete this comment
}

type response {
	// TODO: add members here and delete this comment
}

service {{.serviceName}} {
	@handler GetUser // TODO: set handler name and delete this comment
	get /users/id/:userId(request) returns(response)

	@handler CreateUser // TODO: set handler name and delete this comment
	post /users/create(request)
}
