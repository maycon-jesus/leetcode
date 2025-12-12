# Explicação Detalhada - Count Mentions Per User

## Visão Geral do Problema

O problema simula um sistema de menções de usuários semelhante a plataformas de chat (Discord, Slack, etc.), onde:
- Usuários podem ser mencionados em mensagens
- Usuários podem ficar offline temporariamente
- Diferentes tipos de menção têm comportamentos distintos (ALL, HERE, id específico)

## Desafios Principais

### 1. Ordem dos Eventos é Crucial
Eventos com o mesmo timestamp devem ser processados em ordem específica: eventos `OFFLINE` antes de `MESSAGE`. Isso garante que mudanças de status sejam aplicadas antes das mensagens.

### 2. Gerenciamento de Estado Temporal
Usuários ficam offline por exatamente 60 unidades de tempo. Precisamos rastrear quando cada usuário volta online.

### 3. Tipos Diferentes de Menção
- `ALL`: menciona todos (online + offline)
- `HERE`: menciona apenas online
- `id<number>`: menciona usuários específicos (podem estar offline)
- Duplicatas devem ser contadas

## Estrutura da Solução

### Estruturas de Dados

```go
users := make([]int, numberOfUsers)              // Contador de menções por usuário
usersOffline := make([]bool, numberOfUsers)      // Status: true = offline, false = online
offlineDeadline := make(map[int][]int)           // timestamp -> lista de userIDs que voltam online
```

#### Por que essas estruturas?

1. **`users[]`**: Array simples para contar menções. Acesso O(1) por ID.

2. **`usersOffline[]`**: Array booleano para verificar rapidamente se um usuário está offline. Crucial para processar menções "HERE".

3. **`offlineDeadline{}`**: Mapa que agrupa usuários por timestamp de retorno.
   - Chave: timestamp quando usuários voltam online
   - Valor: lista de IDs que voltam naquele momento
   - Permite processar múltiplos usuários que voltam no mesmo timestamp de uma vez

## Passo a Passo do Algoritmo

### Fase 1: Ordenação

```go
func sortEvents(events [][]string) [][]string {
    sort.Slice(events, func(i, j int) bool {
        timeI, _ := strconv.Atoi(events[i][1])
        timeJ, _ := strconv.Atoi(events[j][1])

        // Primeiro critério: timestamp
        if timeI != timeJ {
            return timeI < timeJ
        }

        // Segundo critério: OFFLINE vem antes de MESSAGE
        return events[i][0] == "OFFLINE" && events[j][0] != "OFFLINE"
    })
    return events
}
```

**Por que ordenar assim?**
- Timestamps em ordem crescente: processa eventos cronologicamente
- OFFLINE antes de MESSAGE no mesmo timestamp: garante que status é atualizado antes de processar menções

**Exemplo:**
```
Entrada: [["MESSAGE","10","HERE"], ["OFFLINE","10","0"]]
Após ordenação: [["OFFLINE","10","0"], ["MESSAGE","10","HERE"]]
```

### Fase 2: Processamento de Eventos

#### Loop Principal

```go
for _, event := range events {
    eventType := event[0]
    timestamp, _ := strconv.Atoi(event[1])
    eventData := event[2]

    // 1. Processar retornos online
    // 2. Processar evento atual
}
```

#### Subprocesso: Retornos Online

```go
for deadline := range offlineDeadline {
    if deadline <= timestamp {
        for _, userId := range offlineDeadline[deadline] {
            usersOffline[userId] = false
        }
        delete(offlineDeadline, deadline)
    }
}
```

**O que faz:**
- Verifica todos os deadlines agendados
- Se `deadline <= timestamp atual`, usuários voltam online
- Remove deadline do mapa (limpeza de memória)

**Complexidade:** O(D × U_d) onde D = número de deadlines e U_d = usuários por deadline

### Fase 3: Processamento por Tipo de Evento

#### Tipo 1: MESSAGE com "ALL"

```go
if eventData == "ALL" {
    for i := 0; i < numberOfUsers; i++ {
        users[i]++
    }
}
```

**Comportamento:**
- Incrementa contador de TODOS os usuários
- Não verifica status online/offline
- **Complexidade:** O(U)

#### Tipo 2: MESSAGE com "HERE"

```go
else if eventData == "HERE" {
    for i, offline := range usersOffline {
        if !offline {
            users[i]++
        }
    }
}
```

**Comportamento:**
- Itera sobre todos os usuários
- Incrementa apenas se `!offline` (usuário está online)
- **Complexidade:** O(U)

#### Tipo 3: MESSAGE com IDs específicos

```go
else {
    start := 0
    for end, char := range eventData {
        if char == ' ' {
            mention := eventData[start+2 : end]  // Pula "id" prefix
            start = end + 1
            userId, _ := strconv.Atoi(mention)
            users[userId]++
        }
    }
    // Processa última menção (não tem espaço após)
    mention := eventData[start+2:]
    userId, _ := strconv.Atoi(mention)
    users[userId]++
}
```

**Como funciona o parsing:**

```
Input: "id1 id0 id3"
       ↓
Iteração 1: encontra ' ' na posição 3
  - Extrai "id1"[2:] = "1"
  - users[1]++
  - start = 4

Iteração 2: encontra ' ' na posição 7
  - Extrai "id0"[2:] = "0"
  - users[0]++
  - start = 8

Após loop: processa resto
  - Extrai "id3"[2:] = "3"
  - users[3]++
```

**Por que `start+2`?**
- Pula o prefixo "id" de cada menção
- "id1" → [2:] → "1"

**Complexidade:** O(M) onde M = comprimento da string

#### Tipo 4: OFFLINE

```go
case "OFFLINE":
    endOffline := timestamp + 60
    userId, _ := strconv.Atoi(eventData)

    // Cria lista se não existe
    if _, exists := offlineDeadline[endOffline]; !exists {
        offlineDeadline[endOffline] = make([]int, 0)
    }

    usersOffline[userId] = true
    offlineDeadline[endOffline] = append(offlineDeadline[endOffline], userId)
```

**Etapas:**
1. Calcula quando usuário volta: `timestamp + 60`
2. Marca usuário como offline imediatamente
3. Agenda retorno no timestamp calculado

**Exemplo:**
```
Evento: ["OFFLINE", "10", "2"]

Ações:
- endOffline = 10 + 60 = 70
- usersOffline[2] = true
- offlineDeadline[70] = [2]

No timestamp 70 ou após, user 2 volta online
```

## Exemplo Completo de Execução

Vamos processar o exemplo do `main()`:

```go
numberOfUsers = 5
events = [
    {"OFFLINE", "28", "1"},
    {"OFFLINE", "14", "2"},
    {"MESSAGE", "2", "ALL"},
    {"MESSAGE", "28", "HERE"},
    {"OFFLINE", "66", "0"},
    {"MESSAGE", "34", "id2"},
    {"MESSAGE", "83", "HERE"},
    {"MESSAGE", "40", "id3 id3 id2 id4 id4"},
]
```

### Após Ordenação

```
[
    {"MESSAGE", "2", "ALL"},           # timestamp 2
    {"OFFLINE", "14", "2"},            # timestamp 14
    {"OFFLINE", "28", "1"},            # timestamp 28 - OFFLINE primeiro
    {"MESSAGE", "28", "HERE"},         # timestamp 28 - MESSAGE depois
    {"MESSAGE", "34", "id2"},          # timestamp 34
    {"MESSAGE", "40", "id3 id3 id2 id4 id4"},
    {"OFFLINE", "66", "0"},
    {"MESSAGE", "83", "HERE"},
]
```

### Execução Passo a Passo

**Estado Inicial:**
```
users = [0, 0, 0, 0, 0]
usersOffline = [false, false, false, false, false]
offlineDeadline = {}
```

---

**Evento 1:** `["MESSAGE", "2", "ALL"]`
- Não há deadlines para processar
- Tipo: ALL → incrementa todos
- **users = [1, 1, 1, 1, 1]**

---

**Evento 2:** `["OFFLINE", "14", "2"]`
- Não há deadlines ≤ 14
- User 2 fica offline até timestamp 74
- usersOffline[2] = true
- offlineDeadline[74] = [2]

---

**Evento 3:** `["OFFLINE", "28", "1"]`
- Não há deadlines ≤ 28
- User 1 fica offline até timestamp 88
- usersOffline[1] = true
- offlineDeadline[88] = [1]

**Estado:**
```
users = [1, 1, 1, 1, 1]
usersOffline = [false, true, true, false, false]
offlineDeadline = {74: [2], 88: [1]}
```

---

**Evento 4:** `["MESSAGE", "28", "HERE"]`
- Não há deadlines ≤ 28
- Tipo: HERE → incrementa apenas online
- Online: 0, 3, 4 (users 1 e 2 estão offline)
- **users = [2, 1, 1, 2, 2]**

---

**Evento 5:** `["MESSAGE", "34", "id2"]`
- Não há deadlines ≤ 34
- Menciona user 2 (mesmo offline)
- **users = [2, 1, 2, 2, 2]**

---

**Evento 6:** `["MESSAGE", "40", "id3 id3 id2 id4 id4"]`
- Não há deadlines ≤ 40
- Parsing: id3, id3, id2, id4, id4
- Incrementa: 3 (x2), 2 (x1), 4 (x2)
- **users = [2, 1, 3, 4, 4]**

---

**Evento 7:** `["OFFLINE", "66", "0"]`
- Não há deadlines ≤ 66
- User 0 offline até 126
- usersOffline[0] = true
- offlineDeadline[126] = [0]

**Estado:**
```
usersOffline = [true, true, true, false, false]
offlineDeadline = {74: [2], 88: [1], 126: [0]}
```

---

**Evento 8:** `["MESSAGE", "83", "HERE"]`
- **Processar deadlines:**
  - 74 ≤ 83? Sim → user 2 volta online
  - 88 ≤ 83? Não
  - usersOffline[2] = false
  - delete offlineDeadline[74]
- Tipo: HERE → incrementa online
- Online: 2, 3, 4 (users 0 e 1 offline)
- **users = [2, 1, 4, 5, 5]**

**Estado Final:**
```
users = [2, 1, 4, 5, 5]
usersOffline = [true, true, false, false, false]
```

**Resultado Esperado:** `[2, 1, 4, 5, 5]` ✅

## Análise de Complexidade

### Complexidade de Tempo

**Componentes:**

1. **Ordenação:** O(E log E)
   - E eventos para ordenar

2. **Loop Principal:** O(E)
   - Itera sobre cada evento uma vez

3. **Processar Deadlines (por evento):** O(D × U_d)
   - D = número de deadlines no momento
   - U_d = média de usuários por deadline
   - No pior caso: O(E × U)

4. **Processar Evento:**
   - ALL: O(U)
   - HERE: O(U)
   - IDs: O(M) onde M = comprimento da string
   - OFFLINE: O(1)

**Total:** O(E log E + E × U)
- Dominado pela ordenação e pelo processamento de eventos ALL/HERE

### Complexidade de Espaço

1. **Arrays:** O(U)
   - users[], usersOffline[]

2. **Mapa de Deadlines:** O(E)
   - No pior caso, todos os eventos são OFFLINE

3. **Ordenação:** O(E)
   - Espaço temporário para sort

**Total:** O(U + E)

## Possíveis Otimizações

### 1. Otimizar Processamento de Deadlines

**Problema Atual:** Itera sobre todos os deadlines em cada evento.

**Solução:** Usar heap (min-heap) de deadlines
```go
type DeadlineHeap []Deadline
// Processa apenas deadlines que expiraram
```
**Ganho:** O(D × log D) em vez de O(D × E)

### 2. Parsing de Menções

**Alternativa:** Usar `strings.Fields()` e `strings.TrimPrefix()`
```go
mentions := strings.Fields(eventData)
for _, mention := range mentions {
    id := strings.TrimPrefix(mention, "id")
    userId, _ := strconv.Atoi(id)
    users[userId]++
}
```
**Trade-off:** Mais legível, mesma complexidade O(M)

### 3. Evitar Re-slice do Mapa

Em vez de deletar deadlines, marcar como processados:
```go
processed := make(map[int]bool)
```
**Trade-off:** Usa mais memória, mas evita operações de delete

## Casos Extremos (Edge Cases)

### 1. Múltiplos Eventos no Mesmo Timestamp
```
[["OFFLINE","10","0"], ["OFFLINE","10","1"], ["MESSAGE","10","HERE"]]
```
✅ Ordenação garante que OFFLINE seja processado primeiro

### 2. Usuário Volta Exatamente no Timestamp da Mensagem
```
[["OFFLINE","10","0"], ["MESSAGE","70","HERE"]]
```
✅ User 0 volta em 70, processado antes da mensagem

### 3. Menções Duplicadas
```
["MESSAGE","5","id1 id1 id1"]
```
✅ Cada menção incrementa separadamente → users[1] += 3

### 4. Todos os Usuários Offline
```
[["OFFLINE","1","0"], ["OFFLINE","1","1"], ["MESSAGE","2","HERE"]]
```
✅ HERE não menciona ninguém

### 5. String de Menção Vazia?
Segundo constraints, sempre há pelo menos 1 menção → não precisa tratar

## Lições Importantes

1. **Ordenação é fundamental** para problemas de simulação temporal
2. **Estruturas de dados adequadas** simplificam a lógica (hash map para deadlines)
3. **Cuidado com parsing manual** de strings (off-by-one errors)
4. **Considere edge cases** de timestamps iguais
5. **Limpeza de dados** (delete deadlines processados) evita memory leak

## Conclusão

A solução equilibra:
- ✅ Simplicidade de implementação
- ✅ Eficiência aceitável para os constraints (≤ 100 users/events)
- ✅ Correção em todos os casos

Para problemas maiores, otimizações com heap seriam necessárias, mas para este problema, a abordagem atual é ideal.
