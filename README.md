# Projeto de Consulta de Endereço

Este projeto é uma API simples em Go que consulta dados de endereço utilizando duas APIs externas: **BrasilAPI** e **ViaCEP**. Ao receber uma requisição com um CEP, a aplicação dispara requisições simultâneas para ambas as APIs e retorna os dados do endereço obtidos da primeira resposta bem-sucedida.

## Funcionalidades

- Consulta de endereço a partir do CEP informado.
- Integração com as APIs:
  - [BrasilAPI](https://brasilapi.com.br/)
  - [ViaCEP](https://viacep.com.br/)
- Resposta no formato JSON com os campos do endereço e a indicação da API que retornou a resposta.
- Timeout de 1 segundo para evitar esperas excessivas se nenhuma API responder.

## Requisitos

- [Go](https://golang.org/dl/) (versão 1.13 ou superior)
- Conexão com a internet para acessar as APIs externas.

## Como Executar

Siga os passos abaixo para rodar o projeto localmente:

1. **Clone o repositório**

   ```bash
   git clone https://seu-repositorio-url.git
   cd nome-do-repositorio
   ```

2. **Instale as dependências**

   No Go, as dependências são gerenciadas através dos módulos. Caso o repositório já contenha um arquivo `go.mod`, as dependências serão automaticamente baixadas ao rodar o comando de build ou run. Caso não esteja configurado, inicialize o módulo:

   ```bash
   go mod init nome-do-modulo
   go mod tidy
   ```

3. **Execute a aplicação**

   Utilize o comando abaixo para rodar a aplicação:

   ```bash
   go run main.go
   ```

   Você deverá ver a seguinte mensagem indicando que o servidor está ativo:

   ```
   Servidor rodando na porta 8080...
   ```

4. **Realize uma requisição para buscar um endereço**

   Abra o seu navegador ou utilize uma ferramenta como `curl` ou [Postman](https://www.postman.com/) e acesse a seguinte URL (substitua `01001000` pelo CEP desejado):

   ```bash
   http://localhost:8080/buscar?cep=01001000
   ```

   Exemplo usando `curl`:

   ```bash
   curl "http://localhost:8080/buscar?cep=01001000"
   ```

   A resposta será um JSON com os dados do endereço, semelhante ao exemplo abaixo:

   ```json
   {
     "cep": "01001-000",
     "logradouro": "Praça da Sé",
     "bairro": "Sé",
     "cidade": "São Paulo",
     "estado": "SP",
     "api_origem": "BrasilAPI"
   }
   ```

   Caso nenhuma API responda dentro de 1 segundo, a API retornará um erro de timeout.

## Estrutura do Projeto

- **main.go**: Contém a função `main` que inicia o servidor HTTP e define a rota `/buscar`.
- **fetchAddress**: Função utilizada para realizar as requisições assíncronas às APIs externas.
- **handleCEPRequest**: Handler responsável por processar as requisições, extrair o CEP e disparar as requisições para as APIs.

## Possíveis Melhorias

- Implementação de cache para consultas repetidas.
- Adicionar logs para monitoramento e depuração.
- Tratar e retornar mensagens de erro mais detalhadas.

## Contribuição

1. Faça um fork do projeto.
2. Crie uma branch para sua feature: `git checkout -b minha-nova-feature`
3. Commit suas alterações: `git commit -am 'Adiciona nova feature'`
4. Push na branch: `git push origin minha-nova-feature`
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

---

Com esses passos, você já estará pronto para executar e testar a API de consulta de endereço. Se tiver alguma dúvida ou sugestão, sinta-se à vontade para abrir uma issue ou contribuir com o projeto!