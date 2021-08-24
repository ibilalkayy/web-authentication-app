# [web-authentication-app](https://web-authentication-app.herokuapp.com/)

![Screenshot 2021-08-24 at 13-21-05 Home](https://user-images.githubusercontent.com/64713734/130583988-82998f8c-18a7-4dc7-b2a9-17ff685caae5.png)

This repository contains the code of a web authentication app that is a simple web app containing home, signup, login, about, and logout pages. Here it will authenticate the user and by the use of sessions, it made easier to limit the visitation of different pages.

This software is written in Golang, HTML, and CSS. It contains four directories in which Go code, Database, HTML and CSS templates are written separately. Those directories are.
- **database**. It handles the database stuff and I used **MongoDB** database in this code.
- **handler**. It contains many functions like home, signup, login, about, etc for performing various tasks.
- **templates**. It contains all the HTML files.
- **static**. It contains all CSS files to design the pages.

---

### Table of Contents

The headers will be used to reference the location of destination.

- [The First Directory](#the-first-directory)
- [The Second Directory](#the-second-directory)
- [The Third Directory](#the-third-directory)
- [The Forth Directory](#the-forth-directory)
- [Author Info](#author-info)

---

# The First Directory
In the first directory **database**, I created one file that is **db.go**. It contains four functions that will handle different tasks.
- The first function is ```Connection()```. It will make the connection with the **MongoDB** database. This function is equal to the ```Connect``` variable that will be called in further functions.
- The second function is ```Insertdata()```. It will enter the data into database. It uses ```InsertOne``` to insert the data.
- The third function is ```Findaccount()```. It will find the account data in the database. It uses ```FindOne``` to find the data based on email address and the password. It also uses ```CompareHashAndPassword()``` to compare passwords and check their similarity. 
- The forth function is ```Updatedata()```. It will update an about field in the database. It filters the data based on email address and the password and it allows the user to modify the data.

# The Second Directory
In the second directory **handler**, I created one file that is **handler.go**. It contains ten functions. Let's see what they do?
- The first function is ```init()```. It will provide cookie sessions based on the secret key. It will be further used in the code.
- The second function is ```makeTemplate()```. It will help us get rid of the common paths used every time.
- The third function is ```Home()```. It will be our first page when we visit our site. It contains home template and another template containg error message if the page is not found.
- The fourth function is ```hashAndSalt()```. It will convert the password into hash. I will use that hash password in the database.
- The fifth function is ```comparePasswords()```. It will help us when we want to compare password and confirm password to check whether both are similar or not.
- The sixth function is ```Signup()```. It will open the signup page, take the form values, compare the passwords, check whether both are similar or not. If they're similar then insert the data in the database, show the success message and redirect to the login page. If the passwords are not similar then show the failure message and give an option to signup again.
- The seventh function is ```Login()```. It will open the login page, take the form values, find the data in the database. If the data is found then make the session value ```true``` and redirect to the about page. If the data is not found then show the failure message and give an option to login again.
- The eighth function is ```About()```. It will handle the process of authentication. Once a user is logged in, it will show an about page until the user is logged out or the session is closed. If the session is closed then for accessing an about page, a user has to first login.
- The ninth function is ```showAbout()``` It is the used in the ```About()``` function to show the user an about page. It contains two methods. One is ```GET``` and another is ```POST```. 
  - In the ```GET``` method, it will check whether an about field is empty in the database or not? If it is empty then show an about form otherwise show the about page to view the about data.
  - In the ```POST``` method, it will take the about form data and insert/update it in the database.
- The last function is ```Logout()```. It will set the session value to ```false``` and execute the logout template. Once a session is closed, a user has to login again to visit an about page.

# The Third Directory
In the third directory **templates**, I created nine files. All of them are HTML files and each of them show different results. Let's start with the **base.html**.
- In the **base.html** file, I wrote the common HTML code so that I don't have to use it in each file. I called the uncommon code and I will define it in each file according to the page.
- In the **home.html**, I wrote the title of the page, navigation bar, and a body with a simple message.
- In the **signup.html**, I wrote the title of the page, navigation bar, and a signup form in a body.
- In the **login.html**, I wrote the title of the page, navigation bar, and a login form in a body.
- In the **aboutform.html**, I wrote the title of the page, navigation bar, an error message if the data is not inserted in the database and an about form in a body.
- In the **about.html**, I wrote the title of the page, navigation bar, and about data in a body. 
- In the **logout.html**, I wrote the title of the page, navigation bar, and a body with logout message.
- In the **footer.html**, I wrote the footer in a page.
- In the **pageerror.html**, I wrote the title of the page, no navigation bar(just defined it), and a body with a 404 page not found message.

# The Forth Directory
In the forth directory **static**, I created two files and those are CSS files which are responsible for the designing of the pages.
- In the **main.css** file, I designed header, navbar, and footer. 
- In the **form.css** file, I designed the signup, login and the about form.

## Author Info

- YouTube - [ibilalkayy](https://www.youtube.com/channel/UCBLTfRg0Rgm4FtXkvql7DRQ)
- LinkedIn - [@ibilalkayy](https://www.linkedin.com/in/ibilalkayy/)

