# CLIENT-SERVER
 
### Cel zadania: 
  Wprowadzenie do komunikacji klient<->serwer 

### Treść zadania:
#### Pierwsza część - SERWER
Ta część zadania polega na stworzeniu aplikacji(serwer), która będzie czekać na nowe żądania.
Aplikacja ma posiadać 2 endpointy, które będą służyć do tworzenia nowych użytkowników oraz ich usuwania:
 - tworzenie użytkownika: `POST /users/`
 - usuwanie użytkownika `DELETE /users/{id}/`
 - pobiranie użytkownika `GET /users/{id}/`

Jako multipleksera zapytań(routera) należy użyć `Router` z pakietu [gorilla/mux](https://godoc.org/github.com/gorilla/mux).

Typ użytkownika ma wyglądać następująco:
```go
  type User struct {
    ID      uint64
    Name    string
    Surname string
    Email   string
  }
```

W aplikacji, użytkownicy mają być przetrzymywani w
następującej strukturze:
```go
type database struct {
  m sync.RWMutex
  users map[uint64]User
}
```

Typ `database` powinien posiadać 3 metody, `Delete`, `Get`, `Set`, służące do komunikacji z wewnętrzną mapą.
W celu uniemożliwienia wyścigów podczas zapisu/odczytu/usuwania, należy wykorzystać załączony `RWMutex`. O tym jak to zrobić będzie załączone w przydatnych materiałach.

Aplikacja ma także posiadać gorutynę, której zadaniem bedzie okresowe zapisywanie bazy na dysku(dowolny format) - okres pomiędzy kolejnymi zapisami ma być regulowany za pomocą flagi(dowolna nazwa).
Podczas uruchamiania aplikacji ma być możliwość użycia powyższego pliku do wczytania bazy.

Serwer ma działać na porcie podanym przez flagę(dowolna nazwa).

#### Druga część - KLIENT
Ta część zadania polega na stworzeniu aplikacji(klient), której zadaniem będzie komunikacja z serwerem w celu tworzenia/usuwania/pobierania użytkowników.
Aplikacja ma działać na zasadzie nieskończonej pętli, która będzie czekać na jedno z trzech następujących poleceń:
  - `adduser` - wyświetla formularz tworzenia nowego użytkownika
  - `deleteuser {id}` - usuwa użytkownika na podstawie podanego id
  - `getuser {id}` - pobiera użytkownika na podstawie podanego id

Aplikacja, po otrzymaniu któregokolwiek z tych poleceń ma wykonać odpowiednią akcję. W razie otrzymania niepoprawnego polecenia, aplikacja ma wypisać informację `Unknown task {nazwa niepoprawnego zadania}` i wrócić do pętli.
Aplikacja ma posiadać flagę określającą port, na którym działa serwer.

#### Dodatkowe zadania:
  - użycie kanałów zamiast mutexów. Można wzorować się na [podanym przykładzie](https://medium.com/stupid-gopher-tricks/more-powerful-synchronization-in-go-using-channels-f4a1c3242ed0) 

#### Przydatne materiały:
  - [gorilla/mux](https://godoc.org/github.com/gorilla/mux)
  - [Mutexes by example](https://gobyexample.com/mutexes)
  - [Mutexy w go tour](https://tour.golang.org/concurrency/9)
  - [Dokumentacja sync/mutex](https://golang.org/pkg/sync/#RWMutex)
