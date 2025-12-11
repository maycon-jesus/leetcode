<p align="center">
  <img src="https://assets.leetcode.com/static_assets/public/LeetCode_logo.png" width="120" alt="LeetCode Logo"/>
</p>

<h1 align="center">📚 Soluções LeetCode</h1>

<p align="center">
  <a href="https://leetcode.com/maycon-jesus/"><img src="https://img.shields.io/badge/LeetCode-Perfil-orange?logo=leetcode" alt="LeetCode Profile"/></a>
  <img src="https://img.shields.io/github/last-commit/maycon-jesus/leetcode?label=Última%20Atualização" alt="Last Updated"/>
  <img src="https://img.shields.io/badge/Problemas%20Resolvidos-37+-blue" alt="Problems Solved"/>
  <img src="https://img.shields.io/badge/Linguagem-Go-00ADD8?logo=go" alt="Language"/>
</p>

<p align="center">
  <img src="https://leetcard.jacoblin.cool/maycon-jesus?theme=radical&font=DM%20Serif%20Display&ext=heatmap" width="500" alt="LeetCode Profile Card" />
</p>

---

## 🎯 Sobre Este Repositório

Uma coleção abrangente de soluções de problemas do LeetCode implementadas em Go. Cada solução é desenvolvida com:
- **Código limpo e eficiente** otimizado para complexidade de tempo e espaço
- **Suítes de testes abrangentes** com cobertura extensiva de casos extremos
- **Explicações detalhadas** de algoritmos e abordagens
- **Testes de benchmark** demonstrando características de performance

> **Nota**: Este repositório foca em compreender algoritmos fundamentais e escrever código com qualidade de produção, não apenas passar nos casos de teste.

---

## 📊 Estatísticas

| Dificuldade | Quantidade | Percentual |
|-------------|-----------|------------|
| 🟢 Fácil    | 15+       | ~40%       |
| 🟡 Médio    | 15+       | ~40%       |
| 🔴 Difícil  | 6+        | ~20%       |

---

## 🗂️ Estrutura do Repositório

```
leetcode/
└── golang/
    ├── 1/          # Two Sum
    ├── 13/         # Roman to Integer
    ├── 14/         # Longest Common Prefix
    └── .../
        ├── README.md           # Descrição do problema e restrições
        ├── {number}.go         # Implementação da solução
        ├── {number}_test.go    # Testes unitários e benchmarks
        └── EXPLANATION.md      # Explicação detalhada (quando disponível)
```

---

## 📝 Padrões de Qualidade das Soluções

Cada solução neste repositório inclui:

### ✅ Qualidade de Código
- Implementação limpa e legível
- Nomes de variáveis significativos
- Tratamento adequado de erros
- Otimizado para performance

### ✅ Testes
- Testes unitários abrangentes (20-30+ casos de teste por problema)
- Cobertura de casos extremos
- Testes de benchmark com métricas de alocação de memória
- Casos de exemplo da descrição do problema

### ✅ Documentação
- **README.md**: Descrição do problema, exemplos, restrições e abordagens de solução
- **EXPLANATION.md**: Passo a passo detalhado da solução, análise de complexidade e técnicas de otimização

### ✅ Performance
- Análise de complexidade de tempo e espaço
- Zero ou mínimas alocações de memória quando possível
- Resultados de benchmark demonstrando performance no mundo real

---

## 🚀 Início Rápido

### Executando Soluções

```bash
# Navegue para o diretório do problema
cd golang/14

# Execute a solução
go run 14.go

# Execute os testes
go test -v

# Execute os benchmarks
go test -bench=. -benchmem
```

### Exemplo de Output

```
=== RUN   TestLongestCommonPrefix
--- PASS: TestLongestCommonPrefix (0.00s)
PASS

BenchmarkLongestCommonPrefix-22    179573352    6.827 ns/op    0 B/op    0 allocs/op
```

---

## 📚 Soluções em Destaque

| Problema | Destaques | Performance |
|---------|-----------|-------------|
| [#13 - Roman to Integer](./golang/13/) | Abordagem limpa baseada em map com otimização early exit | 0 B/op, 0 allocs/op |
| [#14 - Longest Common Prefix](./golang/14/) | Varredura horizontal com zero alocações | 2-23 ns/op, 30 casos de teste |

---

## 🛠️ Tecnologias e Ferramentas

- **Linguagem**: Go 1.21+
- **Testes**: Framework de testes nativo do Go
- **Benchmarking**: Ferramentas de benchmark nativas do Go
- **Controle de Versão**: Git & GitHub

---

## 🤝 Contribuindo

Embora este seja um repositório de aprendizado pessoal, sugestões e melhorias são bem-vindas!

### Como Contribuir
1. Faça um fork do repositório
2. Crie uma branch de feature (`git checkout -b feature/melhoria`)
3. Faça commit das suas mudanças (`git commit -m 'Adiciona melhoria'`)
4. Faça push para a branch (`git push origin feature/melhoria`)
5. Abra um Pull Request

### Diretrizes de Contribuição
- Siga o estilo e estrutura de código existente
- Inclua testes abrangentes
- Adicione explicações para algoritmos complexos
- Certifique-se de que os benchmarks mostram boa performance

---

## 📫 Conecte-se

- 💼 [Perfil LeetCode](https://leetcode.com/maycon-jesus/)
- 🐙 [GitHub](https://github.com/maycon-jesus)

---

## 📜 Licença

Este projeto é open source e está disponível sob a [Licença MIT](LICENSE).

---

<p align="center">
  <i>Construído com 💻 e ☕️</i>
  <br>
  <b>Boa Codificação!</b> 🚀
</p>
