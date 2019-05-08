package cart_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "cart"
)

var _ = Describe("Cart", func() {
	itemA := Item{ID: "itemA", Name: "Item A", Price: 10.20, Qty: 0}
	itemB := Item{ID: "itemB", Name: "Item B", Price: 7.66, Qty: 0}

	Context("initialy", func() {
		cart := Cart{}
		It("has 0 items", func() {
			Expect(cart.TotalUniqueItems()).Should(BeZero())
		})
		It("has 0 units", func() {
			Expect(cart.TotalUnits()).Should(BeZero())
		})
		It("the total amount is 0.00", func() {
			Expect(cart.TotalAmount()).Should(BeZero())
		})
	})

	Context("when a new item is added", func() {
		cart := Cart{}

		cart.AddItem(itemA)

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originAmount := cart.TotalAmount()

		cart.AddItem(itemA)

		It("the shopping cart has 1 more unique item than it had earlier", func(){
			Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount))
		})
		It("the shopping cart has 1 more unit than it had earlier", func() {
			Expect(cart.TotalUnits()).Should(Equal(originUnitCount + 1))
		})
		It("the total amount increases by item price", func() {
			Expect(cart.TotalAmount()).Should(Equal(originAmount + itemA.Price))
		})
	})

	Context("when an existing item is added", func() {
		cart := Cart{}
		

		It("the shopping cart has the same number of unique item as earlier", func() {})
		It("the shopping cart has 1 more unit than it had earlier", func() {})
		It("the total amount increases by item price", func(){})
	})

	Context("that has 0 unit of item A", func(){
		Context("removing item A", func(){
			It("should not change the number of items", func(){})
			It("should not change the number of units", func(){})
			It("should not change the amount", func(){})
		})
	})

	Context("that has 1 unit of item A", func(){
		Context("removing 1 unit item A", func() {
			It("should reduce the number of items by 1", func(){})
			It("should reduce the number of units by 1", func(){})
			It("should reduce the amount by the item price", func(){})
		})
	})

	Context("that has 2 unit of item A", func(){
		Context("removing 1 unit item A", func() {
			It("should not reduce the number of items", func(){})
			It("should reduce the number of units by 1", func(){})
			It("should reduce the amount by the item price", func(){})
		})

		Context("removing 2 unit item A", func() {
			It("should reduce the number of items by 1", func(){})
			It("should reduce the number of units by 1", func(){})
			It("should reduce the amount by the item price", func(){})
		})
	})

})
