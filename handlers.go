package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	finitefield "math.io/crath/finite_field"
	"math.io/crath/ord"
)

func getOrd(ctx *gin.Context) {
	var input OrdSubmit

	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		return
	}

	log.Println(input.Module)

	g, isSet := big.NewInt(0).SetString(input.Element, 0)
	if !isSet || g == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect element"})
		return
	}

	log.Println(g.String())

	m, isSet := big.NewInt(0).SetString(input.Module, 0)
	if !isSet || m == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "could not parse module"})
		return
	}

	solution, err := ord.Ord(g, m)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "incorrect element",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"solution": solution.String(),
		})
	}
}

func getFieldState(ctx *gin.Context) {
	var input FiniteFieldSubmit

	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		return
	}

	px, err := finitefield.BitStringToByte(input.PX)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect `px` field parameters", "description": err})
		return
	}

	result := make([][]FiniteFieldResult, 256)
	for idx := range result {
		result[idx] = make([]FiniteFieldResult, 256)
	}

	for a := byte(0); ; a++ {
		for b := byte(0); ; b++ {
			result[a][b] = FiniteFieldResult{
				Sum: fmt.Sprintf("%08b", finitefield.AddG_2_8(a, b)),
				Mul: fmt.Sprintf("%08b", finitefield.MulG_2_8(a, b, px)),
			}

			if b == byte(0b1111_1111) {
				break
			}
		}

		if a == byte(0b1111_1111) {
			break
		}
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"solution": result,
	})
}
