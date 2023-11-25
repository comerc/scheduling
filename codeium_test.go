package codeium_test

import (
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	// Параметры для расписания
	startTime := time.Date(2023, time.January, 2, 8, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, time.January, 6, 17, 0, 0, 0, time.UTC)
	breakTime := 1 * time.Hour
	developers := []string{"Dev1", "Dev2", "Dev3", "Dev4"}
	sessionsPerDay := 2
	sessionDuration := 1 * time.Hour
	rotationInterval := 1
	expectedPairsPerSession := 2

	// Вызов функции-планировщика
	schedule := Schedule(developers, startTime, endTime, breakTime, sessionsPerDay, sessionDuration, rotationInterval)

	// Проверка, что количество сессий соответствует ожидаемому
	expectedSessions := (endTime.Sub(startTime) / breakTime) * sessionsPerDay
	if len(schedule) != int(expectedSessions) {
		t.Errorf("Unexpected number of sessions. Expected: %d, Got: %d", expectedSessions, len(schedule))
	}

	// Проверка, что каждая сессия содержит правильное количество пар
	for _, session := range schedule {
		pairs := session.Pairs
		if len(pairs) != expectedPairsPerSession {
			t.Errorf("Unexpected number of pairs in session. Expected: %d, Got: %d", expectedPairsPerSession, len(pairs))
		}
	}

	// Проверка, что разработчики равномерно распределены и ротируются каждую новую сессию
	for i := 0; i < len(schedule)-1; i++ {
		session1 := schedule[i]
		session2 := schedule[i+1]

		// Проверка, что каждая пара в следующей сессии содержит новых разработчиков
		for j := 0; j < expectedPairsPerSession; j++ {
			developer1 := session1.Pairs[j].Developer1
			developer2 := session1.Pairs[j].Developer2
			nextDeveloper1 := session2.Pairs[j].Developer1
			nextDeveloper2 := session2.Pairs[j].Developer2

			if developer1 == nextDeveloper1 || developer1 == nextDeveloper2 ||
				developer2 == nextDeveloper1 || developer2 == nextDeveloper2 {
				t.Errorf("Developers are not rotating properly. Session %d: %s, %s; Session %d: %s, %s",
					i, developer1, developer2, i+1, nextDeveloper1, nextDeveloper2)
			}
		}
	}
}
