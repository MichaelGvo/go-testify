package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тестирую MainHandler на корректность запроса, возврат кода ответа 200 и что тело ответа не пустое
func TestMainHandlerdGaveCorrectNotEmptyResp(t *testing.T) {
	//Делаю запрос req функцией NewRequest() из пакета httptest
	//В качестве аргумента функция принимает HTTP метод запроса, URL запроса и тело запроса
	//Обработчик mainHandle() не требует тело запроса, поэтому третьим параметром в NewRequest() передано nil
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)
	//Создаю responseRecorder для записи ответа
	responseRecorder := httptest.NewRecorder()
	//Добавляю переменные expected, status и body - чтобы далее сравнить на возврат кода ответа 200 и что тело ответа не пустое
	expected := http.StatusOK
	status := responseRecorder.Code
	body := responseRecorder.Body.String()
	//Вызываю обработчик mainHandle(), в который передаю responseRecorder и req
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//функцией NoError() из пакета require проверяем на пустоту ответа, иначе также нет смысла дальше проверять (потому require)
	require.NotEmpty(t, body)
	//функцией NoError() из пакета require проверяем на возврат кода ответа 200
	require.Equal(t, status, expected)
}

// Тестирую MainHandler на возврат кода ответа 400 и тело ответа - "wrong city value"
func TestMainHandlerWhenWrongCity(t *testing.T) {
	//Делаю запрос req функцией NewRequest() из пакета httptest
	//В качестве аргумента функция принимает HTTP метод запроса, URL запроса и тело запроса
	//Передаю в URL запроса - город который не поддерживается
	//Обработчик mainHandle() не требует тело запроса, поэтому третьим параметром в NewRequest() передано nil
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=dubai", nil)
	//Создаю responseRecorder для записи ответа
	responseRecorder := httptest.NewRecorder()
	//Добавляю переменные body, expectedBody, status, expectedStatus - чтобы далее удобно было сравнить на "wrong city value" и код ответа 400
	body := responseRecorder.Body.String()
	expectedBody := "wrong city value"
	status := responseRecorder.Code
	expectedStatus := http.StatusBadRequest
	//Вызываю обработчик mainHandle(), в который передаю responseRecorder и req
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//функцией Equal() из пакета assert проверяем на код ответа 400, пользуемся при этом assert
	assert.Equal(t, status, expectedStatus)
	//функцией Equal() из пакета assert проверяем на тело ответа - "wrong city value", пользуемся при этом assert
	assert.Equal(t, body, expectedBody)
}

// Тестирую MainHandler на большее количество кафе, чем есть всего
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//Всего у нас 4 кафе, добавляю переменную totalCount
	totalCount := 4
	//Делаю запрос req функцией NewRequest() из пакета httptest
	//В качестве аргумента функция принимает HTTP метод запроса, URL запроса и тело запроса
	//Передаю в URL запроса - большое число кафе, больше чем есть
	//Обработчик mainHandle() не требует тело запроса, поэтому третьим параметром в NewRequest() передано nil
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=88005553555&city=moscow", nil)
	//Создаю responseRecorder для записи ответа
	responseRecorder := httptest.NewRecorder()
	//Добавляю переменные body, list и listCount
	//В list - функцией Split() превратим body в слайс строк.
	//В качестве первого аргумента она принимает строку, а в качестве второго разделитель (в данном случае запятая — «,»)
	//Далее нам нужен будет len переменной list - ввел listCount
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	listCount := len(list)
	//Вызываю обработчик mainHandle(), в который передаю responseRecorder и req
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//функцией Equal() из пакета assert проверяем на количество кафе, пользуемся при этом assert
	assert.Equal(t, totalCount, listCount)
}
