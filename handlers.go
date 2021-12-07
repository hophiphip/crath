package main

import (
	"log"
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
