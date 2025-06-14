package main

// Problem: Add Two Numbers
// Link: https://leetcode.com/problems/add-two-numbers/
// Difficulty: Medium
// Tags: Linked List, Math
// Time Complexity: O(max(m, n)), where m and n are the lengths of the two linked lists.
// Space Complexity: O(max(m, n)), for the resulting linked list.

// Estrutura de nó para lista ligada que representa dígitos de um número
type ListNode struct {
	Val  int       // valor do dígito (0-9)
	Next *ListNode // ponteiro para o próximo nó
}

// Soma dois números representados por listas ligadas na ordem reversa
// Cada nó contém um único dígito e os dígitos são armazenados em ordem reversa
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // nó auxiliar para facilitar a construção da lista resultado
	current := dummy     // ponteiro para o nó atual na construção da lista
	carry := 0           // variável para armazenar o "vai um" da soma

	// Continua enquanto houver nós em qualquer lista ou carry pendente
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry // inicia a soma com o carry da operação anterior

		// Adiciona o valor do nó atual de l1, se existir
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next // avança para o próximo nó
		}

		// Adiciona o valor do nó atual de l2, se existir
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next // avança para o próximo nó
		}

		carry = sum / 10                        // calcula o novo carry (divisão inteira por 10)
		current.Next = &ListNode{Val: sum % 10} // cria novo nó com o dígito das unidades
		current = current.Next                  // avança para o próximo nó na lista resultado
	}

	return dummy.Next // retorna a lista resultado (pula o nó dummy inicial)
}
