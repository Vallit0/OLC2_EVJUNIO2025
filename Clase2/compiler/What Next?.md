# What Next? 
Seguramente te estas preguntando que debes hacer. 
En primera instancia tenemos que hacer nuestras expresiones, 
definimos una gramatica simple. 

```bash 
Escribe el codigo de Vlang: (Ctrl+D) para terminar
3 + 5 * (2 - 1)
Se termino el analisis sintactico
+-- programa
    |-- expresion
    |   |-- expresion
    |   |   +-- valores
    |   |       +-- valor
    |   |           +-- "3"
    |   |-- "+"
    |   +-- expresion
    |       |-- expresion
    |       |   +-- valores
    |       |       +-- valor
    |       |           +-- "5"
    |       |-- "*"
    |       +-- expresion
    |           |-- "("
    |           |-- expresion
    |           |   |-- expresion
    |           |   |   +-- valores
    |           |   |       +-- valor
    |           |   |           +-- "2"
    |           |   |-- "-"
    |           |   +-- expresion
    |           |       +-- valores
    |           |           +-- valor
    |           |               +-- "1"
    |           +-- ")"
    +-- "<EOF>"
```
Agregue un metodo para generar este arbol cada vez que probamos la estructura.
## Hagamos la declaraciones de las variables 
Nota: Si no le ponemos las labels a los hermanos, no podemos generar nada en la gramatica. 



## Se que existe el Visitor pero...¿Ahora qué? 

Lo primero es que debemos hacer un visitor por las operaciones que queremos hacer
en este caso es interpretar. Por lo que, a partir de esta gramatica, si nos vamos a la carpeta `/parser` notaremos que en `vlang_visitor.go` y en  `base_visitor.go` 
estaran los metodos que nos proponen la interfaz del Visitor y la clase abstracta.
Debes copiar estos métodos a un archivo en el que quieras para declarar un visitor 
que extienda de la interfaz de `visitor` que propones. 

### Creacion de un visitor 
En la carpeta `repl` 
podemos encontrar estos archivos
```Bash 
├── argument.go
├── console.go
├── error_table.go
├── repl.go
├── scope.go
├── variable.go
└── visitor.go
```
de momento utilizaremos `visitor.go` para crear nuestro visitor. 
```mermaid 
```mermaid
classDiagram
    class BaseVlangVisitor {
        +Visit(tree ParseTree) : interface{}
        +VisitChildren(node RuleNode) : interface{}
    }

    class ReplVisitor {
        +VisitPrograma(ctx ProgramaContext) : interface{}
        +Visit<No-Terminal>(ctx <No-Terminal>) : interface{}
        -errorTable : ErrorTable
        -console : Console
    }

    BaseVlangVisitor <|-- ReplVisitor

    class ProgramaContext
    class SumresContext
    class ValorexprContext
    class PrintStatementContext
    class BaseScope
    class ErrorTable
    class Console

    ReplVisitor --> ProgramaContext : visita
    ReplVisitor --> SumresContext : visita
    ReplVisitor --> ValorexprContext : visita
    ReplVisitor --> PrintStatementContext : visita
    ReplVisitor --> BaseScope : mantiene
    ReplVisitor --> ErrorTable : reporta a
    ReplVisitor --> Console : imprime en
```