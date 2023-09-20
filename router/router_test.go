package router

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitializeRouter(t *testing.T) {
	router := InitializeRouter()
	// criar uma instância de roteador

	responseRecorder := httptest.NewRecorder()
	// Isso é usado para registrar a resposta HTTP que será
	// gerada durante o teste. É uma implementação de http.ResponseWriter usada em testes para capturar a saída HTTP.

	request, _ := http.NewRequest("GET", "/", nil)
	//criando uma nova solicitação HTTP GET para a rota "/ping".
	//O objeto request representa a solicitação HTTP que será enviada para o roteador.

	router.ServeHTTP(responseRecorder, request)
	// Está sendo chamando o método ServeHTTP do roteador router para manipular a solicitação
	// request e registrar a resposta resultante no objeto responseRecorder. Isso simula o processamento de uma
	//solicitação HTTP real pelo roteador.

	assert.Equal(t, 200, responseRecorder.Code)
	// Aqui está sendo usado a biblioteca assert para verificar se o código de status HTTP na resposta
	//registrada em responseRecorder é igual a 200. Isso verifica se a rota "/" retornou um código de status 200 (OK).

	assert.Equal(t, "{\"message\":\"pong\"}", responseRecorder.Body.String())
	// Aqui está sendo usado assert novamente para verificar se o corpo da resposta registrada em responseRecorder é igual a "{"message": "pong"}".
	// Isso verifica se o conteúdo da resposta corresponde ao esperado.

}
