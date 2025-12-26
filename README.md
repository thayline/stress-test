:========= Stress Test CLI em Go ==========:

Ferramenta de linha de comando (CLI) escrita em Go para realizar testes de carga (stress test) em serviços web, permitindo configurar número total de requisições e nível de concorrência.

############################################
Objetivo

Executar múltiplas requisições HTTP simultâneas contra uma URL e gerar um relatório final com métricas de desempenho e status das respostas.

############################################
Funcionalidades

-> Execução via Docker
-> Execução de testes via CLI
-> Controle de concorrência com goroutines
-> Garantia do número total de requisições
-> Relatório final com:
    -> Tempo total de execução
    -> Total de requisições realizadas
    -> Quantidade de respostas HTTP 200
    -> Distribuição de outros status HTTP (404, 500, etc.)

############################################
Parâmetros de Entrada (CLI)

Flag	        Descrição
--url	        URL do serviço a ser testado
--requests	    Número total de requisições
--concurrency	Número de requisições simultâneas

Exemplo:
    > go run cmd/stress/main.go --url=http://google.com --requests=1000 --concurrency=10

############################################
Exemplo de Relatório

:===== Stress Test Report =====:
Total time: 3.42s
Total requests: 1000
Status 200: 998
Other status codes:
  500 -> 2

############################################
Executando com Docker

Build da imagem
    > docker build -t stress-test .

Executar o teste
    > docker run stress-test --url=http://google.com --requests=1000 --concurrency=10