# Instalacion Completa de ANTLR4 y Primer Ejemplo
```bash
curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
```
```bash
export CLASSPATH=".:antlr-4.13.1-complete.jar:$CLASSPATH"
alias antlr4='java -jar antlr-4.13.1-complete.jar'
alias grun='java org.antlr.v4.gui.TestRig'
```
```bash
sudo apt install golang-go
```
Instalamos Go
```bash
go version
```
```bash
go install github.com/antlr/antlr4/runtime/Go/antlr@latest
```
Tambien necesitamos instalar Java 
```bash
sudo apt update
sudo apt install default-jre -y
```

```bash
go mod init mi_compilador
go mod tidy
```
