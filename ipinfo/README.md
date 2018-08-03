## Pobieranie i obróbka danych z API [ip-api](http://ip-api.com)
 
#### Cel zadania: 
  Nauka komunikacji z API przy pomocy formatu [JSON](https://www.json.org/)

#### Treść zadania:
  Stworzyć program, który przyjmuje adres IP jako argument(flagę) o nazwie `ip`. 
  Następnie program wykonuje zapytanie do [API ip-api](http://ip-api.com/docs/api:json) którego odpowiedź ma zdekodować do wcześniej zadeklarowanego typu(struktury).
  Po zdekodowaniu odpowiedzi, program ma wyświetlić pola struktury w następującej formie:
  ```
    IP address: xxx.xxx.xxx.xxx
    Hostname: xxxx.xxxx.xxxxx
    City: xxxxxx
    Region: xxxxxx
    Country: xxx
    Loc: xxxx
    Postal: xxxxx
  ```

#### Dodatkowe zadania:
  - walidacja czy podany adres IP jest poprawny (np przy pomocy pakietu [net](https://golang.org/pkg/net/) i typu IP)
  - podział struktury programu na [pakiety](https://www.golang-book.com/books/intro/11)  
  - stworzenie drugiej flagi, `geo`, która jest typem `bool` i która sprawia, że wyświetlane są tylko dane geograficzne podanego adresu IP

##### Przykładowe użycie programu ma wyglądać następująco(dane wyjściowe mogą się różnić):
  ```
  go run *.go -ip 81.190.40.214

  IP address: 81.190.40.214
  Organization: Multimedia Polska S. A.
  City: Szczecin
  Region: West Pomerania
  Country: Poland
  Loc: 53.4358, 14.5094
  Postal: 71-132
  ```

#### Przydatne materiały:
  - [dokumentacja API ip-api](http://ip-api.com/docs/api:json)
  - [dokumentacja pakietu net/http](https://golang.org/pkg/net/http/)
  - [przykładowe zapytania przy użyciu pakietu net/http](https://dlintw.github.io/gobyexample/public/http-client.html)
  - [dokumentacja pakietu encoding/json](https://golang.org/pkg/encoding/json/)
  - [przykłady kodowania i dekodowania JSON](https://gobyexample.com/json)
  - [dokumentacja pakietu flag](https://golang.org/pkg/flag/)
  - [przykłady użycia flag w Go](https://gobyexample.com/command-line-flags)
