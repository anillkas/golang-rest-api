<h3>for Golang</h3>
<pre>$ git clone https://github.com/anillkas/golang-rest-api.git</pre>
<pre>$ cd golang-rest-api</pre>
<pre>$ go run main.go</pre>
<pre>http://localhost:8080 or http://127.0.0.1:8080</pre>
<h3>for create Docker Image</h3>
<pre>$ git clone https://github.com/anillkas/golang-rest-api.git</pre>
<pre>$ cd golang-rest-api</pre>
<pre>$ docker build -t golang-rest-api .</pre>
<pre>$ docker run -d -p 8080:8080 --name docker-rest golang-rest-api</pre>
<pre>http://localhost:8080 or http://127.0.0.1:8080</pre>
<h3>for Pull Docker Image</h3>
<pre>$ docker pull anilkas/golang-rest-api</pre>
<pre>$ docker run -d -p 8080:8080 --nane golang-api anilkas/golang-rest-api</pre>
<pre>http://localhost:8080 or http://127.0.0.1:8080</pre>
<h3>How to used</h3>
<pre>http://localhost:8080 or http://host-ip:8080</pre>
<pre>
GET /post/1 
GET /comment/1
GET /todo/1
</pre>
