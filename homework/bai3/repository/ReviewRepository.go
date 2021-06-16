package repository

import (
	"github.com/TechMaster/golang/08Fiber/Repository/model"
	// repo "github.com/TechMaster/golang/08Fiber/Repository/repository"
)

type ReviewRepo struct {
	reviews map[int64]*model.Review
	autoID  int64 //đây là biến đếm tự tăng gán giá trị cho id của Review
}

var Reviews ReviewRepo

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	review.Id = nextID
	r.reviews[nextID] = review //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}
func (r *ReviewRepo) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}
func (r *ReviewRepo) GetRatingByBookId(bookId int64) float32 {
	var rating float32 = 0
	var count int = 0
	for _, review := range r.reviews {
		if review.BookId == bookId {
			rating += float32(review.Rating)
			count++
		}
	}
	return rating / float32(count)
}
