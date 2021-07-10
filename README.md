# Fake Hash - Ports & Adapters

Série de _interfaces_ e suas implementações para auxiliar no desenvolvimento de serviços em _Go_ baseados na arquitetura
Hexagonal. Permite que esses serviços evitem referências diretas a bibliotecas externas ao longo de todas as suas
camadas, isolando essas referências apenas na camada de inicialização.

## Lista de especificações e implementações disponíveis

* Integração com variáveis de ambiente
   * Variáveis do sistema
* _Log_
   * [uber-go/zap](https://github.com/uber-go/zap)
* _Server HTTP_
   * net/http + [gorilla/mux](https://github.com/gorilla/mux)

## Demais pacotes

### `errorhandler`

Possui _interfaces_ que os erros dos serviços deverão implementar para que um _handler_ de erros possa gerenciá-los.
Possui também uma implementação padrão que trata erros em **APIs** **HTTP**.

### `stoppage`

Possui uma _interface_, `Stopper`, que define como os _Adapters_ devem se comportar durante o encerramento da aplicação,
executando métodos das suas respectivas bibliotecas que precisam ser executados neste período.

Possui também uma `struct` que auxilia os serviços a gerenciar as implementações de `Stopper` e chamar seus métodos nas
camadas apropriadas das aplicações.

### `toolkit`

Possui _helpers_ que auxiliam os serviços a executarem determinadas atividades.

#### Correlation ID

ID que pode ser enviado nas requisições **HTTP** e é utilizado nas aplicações para agrupar informações de cada
requisição. Possui métodos para gerenciar esses **IDs** nos `context.Context`.

**Obs.:** Este pacote pode ser desacoplado deste projeto e se tornar um novo módulo, mas para poupar tempo manterei ele
aqui.
