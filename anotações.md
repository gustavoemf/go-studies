### Hellow world!
- variáveis não utilizadas não pode, para contornar isso devemos utilizar "_" (underscore),
isso porque funções possuem a opção de retornar mais de um valor, _ captura os valores que
não iremos usar

### Declaração curta de variáveis / Operador curto de declaração
- := marmota
    - tipagem automática
    - só pode repetir se tiver variáveis novas, somente em declarações
    - = (atribuição) é diferente de := (declaração + atribuição)
    - só funciona dentro de clodeblocks (scope)
- operadores e expressões
    - expressão é algo que produz resultado mas não gera uma ação (evaluate)

### Keyword "var"
- só use a declaração longa de variáveis quando não inicializar uma variável
- use a forma curta para declarar e inicializar

### Tipos

- Go é uma linguagem de tipos estáticos
- Tipos de dados primitivos: int, string e bool
- Tipos de dados compostos: Tipos compostos de tipos primitivo criados pelo usuário

### Valor zero
- valores presentes em uma variável antes dela ser inicializada
- os zeros:
    - int: 0
    - float: 0.0
    - bool: false
    - string: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- usar := sempre que possível e var em package-level


### fmt
- interpreted string literal vs raw string literal
- Rune: cada caracter de uma string
- um literal é uma notação para representar um valor fixo no código fonte
- Format printing: documentação
    - Grupo #1: Print -> standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
    - Grupo #2: Print -> string, podendo ser usado como variável
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err errror)
- entre cada item no Println há um espaço enquanto que no Sprint não ao concatenar

### criando tipos
- criar um tipo próprio habilita o uso de métodos e interfaces
- tipo subjacente é o tipo interno daquele tipo que foi criado
- apesar de subjacente não posso atribuir os valores de um para o outro

### conversão de tipos
- em Go não existe coersion nem casting apenas convertion
- T(x) -> "T" é o tipo que eu quero e "x" o valor que quero converter para "T"

### strings
- carreira de bytes seuquenciais
- strings são imutáveis
- uma string é um "slice of bytes"
- conversão para slice of bytes: []byte(x)

### sistemas numéricos
- base 10: decimal 0~9
- base 2: binário 0~1
- base 16: hexadecimal 0~f

### constantes
- são valores imutáveis
- podem ser tipadas ou não
    - as não tipadas só terão um tipo atribuído a elas quando forem usadas
    - constantes em Go podem mudar de tipo, mas nunca de significado
    - se mudar o significado -> conversão explícita obrigatória
- tipos de constantes são definidos no uso enquanto o tipo de variáveis são definidos durante a atribuição
- podem ser declaradas de uma só vez como uma lista, assim como os imports