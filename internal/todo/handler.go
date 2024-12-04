package todo

import "time"

type Todo struct{
	Body string `json:"body"`
	CreatedAt time.Time 
}

type Todos []Todo


func(t *Todos)AddTodo(body string){
	todo := Todo{
		Body: body,
		CreatedAt: time.Now(),
	}
	*t = append(*t, todo)
}

func(t *Todos)RemoveTodo(index int){
	arr := *t
	*t = append(arr[:index],arr[:index + 1]... )
}

func(t *Todos)GetList()*Todos{
	return &Todos{}
}