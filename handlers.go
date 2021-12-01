package main

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"math.io/crath/ord"
)

func getOrd(ctx *gin.Context) {
	var input OrdSubmit

	if err := ctx.BindJSON(&input); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect parameters"})
		return
	}

	g, isSet := big.NewInt(0).SetString(input.Element, 0)
	if !isSet || g == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect element"})
		return
	}

	m, isSet := big.NewInt(0).SetString(input.Module, 0)
	if !isSet || m == nil || !m.ProbablyPrime(1) {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "incorrect module"})
		return
	}

	solution := ord.Ord(g, m)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"solution": solution.String(),
	})
}
