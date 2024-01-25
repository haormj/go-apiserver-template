package idl

type AddUserInput struct {
	Input
	Data AddUserInputData `json:"data"`
}

type AddUserInputData struct {
	Name string `json:"name" validate:"required,omitempty"`
}

type AddUserOutput struct {
	Output
}

type UpdateUserInput struct {
	Input
	Data UpdateUserInputData `json:"data"`
}

type UpdateUserInputData struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name"`
}

type UpdateUserOutput struct {
	Output
}

type ListUsersInput struct {
	Input
	Data ListUsersInputData `json:"data"`
}

type ListUsersInputData struct {
	Name *string `json:"name"`
}

type ListModsOutput struct {
	Output
	Data []ListModsOutputData `json:"data"`
}

type ListModsOutputData struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	InsertTime string `json:"insert_time"`
	UpdateTime string `json:"update_time"`
}
