package endpoint

import "todo/usecase"

type TodoEndpoint interface {
}

type todoEndpoint struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoEndpoint(todoUseCase usecase.TodoUseCase) TodoEndpoint {
	return &todoEndpoint{
		todoUseCase: todoUseCase,
	}
}
