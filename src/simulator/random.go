package simulator

import (
	"math/rand"
	"time"

	"github.com/aselimkaya/nbasimulator/src/collection"
)

func setAttackOrder(team1, team2 *collection.TeamGameInfo) (*collection.TeamGameInfo, *collection.TeamGameInfo) {
	if generateRandomNumber(2) == 0 {
		return team1, team2
	}
	return team2, team1
}

//isSuccessfulAttack determines the attack's situation. 45% of attacks are successful
func isSuccessfulAttack() bool {
	attack := generateRandomNumber(20)

	if attack < 9 {
		return true
	}

	return false
}

//shoot function determines the shoot type and its result. 70% shoots are 2 pointer and 30% are 3 pointer
func shoot() (int, bool) {
	shootingType := generateRandomNumber(10)

	if shootingType < 7 {
		return shootTwoPointer()
	}
	return shootThreePointer()

}

//shootTwoPointer describes the 2 pointer shooting type. 40% of 2 pointers are successful
func shootTwoPointer() (int, bool) {
	shoot2PT := generateRandomNumber(5)

	if shoot2PT < 2 {
		return 2, true
	}
	return 2, false
}

//shootThreePointer describes the 3 pointer shooting type. 30% of 3 pointers are successful
func shootThreePointer() (int, bool) {
	shoot3PT := generateRandomNumber(10)

	if shoot3PT < 3 {
		return 3, true
	}
	return 3, false
}

//getPlayersWhoScoreAndAssist function randomly select two players who scores and assists among team
func getPlayersWhoScoreAndAssist(teamLen int) (int, int) {
	score, assist := generateRandomNumber(teamLen), generateRandomNumber(teamLen)

	// must be different players
	for score == assist {
		assist = generateRandomNumber(teamLen)
	}

	return score, assist
}

func generateRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
