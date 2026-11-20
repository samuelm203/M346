**Namenskonventionen:** Bezeichner beginnen entweder mit einem Buchstaben oder einem Unterstrich, es wird zwischen Gross- und Kleinschreibung unterschieden, man verwendet camelCase

Werden Var deklariert und nirgend gebraucht kann nicht kompiliert
werden

Bezeichnnung wie folgt: var [Name] [Wert]  

Neue Variable wird mit einem Nullwert des jeweiligen Types initialisiert  (0 bei Ganzzahlen, false bei Boolean oder "" bei Zeichenketten) 

Man kann Variablen auch mit einem Wert initialisieren: 
var [Name] [Wert] = [Wert]

Beispiele 
var result int = 5
var caption string = "Hello World"
var isReady bool = true

Gibt man einen Wert an kann man den Typ weglassen, der Compiler erkennt den Typ automatisch:
var result = 5
var caption = "Hello World"
var isReady = true

Weiter vereinfacht kann man Variablen auch mit dem Kurzschreibweise ":=" deklarieren und initialisieren:
result := 5
caption := "Hello World"
isReady := true

Es ist auch möglich mehrere Variablen gleichzeitig zu deklarieren und zu initialisieren:
var a, b, c = 1, "Hey", false
d, e, f := 2, "Ho", true

### Konstanten 
Konstanten (Spec) lassen sich mit dem Schlüsselwort const definieren
Da sich Konstanten im weitern Verlauf des Programms nicht ändern lassen, muss der Wert bei der Deklaration angegeben werden.
const [Name] [Wert] = [Wert]

Beispiele 
const g float32 = 9.81
const pi float64 = 3.142128393
const name = "Office"

### Primitive Datentypen 
Zahlen
    Ganzzahlen: vorzeichenbehaftet (int) und vorzeichenlos (uint)
    Die Zahl stehen für dei Anzahl der Bits, die für die Zahl verwendet werden.
    Gleitkommazahlen
    
Wahrheitswerte 
Zeichen 
Zeichenketten

### Formatierte Ausgabe
Die formatierte Ausgabe erfolgt mit der Funktion fmt.Printf() und den Formatierungszeichenfolgen.

### Kommentare
Gleich wie Java und C++ werden Kommentare mit // für einzeilige Kommentare und /* */ für mehrzeilige Kommentare geschrieben

### Verzweigungen 
Go kennt die if-else Verzweigung, die switch-case Verzweigung. 

Ein beispiel für die if-else Verzweigung:
```go
x := 5
if x > 0 {
    fmt.Println("x ist positiv")
} else if x < 0 {
    fmt.Println("x ist negativ")
} else {
    fmt.Println("x ist null")
}
```

Ein beispiel für die switch-case Verzweigung:
```go
day := "Montag"
switch day {
case "Montag":
    fmt.Println("Heute ist Montag")
case "Dienstag":
    fmt.Println("Heute ist Dienstag")
default:
    fmt.Println("Heute ist ein anderer Tag")
}
```
Es gibt auch fall-through Verzweigungen, die man mit dem Schlüsselwort fallthrough erzwingen kann.


