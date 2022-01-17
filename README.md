<div id="top"></div>

  <h1 align="center">Test Stockbit Backend Engineer</h1>

</div>

<!-- DOCUMENTATION PROJECT -->
## Documentasi Project

### * Menjalankan program
- Melihat dan menjalankan <b> unit testing</b> yang di buat. <b>"Pertama masuk ke folder soal2"</b> kemudian jalankan di command line
```
./tesh.sh
```
- Menjalankan Program <b>Rest API</b> . <b>"Pertama masuk ke folder soal2"</b> kemudian jalankan di command line
```
go run main.go
```
  setelah program berhasil berjalan, akses sesuai <b>endpoint</b> yang telah dibuat melalui <b>"browser"</b> atau aplikasi API testing seperti <b>"Postman"</b>
<br/>

Berikut fitur dan endpoint Rest API yang terdapat dalam project ini :
| Method | Feature | Endpoint | Authentication
|---|---|---|---|
| GET | Mencari data movie berdasarkan judul dan page | http://127.0.0.1:8080/search/:searchword/:pagination | No |
| GET | Mendapatkan data detail movie | http://127.0.0.1:8080/movie/:imdbID | No |
 
 <br/>
 
 ### * Contoh akses endpoin
 - Contoh mengakses endpoint <b> Mencari data movie berdasarkan judul dan page</b>.akses melalui <b>"browser"</b> atau aplikasi API testing seperti <b>"Postman"</b>
```
http://localhost:8080/search/spiderman/1
```

- Contoh mengakses endpoint <b> detail movie</b>.akses melalui <b>"browser"</b> atau aplikasi API testing seperti <b>"Postman"</b>
```
http://localhost:8080/movie/tt2233044
```

<br/>
 
### Built With

![VS Code](https://img.shields.io/badge/-Visual%20Studio%20Code-05122A?style=flat&logo=visual-studio-code&logoColor=FFFFFF)&nbsp;
![MySQL](https://img.shields.io/badge/-MySQL-05122A?style=flat&logo=mysql&logoColor=FFFFFF)&nbsp;
![Golang](https://img.shields.io/badge/-Golang-05122A?style=flat&logo=go&logoColor=FFFFFF)&nbsp;

<p align="right">(<a href="#top">back to top</a>)</p>

Contributor :
<br>
<p align="left">
    <a href="https://www.linkedin.com/in/nuril-huda-87b279214/" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg"
            alt="nuril huda" height="30" width="40" /></a>
    <a href="https://www.hackerrank.com/nurilhuda7337" target="blank"><img align="center"
            src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/hackerrank.svg"
            alt="nuril7337" height="30" width="40" /></a>
</p>

<p align="right">(<a href="#top">back to top</a>)</p>
<h3>
<p align="center">:copyright: 2022 | Built with :heart: from us</p>
</h3>
<!-- end -->
