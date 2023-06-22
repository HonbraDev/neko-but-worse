package uwu

import "math/rand"

func randomFace() string {
	eye := eyes[rand.Intn(len(eyes))]
	mouth := mouths[rand.Intn(len(mouths))]

	return eye[0] + mouth + eye[1]
}
