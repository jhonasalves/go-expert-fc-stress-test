# Desafio Técnico Go Expert - Stress Test

## Objetivo

Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. Após a execução dos testes, o sistema gerará um relatório com informações detalhadas.

## Funcionalidades

### Entrada de Parâmetros via CLI
- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

### Execução do Teste
- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

### Geração de Relatório
O relatório gerado ao final dos testes incluirá:
- Tempo total gasto na execução.
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Execução da Aplicação

### Gerar a Imagem Docker
Para gerar a imagem Docker da aplicação, execute o seguinte comando na raiz do projeto:

```bash
docker build -t loadtester .
```

### Executar a Aplicação
A aplicação pode ser executada via Docker. Exemplo de uso:

```bash
docker run loadtester --url https://httpbin.org/status/200,404,500 --requests 1000 --concurrency 10
```
Exemplo de Resultado:
```bash
--- Load Test Report ---
Total execution time: 17.370500767s
Total requests made: 1000
Requests with status 200: 340


---------------------- Status Codes ----------------------
Status     Number of Requests  
---------------------------------------------------------
200        340                 
404        324                 
500        336                 

Test completed successfully!
```

## Sobre o httpbin

O [httpbin](https://httpbin.org) é um serviço simples que permite testar e depurar requisições HTTP. Ele fornece endpoints para simular diferentes cenários de requisições e respostas HTTP, como redirecionamentos, autenticação, cabeçalhos personalizados, entre outros.

Neste projeto, o httpbin está sendo utilizado como exemplo para testes de requisições HTTP, permitindo validar o comportamento do código em diferentes situações.