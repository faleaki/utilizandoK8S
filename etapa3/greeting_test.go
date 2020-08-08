package main

import "testing"

func TestGreeting( t *testing.T) {
      mensagem := "Code.education Rocks!"
      retorno := greeting(mensagem)
      esperado := "<b>"+mensagem+"</b>"
      
      if retorno == ""  {
	     t.Error("greeting function returned empty string")
      }
      if retorno != esperado {
	      t.Errorf("greeting function returned %v, expected %v", retorno, esperado)
      }
}
