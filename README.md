# Frontend

Frontend is a simple apps, I created for our Frontend Engineer, Christ Ryandi Aprillo Djaya, so he can help us in creating templates.
So the backend can concentrate on creating API and frontend engineer can create template parallely.

Preparations
------
- Download and Install Golang. The newest golang is OK.


How to create a new endpoint
------
1. Open app.go
2. Add `mux.HandleFunc("/<endpoint url>", <function name>)` below the lowest mux.HandlerFunc
3. Create a function in new file (or same file if it related), u can follow the PaymentPage or AjaxPaymentPage Example
4. Change the data as you want. The PaymentPage Example will print HTML while the AjaxPaymentPage will print JSON
5. Save the file
6. Go to your Terminal
7. Go to location of this folder
8. Run `go build`
8a. If there some error, you must fix it first
9. Run `./frontend`
91. If there some error, you must fix it first
10. Open your browser
11. Go to localhost:8080/<endpoint url>
12. Your page is generated!
