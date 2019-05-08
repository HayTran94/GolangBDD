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

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originAmount := cart.TotalAmount()

		cart.AddItem(itemA)

		It("the shopping cart has 1 more unique item than it had earlier", func(){
			Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount + 1))
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
		cart.AddItem(itemA)

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originTotalAmount := cart.TotalAmount()

		cart.AddItem(itemA)

		It("the shopping cart has the same number of unique item as earlier", func() {
			Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount))
		})
		It("the shopping cart has 1 more unit than it had earlier", func() {
			Expect(cart.TotalUnits()).Should(Equal(originUnitCount + 1))
		})
		It("the total amount increases by item price", func(){
			Expect(cart.TotalAmount()).Should(Equal(originTotalAmount + itemA.Price));
		})
	})

	Context("that has 0 unit of item A", func(){
		cart := Cart {}
		
		cart.AddItem(itemB)
		cart.AddItem(itemB)

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originTotalAmount := cart.TotalAmount()

		Context("removing item A", func(){
			cart.RemoveItem(itemA.ID, 1)

			It("should not change the number of items", func(){
				Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount))
			})
			It("should not change the number of units", func(){
				Expect(cart.TotalUnits()).Should(Equal(originUnitCount))
			})
			It("should not change the amount", func(){
				Expect(cart.TotalAmount()).Should(Equal(originTotalAmount))
			})
		})
	})

	Context("that has 1 unit of item A", func(){
		cart := Cart{}

		cart.AddItem(itemB) 
		cart.AddItem(itemB) 

		cart.AddItem(itemA)

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originTotalAmount := cart.TotalAmount()

		Context("removing 1 unit item A", func() {
			cart.RemoveItem(itemA.ID, 1)
			It("should reduce the number of items by 1", func(){
				Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount - 1))
			})
			It("should reduce the number of units by 1", func(){
				Expect(cart.TotalUnits()).Should(Equal(originUnitCount - 1))
			})
			It("should reduce the amount by the item price", func(){
				Expect(cart.TotalAmount()).Should(Equal(originTotalAmount - itemA.Price))
			})
		})
	})

	Context("that has 2 unit of item A", func(){
		cart := Cart{}
		cart.AddItem(itemB) 
		cart.AddItem(itemB)
		
		cart.AddItem(itemA)
		cart.AddItem(itemA)

		originUniqueItemCount := cart.TotalUniqueItems()
		originUnitCount := cart.TotalUnits()
		originTotalAmount := cart.TotalAmount()

		Context("removing 1 unit item A", func() {
			cart.RemoveItem(itemA.ID, 1)

			It("should not reduce the number of items", func(){
				Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount))
			})
			It("should reduce the number of units by 1", func(){
				Expect(cart.TotalUnits()).Should(Equal(originUnitCount - 1))
			})
			It("should reduce the amount by the item price", func(){
				Expect(cart.TotalAmount()).Should(Equal(originTotalAmount - itemA.Price))
			})
		})

		Context("removing 2 unit item A", func() {
			cart.RemoveItem(itemA.ID, 2)

			It("should reduce the number of items by 1", func(){
				Expect(cart.TotalUniqueItems()).Should(Equal(originUniqueItemCount - 1))
			})
			It("should reduce the number of units by 1", func(){
				Expect(cart.TotalUnits()).Should(Equal(originUnitCount - 2))
			})
			It("should reduce the amount by the item price", func(){
				Expect(cart.TotalAmount()).Should(Equal(originTotalAmount - 2*itemA.Price))
			})
		})
	})

})
