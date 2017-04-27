package main

import "fmt"

//Declaração da estrutura de Pilha
type Pilha struct {
  topo *Elemento; //define o topo
  tamanho int; //tamanho da pilha
}

//Declaração do elemento da pilha
type Elemento struct {
  valor interface{};//Um tipo interface vazio aceita qualquer tipo de dado
  proximo *Elemento;//Referencia para o próximo elemento
}

//Método para retornar o tamanho da pilha
func (pilha *Pilha) length() int {
  return pilha.tamanho;
}

//Método de push na pilha
func (pilha *Pilha) push(valor interface{}) {
  pilha.topo = &Elemento{valor, pilha.topo}//Cria o elemento, e já adiciona o seu endereço no topo da pilha
  pilha.tamanho++//incrementa o seu tamanho
}

//Método de pop na pilha
func (pilha *Pilha) pop() bool{
  if(pilha.tamanho > 0) {
    pilha.topo = pilha.topo.proximo
    pilha.tamanho--
    return true
  }
  return false
}

//Método para exibição dos elementos da Pilha
func (pilha *Pilha) show() {
  temp := pilha.topo //Atribui a variável temporária o topo da Pilha
  fmt.Println("-------------")
  for temp.valor != nil { //Verifica se o valor a ser exibido agora é nulo
    fmt.Println(temp.valor)//Imprime o valor atual
    if(temp.proximo != nil){ //Valida o próximo valor, se ele for nulo, não atribui mais nenhum valor
      temp = temp.proximo//Atribui ao próximo o endereço próximo
    }else{//Sai do método
      fmt.Println("-------------")
      return
    }
  }
}

func main() {
  minhaPilha := new(Pilha)//"Instância" da pilha [Lembre-se que GO não é OO]

  //Pushs
  minhaPilha.push("Java")
  minhaPilha.push("Ruby")
  minhaPilha.push("NodeJS")
  minhaPilha.push("Golang")
  minhaPilha.push("JavaScript")

  //Exibe o tamanho da pilha
  fmt.Println("Tamanho da pilha: ", minhaPilha.length())

  //Mostra os valores da pilha
  minhaPilha.show()

  //Pop na pilha
  minhaPilha.pop()

  fmt.Println("")  

  //Exibe o tamanho da pilha
  fmt.Println("Tamanho da pilha: ", minhaPilha.length())
  //Mostra os valores da pilha
  minhaPilha.show()


}