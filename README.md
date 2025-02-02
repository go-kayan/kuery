# Kuery - Construtor de Queries SQL em Go

`Kuery` é uma biblioteca simples e fluida para construir queries SQL em Go de forma dinâmica e segura.  
Ela suporta comandos `SELECT`, `INSERT`, `UPDATE` e `DELETE` com uma API intuitiva baseada em encadeamento de métodos.

---

## 📦 Instalação

Para adicionar a biblioteca ao seu projeto, execute:

```sh
go get github.com/go-kayan/kuery
```
---

## 📌 Como Usar
### 🔹 **Insert**

```go
package main

import (
	"fmt"
	"github.com/go-kayan/kuery"
)

func main() {
	query := kuery.Insert("usuarios").
		Columns("nome", "email").
		Values("João", "joao@email.com").
		Values("Maria", "maria@email.com").
		Returning("id", "nome").
		Build()

	fmt.Println(query)
}

//INSERT INTO usuarios (nome, email) VALUES ('João', 'joao@email.com'), ('Maria', 'maria@email.com') RETURNING id, nome
```

### 🔹 **Delete**

```go
package main

import (
	"fmt"
	"kuery"
)

func main() {
	query := kuery.Delete("usuarios").
		Where("id = 1").
		Build()

	fmt.Println(query)
}

// DELETE FROM usuarios WHERE id = 1
```
### 🔹 **Update**
```go
package main

import (
	"fmt"
	"kuery"
)

func main() {
	query := kuery.Update("usuarios").
		Set(map[string]interface{}{
			"nome":  "João Silva",
			"email": "joaosilva@email.com",
		}).
		Where("id = 1").
		Build()

	fmt.Println(query)
}
// UPDATE usuarios SET nome = 'João', email = 'joao@email.com' WHERE id = 1 
```
### 🔹 **Select**

```go
package main

import (
	"fmt"
	"kuery"
)

func main() {
	query := kuery.Select("id", "nome").
		From("usuarios").
		Where("idade > 18").
		And("ativo = true").
		OrderBy("nome ASC").
		Limit(10).
		Build()

	fmt.Println(query)
}
// SELECT id, nome FROM usuarios WHERE idade > 18 AND ativo = true ORDER BY nome ASC LIMIT 10
```

