grep -q root /etc/passwd  && echo 1
grep -q nobody /etc/passwd && echo 2
grep -q natas16 /etc/passwd && echo 3