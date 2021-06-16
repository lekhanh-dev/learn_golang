package controller

import (
	"fmt"

	"github.com/TechMaster/golang/08Fiber/Repository/model"
	repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(repo.Reviews.GetAllReviews())
}

func CreateReview(c *fiber.Ctx) error {
	review := new(model.Review)

	err := c.BodyParser(&review)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// Check book has exist
	isBookExist := CheckBookExits(review.BookId)
	if isBookExist {
		// add new review
		reviewId := repo.Reviews.CreateNewReview(review)
		// update rating book
		// computed new rating
		var rating float32 = repo.Reviews.GetRatingByBookId(review.BookId)
		// find book added review
		book := FindBookById(review.BookId)
		book.Rating = rating
		UpdateBookV2(book)

		return c.SendString(fmt.Sprintf("New review is created successfully with id = %d : %f", reviewId, rating))
	} else {
		return c.SendString("Book not exits")
	}
}
