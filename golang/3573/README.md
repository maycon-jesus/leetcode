# 3573. Best Time to Buy and Sell Stock V

**Difficulty:** Medium

## Problem Description

You are given an integer array `prices` where `prices[i]` is the price of a stock in dollars on the i-th day, and an integer `k`.

You are allowed to make at most `k` transactions, where each transaction can be either of the following:

- **Normal transaction:** Buy on day i, then sell on a later day j where i < j. You profit `prices[j] - prices[i]`.
- **Short selling transaction:** Sell on day i, then buy back on a later day j where i < j. You profit `prices[i] - prices[j]`.

**Note:** You must complete each transaction before starting another. Additionally, you can't buy or sell on the same day you are selling or buying back as part of a previous transaction.

Return the maximum total profit you can earn by making at most `k` transactions.

## Examples

### Example 1:
```
Input: prices = [1,7,9,8,2], k = 2
Output: 14
```

**Explanation:**
We can make $14 of profit through 2 transactions:
- A normal transaction: buy the stock on day 0 for $1 then sell it on day 2 for $9.
- A short selling transaction: sell the stock on day 3 for $8 then buy back on day 4 for $2.

### Example 2:
```
Input: prices = [12,16,19,19,8,1,19,13,9], k = 3
Output: 36
```

**Explanation:**
We can make $36 of profit through 3 transactions:
- A normal transaction: buy the stock on day 0 for $12 then sell it on day 2 for $19.
- A short selling transaction: sell the stock on day 3 for $19 then buy back on day 4 for $8.
- A normal transaction: buy the stock on day 5 for $1 then sell it on day 6 for $19.

## Constraints

- `2 <= prices.length <= 10^3`
- `1 <= prices[i] <= 10^9`
- `1 <= k <= prices.length / 2`

## Solution Approach

Esta solução utiliza **Programação Dinâmica** com três estados possíveis para cada dia e número de transações:

### Estados DP

Definimos `dp[i][j][state]` como o lucro máximo após o dia `i`, com `j` transações completas, no estado:
- `state = 0`: Sem posição (sem ações ou short)
- `state = 1`: Posição long (segurando ação comprada)
- `state = 2`: Posição short (vendeu a descoberto)

### Transições

Para cada dia `i` e número de transações `j`:

**Do estado sem posição (0):**
- Não fazer nada: `dp[i+1][j][0]`
- Comprar ação (iniciar normal transaction): `dp[i+1][j][1] = dp[i][j][0] - price`
- Vender a descoberto (iniciar short selling): `dp[i+1][j][2] = dp[i][j][0] + price`

**Do estado long (1):**
- Não fazer nada: `dp[i+1][j][1]`
- Vender ação (completar normal transaction): `dp[i+1][j+1][0] = dp[i][j][1] + price`

**Do estado short (2):**
- Não fazer nada: `dp[i+1][j][2]`
- Comprar de volta (completar short selling): `dp[i+1][j+1][0] = dp[i][j][2] - price`

### Complexidade

- **Tempo:** O(n × k) onde n é o número de dias
- **Espaço:** O(n × k)

### Otimizações Possíveis

A solução pode ser otimizada para usar O(k) de espaço utilizando apenas duas linhas do array DP (atual e próxima), já que cada dia depende apenas do dia anterior.

## Como Executar

### Executar o código:
```bash
go run 3573.go
```

### Executar os testes:
```bash
go test -v
```

### Executar testes com coverage:
```bash
go test -cover
```

### Executar benchmarks:
```bash
go test -bench=. -benchmem
```

## Resultados dos Testes

Os testes cobrem os seguintes cenários:

- ✅ Exemplos fornecidos no problema
- ✅ Transações normais (compra e venda)
- ✅ Transações de short selling
- ✅ Estratégias mistas
- ✅ Preços sempre crescentes
- ✅ Preços sempre decrescentes
- ✅ Preços constantes
- ✅ Casos extremos (arrays vazios, valores mínimos/máximos)
- ✅ Diferentes valores de k
- ✅ Testes de benchmark para análise de performance

**Coverage:** 100% de cobertura de código

