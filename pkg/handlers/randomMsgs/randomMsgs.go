package randomMsgs

import "math/rand"

// 
func SetRandomMsg() (string, string) {
	var (
		winMsg string
		loseMsg string
	)

	successMsgList := [5]string{
	
		"🙌Yippee-ki-yay you win!🙌\n",
		"🙌Like Taco Bueno night!!!🙌\n",
		"🙌\"All winners cheat\" -Board of Losers🙌\n",
		"🙌You did it. Good jerb.🙌\n",
		"🙌Hypebot 4000 approves.🙌\n",
	}
	failMsgList := [5]string{
		"🔻☠ You do not pass Go. You do not collect $200.☠🔻\n",
		"🔻☠ You zigged when you should have zagged.☠🔻\n",
		"🔻☠ All your base are belong to us.☠🔻\n",
		"🔻☠ Snake? Snake?!? SNAAAAAAAKE!☠🔻\n",
		"🔻☠ You came, you saw, you returned home in shame.☠🔻\n",
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