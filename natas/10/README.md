# Level 10

The $key parameter which is used for grep isn't checked or sanatized at all.

It can be tricked to grep the password like this:
> .* /etc/natas_webpass/natas10 #

More advanced techniques like executing a subshell are also possible but not required.

Anyways..
> next pw: s09byvi8880wqhbnonMFMW8byCojm8eA