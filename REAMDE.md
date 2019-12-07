# Gelato checkout 

Library for golang to manipulate prices in order

Currently implements 2 types of rules:
 * rule on SKU item - special price of multiple items
 * rule on total of items - e.g. > $200 then 10% of for total price   
 
## Example of use
 
````
co = ​new​ CheckOut(pricingRules);
co.scan(item);
co.scan(item);
fmt.Println( co.total() )
````




