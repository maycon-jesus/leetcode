3562. Maximum Profit from Trading Stocks with Discounts
Hard
Topics
premium lock iconCompanies
Hint

You are given an integer n, representing the number of employees in a company. Each employee is assigned a unique ID from 1 to n, and employee 1 is the CEO. You are given two 1-based integer arrays, present and future, each of length n, where:

    present[i] represents the current price at which the ith employee can buy a stock today.
    future[i] represents the expected price at which the ith employee can sell the stock tomorrow.

The company's hierarchy is represented by a 2D integer array hierarchy, where hierarchy[i] = [ui, vi] means that employee ui is the direct boss of employee vi.

Additionally, you have an integer budget representing the total funds available for investment.

However, the company has a discount policy: if an employee's direct boss purchases their own stock, then the employee can buy their stock at half the original price (floor(present[v] / 2)).

Return the maximum profit that can be achieved without exceeding the given budget.

Note:

    You may buy each stock at most once.
    You cannot use any profit earned from future stock prices to fund additional investments and must buy only from budget.

 

Example 1:

Input: n = 2, present = [1,2], future = [4,3], hierarchy = [[1,2]], budget = 3

Output: 5

Explanation:

    Employee 1 buys the stock at price 1 and earns a profit of 4 - 1 = 3.
    Since Employee 1 is the direct boss of Employee 2, Employee 2 gets a discounted price of floor(2 / 2) = 1.
    Employee 2 buys the stock at price 1 and earns a profit of 3 - 1 = 2.
    The total buying cost is 1 + 1 = 2 <= budget. Thus, the maximum total profit achieved is 3 + 2 = 5.

Example 2:

Input: n = 2, present = [3,4], future = [5,8], hierarchy = [[1,2]], budget = 4

Output: 4

Explanation:

    Employee 2 buys the stock at price 4 and earns a profit of 8 - 4 = 4.
    Since both employees cannot buy together, the maximum profit is 4.

Example 3:

Input: n = 3, present = [4,6,8], future = [7,9,11], hierarchy = [[1,2],[1,3]], budget = 10

Output: 10

Explanation:

    Employee 1 buys the stock at price 4 and earns a profit of 7 - 4 = 3.
    Employee 3 would get a discounted price of floor(8 / 2) = 4 and earns a profit of 11 - 4 = 7.
    Employee 1 and Employee 3 buy their stocks at a total cost of 4 + 4 = 8 <= budget. Thus, the maximum total profit achieved is 3 + 7 = 10.

Example 4:

Input: n = 3, present = [5,2,3], future = [8,5,6], hierarchy = [[1,2],[2,3]], budget = 7

Output: 12

Explanation:

    Employee 1 buys the stock at price 5 and earns a profit of 8 - 5 = 3.
    Employee 2 would get a discounted price of floor(2 / 2) = 1 and earns a profit of 5 - 1 = 4.
    Employee 3 would get a discounted price of floor(3 / 2) = 1 and earns a profit of 6 - 1 = 5.
    The total cost becomes 5 + 1 + 1 = 7 <= budget. Thus, the maximum total profit achieved is 3 + 4 + 5 = 12.

 

Constraints:

    1 <= n <= 160
    present.length, future.length == n
    1 <= present[i], future[i] <= 50
    hierarchy.length == n - 1
    hierarchy[i] == [ui, vi]
    1 <= ui, vi <= n
    ui != vi
    1 <= budget <= 160
    There are no duplicate edges.
    Employee 1 is the direct or indirect boss of every employee.
    The input graph hierarchy is guaranteed to have no cycles.

---

## Solução

### Abordagem

Este problema pode ser resolvido usando **Programação Dinâmica em Árvore** combinada com **Knapsack (Mochila)**.

#### Principais Insights:

1. **Estrutura de Árvore**: A hierarquia dos funcionários forma uma árvore com o CEO (funcionário 1) como raiz.

2. **Desconto Condicional**: Um funcionário pode comprar ações com desconto (metade do preço) apenas se seu chefe direto também comprar ações.

3. **Problema de Otimização**: Precisamos maximizar o lucro total sem exceder o orçamento, considerando as dependências de desconto.

### Solução com DP em Árvore

A solução utiliza DFS (Depth-First Search) para processar a árvore de baixo para cima, calculando estados de programação dinâmica para cada nó.

#### Estados DP:

Para cada nó `u`, mantemos dois arrays DP:

- **`dp0[budget]`**: Lucro máximo usando até `budget` quando **NÃO compramos ações do nó u**
- **`dp1[budget]`**: Lucro máximo usando até `budget` quando **COMPRAMOS ações do nó u**

Durante o processamento de cada nó, também rastreamos:

- **`subProfit0[budget]`**: Lucro máximo dos filhos quando o desconto **NÃO está disponível**
- **`subProfit1[budget]`**: Lucro máximo dos filhos quando o desconto **está disponível** (porque o pai comprou)

#### Algoritmo:

1. **Construir o grafo** da hierarquia (árvore adjacente)

2. **DFS recursivo** a partir do CEO (nó 0):
   - Para cada nó, processar recursivamente todos os filhos
   - Combinar resultados dos filhos usando knapsack:
     - Iterar sobre todas as combinações de orçamento
     - Maximizar lucro combinando diferentes alocações de orçamento entre filhos

3. **Calcular estados finais** para o nó atual:
   - `dp0[i]`: Não comprar este nó → usar todo orçamento nos filhos sem desconto
   - `dp1[i]`: Comprar este nó → considerar:
     - Comprar com desconto (se disponível): custo = `present[u] / 2`
     - Comprar sem desconto: custo = `present[u]`
     - Alocar orçamento restante nos filhos com desconto disponível

4. **Retornar** `dp0[budget]` da raiz (CEO)

### Complexidade

- **Tempo**: O(n × budget²) onde processamos cada nó e para cada nó fazemos operações de knapsack
- **Espaço**: O(n × budget) para armazenar os arrays DP

### Como Executar

```bash
go test -v
```

