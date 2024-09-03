package main

import (
	"deptdb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Root path handler
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Dept API!")
	})

	// Favicon handler
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	// Handler for retrieving all department records
	r.GET("/dept", func(c *gin.Context) {
		records, err := deptdb.GetRecords()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, records)
	})

	// Handler for retrieving a specific department record by deptno
	r.GET("/dept/:deptno", func(c *gin.Context) {
		deptno, err := strconv.Atoi(c.Param("deptno"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department number"})
			return
		}

		records, err := deptdb.GetRecords()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var foundDept *deptdb.Dept
		for _, record := range records {
			if record.Deptno == deptno {
				foundDept = &record
				break
			}
		}

		if foundDept == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
			return
		}

		c.JSON(http.StatusOK, foundDept)
	})

	// Handler for inserting a new department record
	r.POST("/dept", func(c *gin.Context) {
		var newDept deptdb.Dept
		if err := c.ShouldBindJSON(&newDept); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := deptdb.InsertRecord(newDept); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Record inserted"})
	})

	// Handler for updating an existing department record
	r.PUT("/dept/:deptno", func(c *gin.Context) {
		deptno, err := strconv.Atoi(c.Param("deptno"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department number"})
			return
		}

		var updatedDept deptdb.Dept
		if err := c.ShouldBindJSON(&updatedDept); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedDept.Deptno = deptno // Ensure the deptno from the URL is used

		if err := deptdb.UpdateRecord(updatedDept); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Record updated"})
	})

	// Handler for deleting a department record
	r.DELETE("/dept/:deptno", func(c *gin.Context) {
		deptno, err := strconv.Atoi(c.Param("deptno"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department number"})
			return
		}

		if err := deptdb.DeleteRecord(deptno); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Record deleted"})
	})

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
