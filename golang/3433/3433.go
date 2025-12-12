package main

import (
	"sort"
	"strconv"
)

func sortEvents(events [][]string) [][]string {
	sort.Slice(events, func(i, j int) bool {
		timeI, _ := strconv.Atoi(events[i][1])
		timeJ, _ := strconv.Atoi(events[j][1])

		// Ordenar por timestamp
		if timeI != timeJ {
			return timeI < timeJ
		}

		// Se timestamps iguais, OFFLINE vem antes
		return events[i][0] == "OFFLINE" && events[j][0] != "OFFLINE"
	})
	return events
}

func countMentions(numberOfUsers int, events [][]string) []int {
	events = sortEvents(events)
	users := make([]int, numberOfUsers)
	usersOffline := make([]bool, numberOfUsers)
	offlineDeadline := make(map[int][]int)

	for _, event := range events {
		eventType := event[0]
		timestamp, _ := strconv.Atoi(event[1])
		eventData := event[2]

		for deadline := range offlineDeadline {
			if deadline <= timestamp {
				for _, userId := range offlineDeadline[deadline] {
					usersOffline[userId] = false
				}
				delete(offlineDeadline, deadline)
			}
		}

		switch eventType {
		case "MESSAGE":
			if eventData == "ALL" {
				for i := 0; i < numberOfUsers; i++ {
					users[i]++
				}
			} else if eventData == "HERE" {
				for i, offline := range usersOffline {
					if !offline {
						users[i]++
					}
				}
			} else {
				start := 0
				for end, char := range eventData {
					if char == ' ' {
						mention := eventData[start+2 : end]
						start = end + 1
						userId, _ := strconv.Atoi(mention)
						users[userId]++
					}
				}
				mention := eventData[start+2:]
				userId, _ := strconv.Atoi(mention)
				users[userId]++
			}
		case "OFFLINE":
			endOffline := timestamp + 60
			userId, _ := strconv.Atoi(eventData)
			if _, exists := offlineDeadline[endOffline]; !exists {
				offlineDeadline[endOffline] = make([]int, 0)
			}
			usersOffline[userId] = true
			offlineDeadline[endOffline] = append(offlineDeadline[endOffline], userId)
		}
	}
	return users
}
