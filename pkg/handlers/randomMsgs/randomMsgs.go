package handlers

import "math/rand"

// 
func SetRandomMsg() (string, string) {
	var (
		winMsg string
		loseMsg string
	)

	successMsgList := [5]string{
	
		"πYippee-ki-yay you win!π\n",
		"πLike Taco Bueno night!!!π\n",
		"π\"All winners cheat\" -Board of Losersπ\n",
		"πYou did it. Good jerb.π\n",
		"πHypebot 4000 approves.π\n",
	}
	failMsgList := [5]string{
		"π»β  You do not pass Go. You do not collect $200.β π»\n",
		"π»β  You zigged when you should have zagged.β π»\n",
		"π»β  All your base are belong to us.β π»\n",
		"π»β  Snake? Snake?!? SNAAAAAAAKE!β π»\n",
		"π»β  You came, you saw, you returned home in shame.β π»\n",
	}

	switch rand.Intn(5) {
	case 0:
		winMsg = successMsgList[0]
	case 1:
		winMsg = successMsgList[1]
	case 2:
		winMsg = successMsgList[2]
	case 3:
		winMsg = successMsgList[3]
	case 4:
		winMsg = successMsgList[4]

	}

	switch rand.Intn(5) {
	case 0:
		loseMsg = failMsgList[0]
	case 1:
		loseMsg = failMsgList[1]
	case 2:
		loseMsg = failMsgList[2]
	case 3:
		loseMsg = failMsgList[3]
	case 4:
		loseMsg = failMsgList[4]

	}


	return winMsg, loseMsg
}