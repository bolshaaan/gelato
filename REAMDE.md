# Gelato checkout 

Library for golang to manipulate prices in shopping cart

Currently implements 2 types of rules:
 * rule on SKU item - special price of multiple items
 * rule on total of items - e.g. > $200 then 10% of for total price   
 
## Multithreaded
This library is coroutine friendly. When using in many goroutines,
first create one Collection of discounts(rules), and pass it to many instances of carts.
This makes less memory consumption, because collection of rules can be very large.
 
## Example of use
 
````
// make collection of rules
rules := discount.NewCollection()
rules.SetTotalDiscount(1, 20)

// or load this from file
rules = discount.LoadFromFile("path/to/file")

// create cart with rule collection
cart := NewCart(rules)

//add items to cart
cart.Scan(gelato.Item{SKU: "A", Count: 1, Price: 12 })

// print cart Total 
fmt.Println( cart.Total() )
````




