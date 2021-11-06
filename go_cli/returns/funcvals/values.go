package funcvals;

func add(i int, j int)int{return i + j}
func sub(i int, j int)int{return i - j}
func mul(i int, j int)int{return i * j}
func div(i int, j int)int{return i / j}

var opMap = map[string]func (int, int) int {
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

var x [3]int


func main(){

}
