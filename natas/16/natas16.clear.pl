# character to extract
$i=$ENV{'haxGetchar'};
# read file to $string
open FILE, "natas17" or print "Nazi's" and exit; 
while (<FILE>){
 $password=$_;
}
close FILE;

print substr($password,$i,1);