# Instalacion Completa de ANTLR4 Tool y Runtime para GO
## Errores comunes
Si estás leyendo esto es muy probable que ya hayas instalado la versión 
con un comando tipo `curl -O antlr4-complete.jar` y eso no te sirvio. 
El motivo es porque en los ultimos años cambiaron la dirección del repo
de antlr, lo movieron de carpeta (lol). Así que aquí te muestro como arreglar tu errorsito: 

Primero debemos buscar todas las instancias de el repositorio antiguo y actualizarlas 
```bash
find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/antlr/antlr4/runtime/Go/antlr/v4,github.com/antlr4-go/antlr/v4,g' {} \;
```
o bien en una linea como, lo debes correr en la carpeta de los archivos generados. 
```bash
find . -type f -name '*.go' -exec sed -i -e 's,github.com/antlr/antlr4/runtime/Go/antlr/v4,github.com/antlr4-go/antlr/v4,g' {} \;
```

## Instalacion de Go 1.21 
NOTA: 
Estás intentando usar una versión de ANTLR Go Runtime que requiere como mínimo Go 1.21. Aqui hay un tutorial de como instalar Go 1.21. 
```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
```
Luego añadimos al CLASSPATH 
```bash
export PATH=$PATH:/usr/local/go/bin
```
![image](https://github.com/user-attachments/assets/f2f0d9fe-4881-4371-87b5-2816a14b513d)
Actualizamos la consola
```bash 
source ~/.bashrc
```
Luego le damos permiso al
```bash
chmod +x generate.sh
```

## Instalacion de de ANTLR4 con Go 
```bash
go get -u github.com/antlr4-go/antlr
```
![image](https://github.com/user-attachments/assets/a1610ac3-a16f-4789-94ff-b271029f2821)

La estructura de nuestro proyecto debería ser: 
```bash
	.
	├── myproject
	├── parser
	│     ├── mygrammar.g4
	│     ├── antlr-4.13.2-complete.jar
	│     ├── generate.go
	│     └── generate.sh
	├── parsing  # Generated code goes here
	│     └── error_listeners.go
	├── go.mod
	├── go.sum
	├── main.go
	└── main_test.go
```

Si aun no has hecho la instalacion, preferiblemente inicia con esto: 
```bash
go mod init
```
te va a generar un go.mod en donde se encuentran las dependencias del proyecto
entra a el y si antes usaste go, veras algo asi: 
```yaml
module compiler

go 1.18

require (
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10
	github.com/antlr4-go/antlr/v4 v4.13.1
)

require (
	github.com/antlr4-go/antlr v0.0.0-20230518091524-98b52378c522 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
)
```
Especificamente `github.com/antlr/antlr4/runtime/Go/antlr v1.4.10` es la version antigua, asi que se deben asegurar que NO sea el utilizado. La version actual es 
`github.com/antlr4-go/antlr/v4`. Usar ambos puede dar conflictos. 


`go mod tidy` es un comando de Go que limpia y actualiza tu archivo go.mod y go.sum, haciendo que reflejen exactamente lo que tu proyecto realmente necesita.

## Instalar ANTLR4 Tool 
```bash 
pip install antlr4-tools
```
## Generacion de Archivos Facil
Actualmente tengo esta division de archivos 
```bash
.
├── Clase1
│   └── compiler
│       ├── evalVisitor.go
│       ├── generate.go
│       ├── generate.sh
│       ├── go.mod
│       ├── go.sum
│       ├── main.go
│       ├── V4LangGrammar.g4
│       └── V4LangLexer.g4
├── go1.21.5.linux-amd64.tar.gz
├── Instalar_ANTLR4.md
└── README.md

```
Si lo notan, tengo V4LangGrammar y V4LangLexer en esta seccion. Correr el comando de `antlr DLanguage=Go` se vuelve exhaustivo. Así que haremos un archivo `generate.go` y un `generate.sh` para que solo debamos correr `go generate` para generar nuestros archivos. 
En `go generate` escribimos: 

`chmod +x generate.sh`

## Notas importantes
![image](https://github.com/user-attachments/assets/6adc625b-8a0c-4502-8573-d7ecd1f49540)

EL REPOSITORIO QUE SE DEBE USAR ES ESTE: `github.com/antlr4-go/antlr`



```bash 
go get github.com/antlr/antlr4/runtime/Go/antlr@latest
```
```bash 
antlr4 -Dlanguage=Go -visitor -package parser -o parser V4LangGrammar.g4
antlr4 -Dlanguage=Go -visitor -package parser -o parser V4LangGrammar.g4
```
