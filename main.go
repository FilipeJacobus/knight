package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "cache-control", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost/"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/get-positions", func(c *gin.Context) {
		var response interface{}
		statusCodeResp := 500
		var psts1 []string
		var psts2 []string
		var err error
		pst1 := c.Query("player1")
		pst2 := c.Query("player2")

		if pst1 != "" {
			psts1, err = getPositions(pst1)
		} else {
			err = errors.New("Invalid params")
		}
		if pst2 != "" {
			psts2, err = getPositions(pst2)
		} else {
			err = errors.New("Invalid params")
		}

		if err == nil {
			statusCodeResp = 200
			response = map[string]interface{}{
				"player1": psts1,
				"player2": psts2,
			}
		} else {
			response = map[string]interface{}{
				"error": err.Error(),
			}
		}

		c.JSON(statusCodeResp, response)
	})

	r.Run(":9876")

}

func getPositions(ent string) ([]string, error) {

	entX, entY, err := entToPoints(ent)
	if err != nil {
		return nil, err
	}

	var positions []string

	q1x, q1y := firstQuad(entX, entY)
	if q1x != nil {
		positions = append(positions, slcToPosition(q1x))
	}
	if q1y != nil {
		positions = append(positions, slcToPosition(q1y))
	}

	q2x, q2y := secondQuad(entX, entY)
	if q2x != nil {
		positions = append(positions, slcToPosition(q2x))
	}
	if q2y != nil {
		positions = append(positions, slcToPosition(q2y))
	}

	q3x, q3y := thirdQuad(entX, entY)
	if q3x != nil {
		positions = append(positions, slcToPosition(q3x))
	}
	if q3y != nil {
		positions = append(positions, slcToPosition(q3y))
	}

	q4x, q4y := fourthQuad(entX, entY)
	if q4x != nil {
		positions = append(positions, slcToPosition(q4x))
	}
	if q4y != nil {
		positions = append(positions, slcToPosition(q4y))
	}

	return positions, nil
}

func entToPoints(ent string) (x, y int, err error) {

	if len(ent) == 2 {

		for i, v := range chars {

			if string(ent[0]) == v {

				x = i + 1

			}
		}
		if x > 8 || x < 1 {
			return x, y, errors.New("Invalid position")
		}
		y, err = strconv.Atoi(string(ent[1]))

	} else {
		err = errors.New("Invalid position")
	}
	return x, y, err
}

func slcToPosition(ent []int) string {
	x := chars[(ent[0] - 1)]
	y := strconv.Itoa(ent[1])
	return x + y
}

func firstQuad(x, y int) (q1x, q1y []int) {
	ax := x + 2
	ay := y + 1
	if (ax <= 8) && (ay <= 8) {
		q1x = []int{ax, ay}
	}

	bx := x + 1
	by := y + 2
	if (bx <= 8) && (by <= 8) {
		q1y = []int{bx, by}
	}

	return q1x, q1y

}

func secondQuad(x, y int) (q2x, q2y []int) {

	ax := x - 1
	ay := y + 2
	if (ax > 0) && (ay <= 8) {
		q2x = []int{ax, ay}
	}

	bx := x - 2
	by := y + 1
	if (bx > 0) && (by <= 8) {
		q2y = []int{bx, by}
	}

	return q2x, q2y

}

func thirdQuad(x, y int) (q3x, q3y []int) {
	ax := x - 2
	ay := y - 1
	if (ax > 0) && (ay > 0) {
		q3x = []int{ax, ay}
	}

	bx := x - 1
	by := y - 2
	if (bx > 0) && (by > 0) {
		q3y = []int{bx, by}
	}

	return q3x, q3y

}

func fourthQuad(x, y int) (q4x, q4y []int) {

	ax := x + 1
	ay := y - 2
	if (ax <= 8) && (ay > 0) {
		q4x = []int{ax, ay}
	}

	bx := x + 2
	by := y - 1
	if (bx <= 8) && (by > 0) {
		q4y = []int{bx, by}
	}

	return q4x, q4y

}
