# Projeto MCONF

Seja muito bem-vindo ao Projeto MCONF



<br>

> Projeto constituí em contruir uma api em GO que consulta uma api externa e traz o resultado com uma leitura facilitada para humanos e rodar ela em um container no Docker.
<br>


```go
go run server.go
```

<img src="image/comand-vscode.png" alt="exemplo imagem">


```cmd
localhost:3000
```

<img src="image/pagina_renderizada.png" alt="exemplo imagem">

<br>
<br>

## 💻 Pré-Requisitos
<br>
Antes de começar, verifique se você atendeu aos seguintes requisitos:

* Você instalou a versão mais recente de `Go`
* Você tem uma máquina `Linux/Windows/Mac`. 

## 🚀 Instalando <API em go>

Para instalar a <api>, siga estas etapas:
<br>
>Clone o repositório
<br>

Linux:
```
$ git clone https://github.com/JoseTorquato/mconf-joselucas.git
```
<br>

```linux
cd mconf-joselucas/
```

```linux
cd runner
```

```linux
python3 runner.py
```
>agora já temos um script python rodando a api em go.

<img src="image/runner.png" alt="exemplo imagem">

```cmd
localhost:3000
```
<img src="image/dados.png" alt="exemplo imagem">


#### Enquanto isso no terminal:

>Temos os dados sendo mostrado no proprio terminal

<img src="image/dados_terminal.png" alt="exemplo imagem">


<br>


## Docker
<br>

> Com o terminal aberto na raiz do projeto ultilize os seguintes comandos.

```cmd
$ docker build -t mconf/api:joselucas-1 .
```
<img src="image/docker_2.png" alt="exemplo imagem">


## ☕ Usando <api> com docker.

Para usar <api>, siga estas etapas:
<br>

### Container API

```cmd
$ docker run -it --rm -p 3000 mconf/api:joselucas-1
```

```cmd
ls 
```
> Você verá os seguintes intens 
```
Dockerfile  README.md  api.html  go.mod  go.sum  image  server  server.go
```
<img src="image/docker_3.png" alt="exemplo imagem">

Agora só executar o arquivo server com o comando.
```
chmod +x ./server 
```

Abra agora uma pagina web e digite.
```
localhost:3000
```

### Container RUNNER

> Para rodar em container e não localmente comente a linha 8 do programa

<img src="image/runner_docker.png" alt="exemplo imagem">


```linux
$ docker build -t mconf/runner:joselucas-1 .
```
<img src="image/docker_r1.png" alt="exemplo imagem">


```linux
$ docker run -it --rm -p 3000 mconf/runner:joselucas-1
```
> Em construção para os containers se comunicarem




## O que ainda não consegui fazer

- Comunicar os dois containers []
- Buscar dados pelos parametros na linha de comando docker especifico [] 

## Contato

Linkedin: https://www.linkedin.com/in/josetorquato