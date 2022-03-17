## Installation
- Firstly you have to set database(MySQL) address in connections/connection.go.
- After setting you can run the project


## About The Project
- When you run the project, there will be 2 tables in the database, users and wallets.
- First register and then login, otherwise API doesn't allow to get or post datas. 
- To check errors, you can try to get users or other endpoints before registration 
- After registiration please login and try to get user by id and then get all users. If you
write invalid user id you will get an error
- After checking user features, you can add new wallet given by userid
- Then you can add credit to wallet or get debit from the wallet by {walletid}.
- You can check the credit and debit can not be less than 0.
- After add credit, you can get debit from balance. But balance can not be less than 0 that's why
it won't allow to play but you can change your debit amount if it is less than balance you can go on to play
- You have to write correct userid instead of {id} on url. And correct walletid instead of {walletid} 


## EndPoints

### POST
- http://localhost:8080/api/register  

/*To register new user please follow the format below*/
```
{
    "Name": "",
    "Email": "",
    "Password": ""
}
```
  

- http://localhost:8080/api/login    
     
/*To login please follow the format below*/
```
{
    "Email": "",
    "Password": ""
}
```

- http://localhost:8080/api/addWallet       

/*To add new wallet please follow the format below*/
```
{
    "userId": 1                    //it will create a wallet for user who has 1 id
}
```

- http://localhost:8080/api/wallets/{walletid}/credit 

/*To add credit please follow the format below*/
```
{
    "credit": 50            //it will add 50 credit to balance. Credit can not be less than 0
}
```
  

- http://localhost:8080/api/wallets/{walletid}/debit 

/*To subtract debit please follow the format below*/
```
{
    "debit": 30               //it will subtract 30 from balance. Debit can not be less than 0
}
```

- http://localhost:8080/api/logout/{id} 

/*user is logged out and token is deleted by given id*/

### GET

	
- http://localhost:8080/api/getUser/{id}        /*get user by given id*/

- http://localhost:8080/api/getAllUser        /*get All Users*/

- http://localhost:8080/api/wallets/{walletid}/balance     /*get wallet by given*/