## Ping pong
 
#### Cel zadania: 
  Wprowadzenie do kanałów

#### Treść zadania:
  Stworzyć program, który kolejno wypisuje słowa `ping` i `pong`. W tym celu należy wykorzystać jeden kanał `chan string` do przesyłania tych słów oraz 2 gorutyny `ping` oraz `pong`, które przyjmują dany kanał i wykorzystują go do komunikacji/synchronizacji. Po każdym wypisaniu słowa, `ping` i `pong` mają czekać losową ilość milisekund z zakresu `[100, 1000]` przed wpisaniem informacji na kanał. Program ma działać w nieskończonej pętli.

#### Dodatkowe zadania:
  - maybe soon

##### Przykładowe użycie programu ma wyglądać następująco(dane wyjściowe mogą się różnić):
  ```
  $ go run ping.go 
  ping
  pong
  ping
  pong
  ping
  pong
  ping
  pong
  ping
  ``` 

#### Przydatne materiały:
  - [channels w go tour](https://tour.golang.org/concurrency/2)
  - [przykłady channels](https://gobyexample.com/channels)
  - [dokumentacja math/rand](https://golang.org/pkg/math/rand/)
  - [dokumentacja crypto/rand](https://golang.org/pkg/crypto/rand/)
