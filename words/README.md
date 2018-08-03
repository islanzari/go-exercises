## Odczytywanie oraz obróbka danych z interfejsu io.Reader
 
#### Cel zadania: 
  Nauka pracy z interfejsem [io.Reader](https://golang.org/pkg/io/#Reader)

#### Opis zadania:
  Zadanie polega na uzupełnieniu funkcji `words`, której celem jest odczytanie i obróbka danych z podanego interfejsu `io.Reader`. 
  Funkcja ma zwrócić dwa slice stringów, jeden zawierający wszystkie słowa które posiadają parzystą liczbę samogłosek oraz drugi, który zawiera słowa z nieparzystą liczbą samogłosek.

#### Dodatkowe zadania:
  - stworzyć flagę `file`, która będzie wskazywać na plik, który ma zostać odczytany i wykorzystany jako argument do funkcji `words` (flaga ma domyślnie wskazywać na `lorem.txt`)
  - przebudować program by slice stringów zwracane przez funkcję `words` zostały zapisane do plików `odd.txt` oraz `even.txt`

##### Przykładowe użycie programu ma wyglądać następująco:
  ```
  go run main.go
  [Give, Whose, life, moved ...][shall, together, signs, lights ...]
  ```

#### Przydatne materiały:
  - dokumentacja [io.Reader](https://golang.org/pkg/io/#Reader)
  - dokumetnacja [io.Writer](https://golang.org/pkg/io/#Writer)
  - dokumentacja [os.File](https://golang.org/pkg/os/#File)
  - dokumentacja [ioutil](https://golang.org/pkg/io/ioutil/)