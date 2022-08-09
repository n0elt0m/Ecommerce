# Ecommerce
SIGNUP FUNCTION API CALL (POST REQUEST)
http://localhost:9000/users/signup

{
  "first_name": "noel",
  "last_name": "tom",
  "email": "noel@gmail.com",
  "password": "password",
  "phone": "6475256718"
}
Response :"Successfully Signed Up!!"

LOGIN FUNCTION API CALL (POST REQUEST)

http://localhost:9000/users/login

{
  "email": "noel@gmail.com",
  "password": "password"
}
response will be like this

{
    "_id": "62f1093b41684f348e22900c",
    "first_name": "noel",
    "last_name": "tom",
    "password": "$2a$14$swFqDE2zoF0COrGGKlSxquJ82dUNq4n5QM0nOf.uHLOCQIv.hQ5uK",
    "email": "noel@gmail.com",
    "phone": "6475256718",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im5vZWxAZ21haWwuY29tIiwiRmlyc3RfTmFtZSI6Im5vZWwiLCJMYXN0X05hbWUiOiJ0b20iLCJVaWQiOiI2MmYxMDkzYjQxNjg0ZjM0OGUyMjkwMGMiLCJleHAiOjE2NjAxMDI3NzZ9.ewcGcSvATQVSMSHVjSBrCicdaHIBgPsHKcmusdv-CRg",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2NjAxODkxNzZ9.HlrkL9CVnOYFjkSKb1QckN-P6oIxhOkHT3mAhBYeWKw",
    "created_at": "2022-08-08T13:01:47Z",
    "updated_at": "2022-08-08T13:01:47Z",
    "user_id": "62f1093b41684f348e22900c",
    "usercart": [],
    "address": [],
    "orders": []
}
Admin add Product Function (POST REQUEST)

http://localhost:9000/admin/addproduct

{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}
Response : "Successfully added our Product Admin!!"

View all the Products in db GET REQUEST

http://localhost:9000/users/productview

Response

[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "alienwarex15",
    "price": 1500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "giner ale",
    "price": 900,
    "rating": 5,
    "image": "gin.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "whiskey",
    "price": 100,
    "rating": 7,
    "image": "whis.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "acer predator",
    "price": 3000,
    "rating": 10,
    "image": "acer.jpg"
  }
]
Search Product by regex function (GET REQUEST)
defines the word search sorting http://localhost:9000/users/search?name=al

response:

[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 1500,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "ginger Ale",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
Adding the Products to the Cart (GET REQUEST)

http://localhost:9000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

Corresponding mongodb query

Removing Item From the Cart (GET REQUEST)

http://localhost:9000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

Listing the item in the users cart (GET REQUEST) and total price

http://localhost:9000/listcart?id=xxxxxxuser_idxxxxxxxxxx

Addding the Address (POST REQUEST)

http://localhost:9000/addadress?id=user_id**\*\***\***\*\***

The Address array is limited to two addresses home and work address more than two address is not acceptable

{
  "house_name": "white house",
  "street_name": "white street",
  "city_name": "washington",
  "pin_code": "332423432"
}
Editing the Home Address(PUT REQUEST)

http://localhost:9000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Editing the Work Address(PUT REQUEST)

http://localhost:9000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

Delete Addresses(GET REQUEST)

http://localhost:9000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx


Cart Checkout Function and placing the order(GET REQUEST)

After placing the order the items have to be deleted from cart

http://localhost:9000/cartcheckout?id=xxuser_idxxx

Instantly Buying the Products(GET REQUEST) 
http://localhost:9000/instantbuy?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
