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
Luego instalamos ANTL4 
```bash
mkdir -p ~/antlr
cd ~/antlr
curl -O https://www.antlr.org/download/antlr-4.13.1-complete.jar
```
Configuramos el entorno de bash 
```bash
nano ~/.bashrc
export CLASSPATH=".:$HOME/antlr/antlr-4.13.1-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M -cp "$HOME/antlr/antlr-4.13.1-complete.jar" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "$HOME/antlr/antlr-4.13.1-complete.jar" org.antlr.v4.gui.TestRig'
```
Luego ejecutamos el entorno de Bash 
```bash
source ~/.bashrc  # o ~/.zshrc
```
![image](https://github.com/user-attachments/assets/3a0ff98f-dc4e-41d4-96f2-0ce7aed4681a)

