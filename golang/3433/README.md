# 3433. Count Mentions Per User

**Dificuldade:** Medium
**Status:** ✅ Solved
**Tópicos:** Array, Hash Table, String, Simulation

## Descrição do Problema

Você recebe um inteiro `numberOfUsers` representando o número total de usuários e um array `events` de tamanho `n x 3`.

Cada `events[i]` pode ser um dos seguintes tipos:

### 1. Evento de Mensagem
**Formato:** `["MESSAGE", "timestamp", "mentions_string"]`

Indica que um conjunto de usuários foi mencionado em uma mensagem no `timestamp`.

A string `mentions_string` pode conter:
- `id<number>`: onde `<number>` é um inteiro no intervalo [0, numberOfUsers - 1]. Pode haver múltiplos IDs separados por espaço e podem conter duplicatas. Menciona até usuários offline.
- `ALL`: menciona todos os usuários (online e offline).
- `HERE`: menciona apenas usuários online.

### 2. Evento Offline
**Formato:** `["OFFLINE", "timestamp", "id"]`

Indica que o usuário `id` ficou offline no `timestamp` por 60 unidades de tempo. O usuário voltará automaticamente online em `timestamp + 60`.

### Regras Importantes
- Todos os usuários começam online
- Mudanças de status (offline/online) são processadas antes de eventos de mensagem no mesmo timestamp
- Cada menção deve ser contada separadamente (duplicatas contam)

**Retorno:** Um array `mentions` onde `mentions[i]` representa o número total de menções que o usuário com ID `i` recebeu.

## Exemplos

### Exemplo 1
```
Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]
Output: [2,2]

Explicação:
- Timestamp 10: id0 e id1 são mencionados → mentions = [1,1]
- Timestamp 11: id0 fica offline
- Timestamp 71: id0 volta online e "HERE" menciona todos online → mentions = [2,2]
```

### Exemplo 2
```
Input: numberOfUsers = 2, events = [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]
Output: [2,2]

Explicação:
- Timestamp 10: id0 e id1 mencionados → mentions = [1,1]
- Timestamp 11: id0 fica offline
- Timestamp 12: "ALL" menciona todos (incluindo offline) → mentions = [2,2]
```

### Exemplo 3
```
Input: numberOfUsers = 2, events = [["OFFLINE","10","0"],["MESSAGE","12","HERE"]]
Output: [0,1]

Explicação:
- Timestamp 10: id0 fica offline
- Timestamp 12: "HERE" menciona apenas online (só id1) → mentions = [0,1]
```

## Constraints

- `1 <= numberOfUsers <= 100`
- `1 <= events.length <= 100`
- `events[i].length == 3`
- `events[i][0]` será `MESSAGE` ou `OFFLINE`
- `1 <= int(events[i][1]) <= 10⁵`
- Número de menções `id<number>` em cada evento `MESSAGE`: entre 1 e 100
- `0 <= <number> <= numberOfUsers - 1`
- É garantido que o usuário referenciado no evento `OFFLINE` está online quando o evento ocorre

## Solução

### Abordagem

A solução utiliza **simulação de eventos** com as seguintes estratégias:

1. **Ordenação de Eventos**: Ordena eventos por timestamp. Quando timestamps são iguais, eventos `OFFLINE` vêm primeiro (garante que mudanças de status sejam processadas antes de mensagens).

2. **Rastreamento de Estado**:
   - `users[]`: contador de menções para cada usuário
   - `usersOffline[]`: indica se o usuário está offline
   - `offlineDeadline{}`: mapeia timestamps de retorno online aos usuários

3. **Processamento de Eventos**:
   - Antes de cada evento, verifica e processa usuários que voltaram online
   - **MESSAGE**: incrementa contadores conforme tipo de menção (ALL/HERE/id)
   - **OFFLINE**: marca usuário como offline e agenda retorno

4. **Parsing de Menções**: Para menções específicas (`id<number>`), faz parsing da string separando por espaços e extraindo IDs.

### Complexidade

- **Tempo:** O(E log E + E × U)
  - `E log E`: ordenação dos eventos
  - `E × U`: no pior caso, processar cada evento pode iterar sobre todos os usuários (eventos "ALL" ou "HERE")
  - Parsing de menções: O(M) onde M é o comprimento da string de menções

- **Espaço:** O(U + E)
  - `U`: arrays para rastrear estado dos usuários
  - `E`: no pior caso, todos os eventos podem estar no mapa `offlineDeadline`

Onde:
- `E` = número de eventos
- `U` = número de usuários
- `M` = comprimento da string de menções

### Implementação

Ver arquivo [3433.go](3433.go) para a implementação completa em Go.

#### Funções Principais

- **`sortEvents`**: Ordena eventos por timestamp, priorizando eventos OFFLINE em caso de empate
- **`countMentions`**: Função principal que simula os eventos e conta menções

## Como Executar

```bash
go run 3433.go
```

## Links

- [LeetCode Problem #3433](https://leetcode.com/problems/count-mentions-per-user/)